package repository

import (
	"tctApi/internal/organizer"

	"gorm.io/gorm"
)

type OrganizerRepository interface {
	Create(o *organizer.Organizer) error
	FindAll() (*[]organizer.Organizer, error)
	FindById(id uint) (*organizer.Organizer, error)
	Delete(id uint) error
	Update(o *organizer.Organizer) error
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
	var organizers []organizer.Organizer
	err := r.db.Find(&organizers).Error
	return &organizers, err
}

func (r *organizerRepository) FindById(id uint) (*organizer.Organizer, error) {
	var organizer organizer.Organizer
	err := r.db.First(&organizer, id).Error
	if err != nil {
		return nil, err
	}

	return &organizer, nil
}

func (r *organizerRepository) Delete(id uint) error {
	return r.db.Delete(&organizer.Organizer{}, id).Error
}

func (r *organizerRepository) Update(o *organizer.Organizer) error {
	return r.db.Model(&organizer.Organizer{}).Where("id = ?", o.ID).Updates(o).Error
}
