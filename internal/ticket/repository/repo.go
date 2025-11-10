package repository

import (
	"tctApi/internal/ticket"

	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(t *ticket.Ticket) error
	FindAll() (*[]ticket.Ticket, error)
	FindById(id uint) (*ticket.Ticket, error)
	Update(t *ticket.Ticket) error
	Delete(id uint) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) Create(t *ticket.Ticket) error {
	return r.db.Create(t).Error
}

func (r *ticketRepository) FindAll() (*[]ticket.Ticket, error) {
	var tickets []ticket.Ticket
	err := r.db.Preload("Tier").Find(&tickets).Error
	return &tickets, err
}

func (r *ticketRepository) FindById(id uint) (*ticket.Ticket, error) {
	var t ticket.Ticket
	err := r.db.Preload("Tier").First(&t, id).Error
	return &t, err
}

func (r *ticketRepository) Update(t *ticket.Ticket) error {
	return r.db.Model(t).Updates(t).Error
}

func (r *ticketRepository) Delete(id uint) error {
	return r.db.Delete(&ticket.Ticket{}, id).Error
}
