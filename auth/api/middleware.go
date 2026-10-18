package api

import (
	"net/http"
	"strings"

	"mini_sistema/auth/selectors"
	"mini_sistema/auth/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *services.JWTService, userSelector *selectors.UserSelector) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Header Authorization não informado"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido. Use 'Bearer <token>'"})
			return
		}

		claims, err := jwtService.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Busca o usuário atual para checar se a versão do token ainda é válida
		user, err := userSelector.GetByID(claims.UserID)
		if err != nil || user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
			return
		}

		// Se a versão do token for diferente da versão atual no banco, o token foi revogado por um novo login
		if claims.TokenVersion != user.TokenVersion {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Sessão expirada devido a um novo login em outro dispositivo"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.Email)
		c.Set("isAdmin", claims.IsAdmin)

		c.Next()
	}
}
