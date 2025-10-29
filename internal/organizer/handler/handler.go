package handler

import (
	"errors"
	"net/http"
	"strconv"
	"tctApi/internal/organizer"
	"tctApi/internal/organizer/usecase"
	"tctApi/pkg/response"

	"github.com/gin-gonic/gin"
)

type OrganizerHandler struct {
	organizerUC usecase.OrganizerUsecase
}

func NewOrganizerHandler(organizerUC usecase.OrganizerUsecase) *OrganizerHandler {
	return &OrganizerHandler{organizerUC}
}

func (o *OrganizerHandler) Create(c *gin.Context) {
	var req organizer.OrganizerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	newOrg, err := o.organizerUC.Create(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, newOrg)
}

func (o *OrganizerHandler) FindAll(c *gin.Context) {
	data, err := o.organizerUC.FindAll()
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, data)
}

func (o *OrganizerHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		response.Error(c, http.StatusBadRequest, errors.New("invalid id").Error())
		return
	}

	var req organizer.OrganizerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	org, err := o.organizerUC.Update(uint(idUint), &req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, org)
}

func (o *OrganizerHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		response.Error(c, http.StatusBadRequest, errors.New("invalid id").Error())
		return
	}

	err = o.organizerUC.Delete(uint(idUint))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, nil)
}
