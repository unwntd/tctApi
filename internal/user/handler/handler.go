package handler

import (
	"net/http"
	"tctApi/internal/user"
	"tctApi/internal/user/usecase"
	"tctApi/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUC usecase.UserUsecase
}

func NewUserHandler(userUC usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUC}
}

func (h *UserHandler) RegisterRoutes(r *gin.RouterGroup) {
	users := r.Group("/users")
	users.POST("/register", h.Register)
	users.POST("/login", h.Login)
}

func (h *UserHandler) Register(c *gin.Context) {
	var req user.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := h.userUC.Register(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, newUser)
}

func (h *UserHandler) Login(c *gin.Context) {
	var req user.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	token, userData, err := h.userUC.Login(&req)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	resp := user.AuthResponse{
		Token: token,
		User:  userData,
	}

	response.Success(c, resp)
}
