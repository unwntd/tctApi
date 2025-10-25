package repository

import (
	"tctApi/internal/organizer"

	"gorm.io/gorm"
)

type OrganizerRepository interface {
	Create(o *organizer.Organizer) error
	FindAll() (*[]organizer.Organizer, error)
	Delete(id uint) error
	Update(id uint) error
}

type organizerRepository struct {
	db *gorm.DB
}

func NewOrganizerRepository(db *gorm.DB) OrganizerRepository {
	return &organizerRepository{db}
}

func (r *organizerRepository) Create(o *organizer.Organizer) error {
	return r.db.Create(o).Error
}

func (r *organizerRepository) FindAll() (*[]organizer.Organizer, error) {
	var organizers *[]organizer.Organizer
	err := r.db.Find(&organizers).Error
	return organizers, err
}

func (r *organizerRepository) Delete(id uint) error {
	return r.db.Delete(&organizer.Organizer{}, id).Error
}

func (r *organizerRepository) Update(id uint) error {
	return r.db.Model(&organizer.Organizer{}).Where("id = ?", id).Updates(organizer.Organizer{}).Error
}
