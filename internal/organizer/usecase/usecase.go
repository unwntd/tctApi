package usecase

import (
	"fmt"
	"tctApi/internal/organizer"
	"tctApi/internal/organizer/repository"
	"time"
)

type OrganizerUsecase interface {
	FindAll() (*[]organizer.Organizer, error)
	Create(input *organizer.OrganizerRequest) (*organizer.Organizer, error)
	Update(organizerID uint, input *organizer.OrganizerRequest) (*organizer.Organizer, error)
	Delete(organizerID uint) error
}

type organizerUsecase struct {
	repo repository.OrganizerRepository
}

func NewOrganizerUsecase(repo repository.OrganizerRepository) OrganizerUsecase {
	return &organizerUsecase{repo: repo}
}

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

func (o *organizerUsecase) FindAll() (*[]organizer.Organizer, error) {
	organizers, err := o.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all organizers: %w", err)
	}
	return organizers, nil
}

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
		return nil, fmt.Errorf("failed to create organizer: %w", err)
	}

	return org, nil
}

func (o *organizerUsecase) Delete(organizerID uint) error {
	_, err := o.repo.FindById(organizerID)
	if err != nil {
		return err
	}

	return o.repo.Delete(organizerID)
}
