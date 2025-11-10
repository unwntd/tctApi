package usecase

import (
	"fmt"
	tickettier "tctApi/internal/ticket-tier"
	"tctApi/internal/ticket-tier/repository"
	"time"
)

type TicketTierUsecase interface {
	Create(input *tickettier.TicketTierRequest) (*tickettier.TicketTier, error)
	FindAll() (*[]tickettier.TicketTier, error)
	FindById(id uint) (*tickettier.TicketTier, error)
	Update(id uint, input *tickettier.TicketTierRequest) (*tickettier.TicketTier, error)
	Delete(id uint) error
}

type ticketTierUsecase struct {
	repo repository.TicketTierRepository
}

func NewTicketTierUsecase(repo repository.TicketTierRepository) TicketTierUsecase {
	return &ticketTierUsecase{repo: repo}
}

func (tt *ticketTierUsecase) Create(input *tickettier.TicketTierRequest) (*tickettier.TicketTier, error) {
	newTicketTier := &tickettier.TicketTier{
		EventID:      input.EventID,
		Name:         input.Name,
		Description:  input.Description,
		Limit:        *input.Limit,
		AvailableQty: *input.AvailableQty,
		PendingQty:   *input.PendingQty,
	}

	if err := tt.repo.Create(newTicketTier); err != nil {
		return nil, fmt.Errorf("failed to create Ticket Tier: %w", err)
	}

	return newTicketTier, nil
}

func (tt *ticketTierUsecase) FindAll() (*[]tickettier.TicketTier, error) {
	ticketTiers, err := tt.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all Ticket Tier: %w", err)
	}

	return ticketTiers, nil
}

func (tt *ticketTierUsecase) FindById(ticketTierId uint) (*tickettier.TicketTier, error) {
	ticketTier, err := tt.repo.FindById(ticketTierId)
	if err != nil {
		return nil, fmt.Errorf("failed to find Ticket Tier: %w", err)
	}

	return ticketTier, nil
}

func (tt *ticketTierUsecase) Update(ticketTierId uint, input *tickettier.TicketTierRequest) (*tickettier.TicketTier, error) {
	existingTicketTier, err := tt.repo.FindById(ticketTierId)
	if err != nil {
		return nil, err
	}

	ticketTier := &tickettier.TicketTier{
		ID:           ticketTierId,
		EventID:      input.EventID,
		Name:         input.Name,
		Description:  input.Description,
		Limit:        *input.Limit,
		AvailableQty: *input.AvailableQty,
		PendingQty:   *input.PendingQty,
		CreatedAt:    existingTicketTier.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	if err := tt.repo.Update(ticketTier); err != nil {
		return nil, fmt.Errorf("failed to update Ticket Tier: %w", err)
	}

	return ticketTier, nil
}

func (tt *ticketTierUsecase) Delete(ticketTierID uint) error {
	_, err := tt.repo.FindById(ticketTierID)
	if err != nil {
		return fmt.Errorf("failed to delete Ticket Tier: %w", err)
	}

	return tt.repo.Delete(ticketTierID)
}
