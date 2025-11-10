package handler

import (
	"net/http"
	tickettier "tctApi/internal/ticket-tier"
	"tctApi/internal/ticket-tier/usecase"
	"tctApi/pkg/helper"
	"tctApi/pkg/response"

	"github.com/gin-gonic/gin"
)

type TicketTierHandler struct {
	ticketTierUC usecase.TicketTierUsecase
}

func NewTicketTierHandler(ticketTierUC usecase.TicketTierUsecase) TicketTierHandler {
	return TicketTierHandler{ticketTierUC}
}

func (tt *TicketTierHandler) Create(c *gin.Context) {
	var req tickettier.TicketTierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	newOrg, err := tt.ticketTierUC.Create(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, newOrg)
}

func (tt *TicketTierHandler) FindAll(c *gin.Context) {
	data, err := tt.ticketTierUC.FindAll()
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, data)
}

func (tt *TicketTierHandler) FindById(c *gin.Context) {
	idUint, err := helper.ConvertToUint(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	ticketTier, err := tt.ticketTierUC.FindById(idUint)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, ticketTier)
}

func (tt *TicketTierHandler) Update(c *gin.Context) {
	idUint, err := helper.ConvertToUint(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var req tickettier.TicketTierRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	ticketTier, err := tt.ticketTierUC.Update(idUint, &req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, ticketTier)
}

func (tt *TicketTierHandler) Delete(c *gin.Context) {
	idUint, err := helper.ConvertToUint(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err = tt.ticketTierUC.Delete(idUint)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
	}

	response.Success(c, nil)
}
