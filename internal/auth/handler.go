package auth

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service AuthService
}

func NewAuthHandler(s AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) RegisterRoutes(r *gin.RouterGroup) {
	users := r.Group("/auth")
	users.POST("/login", h.Login)
	users.POST("/refresh", h.Refresh)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, refreshToken, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Store refresh token in HttpOnly cookie
	c.SetCookie("refresh_token", refreshToken, 7*24*3600, "/", "", true, true)
	c.JSON(http.StatusOK, tokens)
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshID, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing refresh token"})
		return
	}

	userID, _ := strconv.Atoi(c.Query("user_id")) // or from claims/context
	tokens, newRefreshTokens, err := h.service.Refresh(uint(userID), refreshID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Rotate refresh cookie
	c.SetCookie("refresh_token", newRefreshTokens, 7*24*3600, "/", "", true, true)
	c.JSON(http.StatusOK, tokens)
}
