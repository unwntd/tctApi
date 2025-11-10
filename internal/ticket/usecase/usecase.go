package usecase

import (
	"errors"
	"fmt"
	"tctApi/internal/ticket"
	ticketTierRepo "tctApi/internal/ticket-tier/repository"
	"tctApi/internal/ticket/repository"
	"time"
)

type TicketUsecase interface {
	Create(input *ticket.TicketRequest) (*ticket.Ticket, error)
	FindAll() (*[]ticket.Ticket, error)
	FindById(id uint) (*ticket.Ticket, error)
	Update(id uint, input *ticket.TicketRequest) (*ticket.Ticket, error)
	Delete(id uint) error
}

type ticketUsecase struct {
	repo     repository.TicketRepository
	tierRepo ticketTierRepo.TicketTierRepository
}

func NewTicketUsecase(repo repository.TicketRepository, tierRepo ticketTierRepo.TicketTierRepository) TicketUsecase {
	return &ticketUsecase{repo: repo, tierRepo: tierRepo}
}

func (t *ticketUsecase) Create(input *ticket.TicketRequest) (*ticket.Ticket, error) {
	tier, err := t.tierRepo.FindById(input.TierID)
	msg := fmt.Sprintf("failed to create Ticket: ticket tier with ID %d does not exist", input.TierID)
	if err != nil || tier == nil {
		return nil, errors.New(msg)
	}

	ticket := &ticket.Ticket{
		TierID:       input.TierID,
		Name:         input.Name,
		Description:  input.Description,
		Price:        input.Price,
		Tax:          *input.Tax,
		Limit:        *input.Limit,
		AvailableQty: *input.AvailableQty,
		PendingQty:   *input.PendingQty,
		StartPeriod:  input.StartPeriod,
		EndPeriod:    input.EndPeriod,
		IsRefundable: *input.IsRefundable,
	}

	if err := t.repo.Create(ticket); err != nil {
		return nil, fmt.Errorf("failed to create Ticket: %w", err)
	}

	newTicket, err := t.repo.FindById(ticket.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated Ticket: %w", err)
	}

	return newTicket, nil
}

func (t *ticketUsecase) FindAll() (*[]ticket.Ticket, error) {
	tickets, err := t.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all Tickets: %w", err)
	}

	return tickets, nil
}

func (t *ticketUsecase) FindById(ticketId uint) (*ticket.Ticket, error) {
	ticket, err := t.repo.FindById(ticketId)
	if err != nil {
		return nil, fmt.Errorf("failed to find Ticket: %w", err)
	}

	return ticket, nil
}

func (t *ticketUsecase) Update(ticketId uint, input *ticket.TicketRequest) (*ticket.Ticket, error) {
	existingTicket, err := t.repo.FindById(ticketId)
	if err != nil {
		return nil, err
	}

	tier, err := t.tierRepo.FindById(input.TierID)
	if err != nil {
		return nil, fmt.Errorf("failed to create Ticket: %w", err)
	}

	if tier == nil {
		return nil, fmt.Errorf("failed to create Ticket: ticket tier with ID %d does not exist", input.TierID)
	}

	ticket := &ticket.Ticket{
		ID:           existingTicket.ID,
		TierID:       input.TierID,
		Name:         input.Name,
		Description:  input.Description,
		Price:        input.Price,
		Tax:          *input.Tax,
		Limit:        *input.Limit,
		AvailableQty: *input.AvailableQty,
		PendingQty:   *input.PendingQty,
		StartPeriod:  input.StartPeriod,
		EndPeriod:    input.EndPeriod,
		IsRefundable: *input.IsRefundable,
		CreatedAt:    existingTicket.CreatedAt,
		UpdatedAt:    time.Now(),
	}

	if err := t.repo.Update(ticket); err != nil {
		return nil, fmt.Errorf("failed to update Ticket: %w", err)
	}

	updatedTicket, err := t.repo.FindById(ticketId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated Ticket: %w", err)
	}

	return updatedTicket, nil
}

func (t *ticketUsecase) Delete(ticketId uint) error {
	_, err := t.repo.FindById(ticketId)
	if err != nil {
		return err
	}

	if err := t.repo.Delete(ticketId); err != nil {
		return fmt.Errorf("failed to delete Ticket: %w", err)
	}

	return nil
}
