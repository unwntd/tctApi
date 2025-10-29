package handler

import (
	"net/http"
	"tctApi/internal/role"
	"tctApi/internal/role/usecase"
	"tctApi/pkg/helper"
	"tctApi/pkg/response"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleUC usecase.RoleUsecase
}

func NewRoleHandler(roleUC usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{roleUC}
}

func (r *RoleHandler) Create(c *gin.Context) {
	var req role.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	newRole, err := r.roleUC.Create(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, newRole)
}

func (r *RoleHandler) FindAll(c *gin.Context) {
	data, err := r.roleUC.FindAll()
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, data)
}

func (r *RoleHandler) Update(c *gin.Context) {
	idUint, err := helper.ConvertToUint(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var req role.RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	role, err := r.roleUC.Update(uint(idUint), &req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, role)
}

func (r *RoleHandler) Delete(c *gin.Context) {
	idUint, err := helper.ConvertToUint(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err = r.roleUC.Delete(uint(idUint))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, nil)
}
