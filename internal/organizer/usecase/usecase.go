package usecase

import (
	"fmt"
	"tctApi/internal/organizer"
	"tctApi/internal/organizer/repository"
)

type OrganizerUsecase interface {
	FindAll() (*[]organizer.Organizer, error)
	Create(input *organizer.OrganizerRequest) (*organizer.Organizer, error)
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
