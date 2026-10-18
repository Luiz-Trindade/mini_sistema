package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mini_sistema/auth/models"
)

// getEnv obtém o valor da variável de ambiente ou retorna um valor padrão caso não esteja definida
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func InitDB() *gorm.DB {
	// Leitura das variáveis de ambiente com fallbacks padronizados
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "6432")
	user := getEnv("POSTGRES_USER", "app_user")
	password := getEnv("POSTGRES_PASSWORD", "app_password")
	dbname := getEnv("POSTGRES_DB", "mini_sistema")
	sslmode := getEnv("DB_SSLMODE", "disable")
	timezone := getEnv("DB_TIMEZONE", "America/Belem")

	// DSN dinâmico montado a partir das variáveis
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Importante para PgBouncer em modo 'transaction': desabilita prepared statements no driver
		PrepareStmt: false,
	})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados via PgBouncer: %v", err)
	}

	// Executa as Migrations das entidades
	err = db.AutoMigrate(&models.Permission{}, &models.Group{}, &models.User{})
	if err != nil {
		log.Fatalf("Erro ao executar AutoMigrate: %v", err)
	}

	fmt.Printf("Conexão estabelecida com sucesso no banco '%s' via PgBouncer (Porta %s)!\n", dbname, port)
	return db
}
