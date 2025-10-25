package handler

import (
	"net/http"
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

func (o *OrganizerHandler) FindAll(c *gin.Context) {
	data, err := o.organizerUC.FindAll()
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, data)
}
