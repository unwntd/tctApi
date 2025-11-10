package usecase

import (
	"fmt"
	"tctApi/internal/organizer"
	"tctApi/internal/organizer/repository"
	"time"
)

type OrganizerUsecase interface {
	Create(input *organizer.OrganizerRequest) (*organizer.Organizer, error)
	FindAll() (*[]organizer.Organizer, error)
	FindById(organizerID uint) (*organizer.Organizer, error)
	Update(organizerID uint, input *organizer.OrganizerRequest) (*organizer.Organizer, error)
	Delete(organizerID uint) error
}

type organizerUsecase struct {
	repo repository.OrganizerRepository
}

// NewOrganizerUsecase returns a new instance of OrganizerUsecase with the given repository.
// The returned OrganizerUsecase is used to perform operations on organizers such as creating, updating, and deleting them.
// The repository is used to interact with the database.
func NewOrganizerUsecase(repo repository.OrganizerRepository) OrganizerUsecase {
	return &organizerUsecase{repo: repo}
}

// Create a new organizer with the given information. Returns an error if creation fails.
// The new organizer is returned along with a nil error if creation is successful.
func (o *organizerUsecase) Create(input *organizer.OrganizerRequest) (*organizer.Organizer, error) {
	newOrg := &organizer.Organizer{
		Name:           input.Name,
		Address:        input.Address,
		PicName:        input.PicName,
		PicPhoneNumber: input.PicPhoneNumber,
		PicEmail:       input.PicEmail,
		LogoURL:        input.LogoURL,
	}

	if err := o.repo.Create(newOrg); err != nil {
		return nil, fmt.Errorf("failed to create organizer: %w", err)
	}

	return newOrg, nil
}

// FindAll retrieves all organizers from database
// Returns an error if retrieval fails
func (o *organizerUsecase) FindAll() (*[]organizer.Organizer, error) {
	organizers, err := o.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all organizers: %w", err)
	}
	return organizers, nil
}

func (o *organizerUsecase) FindById(organizerId uint) (*organizer.Organizer, error) {
	org, err := o.repo.FindById(organizerId)
	if err != nil {
		return nil, fmt.Errorf("failed to find organizer: %w", err)
	}

	return org, nil
}

// Update an organizer by ID. Returns an error if the organizer does not exist or if deletion fails.
// The updated organizer is returned along with a nil error if the update is successful.
func (o *organizerUsecase) Update(organizerID uint, input *organizer.OrganizerRequest) (*organizer.Organizer, error) {
	existingOrganizer, err := o.repo.FindById(organizerID)
	if err != nil {
		return nil, err
	}

	org := &organizer.Organizer{
		ID:             existingOrganizer.ID,
		Name:           input.Name,
		Address:        input.Address,
		PicName:        input.PicName,
		PicPhoneNumber: input.PicPhoneNumber,
		PicEmail:       input.PicEmail,
		LogoURL:        input.LogoURL,
		CreatedAt:      existingOrganizer.CreatedAt,
		UpdatedAt:      time.Now(),
	}

	if err := o.repo.Update(org); err != nil {
		return nil, fmt.Errorf("failed to update organizer: %w", err)
	}

	return org, nil
}

// Delete an organizer by ID. Returns an error if the organizer does not exist or if deletion fails.
func (o *organizerUsecase) Delete(organizerID uint) error {
	_, err := o.repo.FindById(organizerID)
	if err != nil {
		return fmt.Errorf("failed to delete organizer: %w", err)
	}

	return o.repo.Delete(organizerID)
}
