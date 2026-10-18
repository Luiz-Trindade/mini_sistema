package api

import (
	"net/http"

	"mini_sistema/auth/selectors"
	"mini_sistema/auth/services"

	"github.com/gin-gonic/gin"
)

type LoginDTO struct {
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required"`
}

type AuthHandler struct {
	authService  *services.AuthService
	jwtService   *services.JWTService
	userSelector *selectors.UserSelector
}

func NewAuthHandler(authService *services.AuthService, jwtService *services.JWTService, userSelector *selectors.UserSelector) *AuthHandler {
	return &AuthHandler{
		authService:  authService,
		jwtService:   jwtService,
		userSelector: userSelector,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var dto LoginDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados da requisição inválidos"})
		return
	}

	user, err := h.authService.Authenticate(dto.Email, dto.Senha)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := h.jwtService.GenerateToken(user.ID, user.Email, user.IsAdmin, user.TokenVersion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token de acesso"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func (h *AuthHandler) Me(c *gin.Context) {
	userID, _ := c.Get("userID")
	email, _ := c.Get("userEmail")
	isAdmin, _ := c.Get("isAdmin")

	c.JSON(http.StatusOK, gin.H{
		"user_id":  userID,
		"email":    email,
		"is_admin": isAdmin,
	})
}

// Logout lê o userID do contexto da requisição protegida e invalida os tokens ativos
func (h *AuthHandler) Logout(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	userID := userIDVal.(uint)

	if err := h.authService.RevokeUserTokens(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao efetuar logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout realizado com sucesso. Sessão encerrada.",
	})
}

func (h *AuthHandler) RegisterRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", h.Login)

		protected := authGroup.Group("")
		protected.Use(AuthMiddleware(h.jwtService, h.userSelector))
		{
			protected.GET("/me", h.Me)
			protected.POST("/logout", h.Logout) // Nova rota protegida
		}
	}
}
