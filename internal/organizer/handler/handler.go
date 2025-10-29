package handler

import (
	"net/http"
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
