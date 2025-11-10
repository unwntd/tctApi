package handler

import (
	"net/http"
	"tctApi/internal/ticket"
	"tctApi/internal/ticket/usecase"
	"tctApi/pkg/helper"
	"tctApi/pkg/response"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	ticketUC usecase.TicketUsecase
}

func NewTicketHandler(ticketUC usecase.TicketUsecase) TicketHandler {
	return TicketHandler{ticketUC}
}

func (t *TicketHandler) Create(c *gin.Context) {
	var req ticket.TicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	newTicket, err := t.ticketUC.Create(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, newTicket)
}

func (t *TicketHandler) FindAll(c *gin.Context) {
	data, err := t.ticketUC.FindAll()
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, data)
}

func (t *TicketHandler) FindById(c *gin.Context) {
	idUint, err := helper.ConvertToUint(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	ticket, err := t.ticketUC.FindById(idUint)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, ticket)
}

func (t *TicketHandler) Update(c *gin.Context) {
	idUint, err := helper.ConvertToUint(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var req ticket.TicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	ticket, err := t.ticketUC.Update(idUint, &req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, ticket)
}

func (t *TicketHandler) Delete(c *gin.Context) {
	idUint, err := helper.ConvertToUint(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err = t.ticketUC.Delete(idUint)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, nil)
}
