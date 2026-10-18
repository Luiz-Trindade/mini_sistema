package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"mini_sistema/auth/api"
	"mini_sistema/auth/commands"
	"mini_sistema/auth/selectors"
	"mini_sistema/auth/services"
	"mini_sistema/database"

	"github.com/gin-gonic/gin"
)

// Embutindo o build estático do Quasar
//
//go:embed frontend/dist/spa/*
var staticFiles embed.FS

func main() {
	createSuperuserFlag := flag.Bool("createsuperuser", false, "Cria um superusuário interactively via terminal")
	flag.Parse()

	db := database.InitDB()

	userSelector := selectors.NewUserSelector(db)
	authService := services.NewAuthService(db, userSelector)
	userService := services.NewUserService(db, userSelector)
	jwtService := services.NewJWTService()

	if *createSuperuserFlag {
		superuserCmd := commands.NewSuperuserCommand(userService)
		superuserCmd.Run()
		return
	}

	authHandler := api.NewAuthHandler(authService, jwtService, userSelector)
	r := gin.Default()

	// Registra as rotas da API (ex: /auth/login)
	authHandler.RegisterRoutes(r)

	// Isola o sub-filesystem apontando para a pasta do build do Quasar
	distFS, err := fs.Sub(staticFiles, "frontend/dist/spa")
	if err != nil {
		log.Fatalf("Erro ao carregar arquivos estáticos do Quasar: %v", err)
	}

	// Pré-carrega o index.html em memória (evita redirecionamentos 301 e acelera a resposta)
	indexHTML, err := fs.ReadFile(distFS, "index.html")
	if err != nil {
		log.Fatalf("Erro ao ler index.html do frontend: %v", err)
	}

	fileServer := http.FileServer(http.FS(distFS))

	// Handler para servir arquivos estáticos e realizar o fallback do Vue Router
	r.NoRoute(func(c *gin.Context) {
		// 1. Se for uma requisição de API não encontrada, retorna 404 em JSON
		if strings.HasPrefix(c.Request.URL.Path, "/api") || strings.HasPrefix(c.Request.URL.Path, "/auth") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Endpoint de API não encontrado"})
			return
		}

		path := strings.TrimPrefix(c.Request.URL.Path, "/")

		// 2. Se for um arquivo físico válido no build (ex: JS, CSS, favicon, fontes)
		if path != "" {
			if f, err := distFS.Open(path); err == nil {
				stat, statErr := f.Stat()
				_ = f.Close()

				if statErr == nil && !stat.IsDir() {
					fileServer.ServeHTTP(c.Writer, c.Request)
					return
				}
			}
		}

		// 3. Para a raiz "/" ou rotas do Vue Router (ex: /login), entrega o index.html diretamente com HTTP 200
		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHTML)
	})

	fmt.Println("Servidor rodando em http://localhost:8080")
	r.Run(":8080")
}
