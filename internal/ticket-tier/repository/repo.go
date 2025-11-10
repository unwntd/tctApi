package repository

import (
	tickettier "tctApi/internal/ticket-tier"

	"gorm.io/gorm"
)

type TicketTierRepository interface {
	Create(tt *tickettier.TicketTier) error
	FindAll() (*[]tickettier.TicketTier, error)
	FindById(id uint) (*tickettier.TicketTier, error)
	Update(tt *tickettier.TicketTier) error
	Delete(id uint) error
}

type ticketTierRepository struct {
	db *gorm.DB
}

func NewTicketTierRepository(db *gorm.DB) TicketTierRepository {
	return &ticketTierRepository{db}
}

func (r *ticketTierRepository) Create(tt *tickettier.TicketTier) error {
	return r.db.Create(tt).Error
}

func (r *ticketTierRepository) FindAll() (*[]tickettier.TicketTier, error) {
	var ticketTiers []tickettier.TicketTier
	err := r.db.Find(&ticketTiers).Error
	return &ticketTiers, err
}

func (r *ticketTierRepository) FindById(id uint) (*tickettier.TicketTier, error) {
	var ticketTier tickettier.TicketTier
	err := r.db.First(&ticketTier, id).Error
	return &ticketTier, err
}

func (r *ticketTierRepository) Update(tt *tickettier.TicketTier) error {
	return r.db.Model(tt).Updates(tt).Error
}

func (r *ticketTierRepository) Delete(id uint) error {
	return r.db.Delete(&tickettier.TicketTier{}, id).Error
}
