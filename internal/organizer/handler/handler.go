package handler

import (
	"net/http"
	"tctApi/internal/organizer"
	"tctApi/internal/organizer/usecase"
	"tctApi/pkg/helper"
	"tctApi/pkg/response"

	"github.com/gin-gonic/gin"
)

type OrganizerHandler struct {
	organizerUC usecase.OrganizerUsecase
}

// NewOrganizerHandler returns a new instance of OrganizerHandler with the given OrganizerUsecase.
// The returned OrganizerHandler is used to perform operations on organizers such as creating, updating, and deleting them.
// The OrganizerUsecase is used to interact with the database.
func NewOrganizerHandler(organizerUC usecase.OrganizerUsecase) *OrganizerHandler {
	return &OrganizerHandler{organizerUC}
}

// Create a new organizer with the given information. Returns an error if creation fails.
// The new organizer is returned along with a nil error if creation is successful.
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

// FindAll retrieves all organizers from database
// Returns an error if retrieval fails
func (o *OrganizerHandler) FindAll(c *gin.Context) {
	data, err := o.organizerUC.FindAll()
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, data)
}

// Update an organizer by ID. Returns an error if the organizer does not exist or if deletion fails.
// The updated organizer is returned along with a nil error if the update is successful.
func (o *OrganizerHandler) Update(c *gin.Context) {
	idUint, err := helper.ConvertToUint(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
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

// Delete an organizer by ID. Returns an error if deletion fails.
func (o *OrganizerHandler) Delete(c *gin.Context) {
	idUint, err := helper.ConvertToUint(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	err = o.organizerUC.Delete(uint(idUint))
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, nil)
}
