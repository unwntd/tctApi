package repository

import (
	"tctApi/internal/organizer"

	"gorm.io/gorm"
)

type OrganizerRepository interface {
	Create(o *organizer.Organizer) error
	FindAll() (*[]organizer.Organizer, error)
	FindById(id uint) (*organizer.Organizer, error)
	Update(o *organizer.Organizer) error
	Delete(id uint) error
}

type organizerRepository struct {
	db *gorm.DB
}

// NewOrganizerRepository returns a new instance of OrganizerRepository with the given database.
// It is used to create a new OrganizerRepository for database operations.
func NewOrganizerRepository(db *gorm.DB) OrganizerRepository {
	return &organizerRepository{db}
}

// Create a new organizer with the given information. Returns an error if creation fails.
// The new organizer is created in the database if creation is successful.
// The created organizer is not returned.
func (r *organizerRepository) Create(o *organizer.Organizer) error {
	return r.db.Create(o).Error
}

// FindAll retrieves all organizers from database
// Returns an error if retrieval fails
// The retrieved organizers are returned as a slice of organizer.Organizer structs.
func (r *organizerRepository) FindAll() (*[]organizer.Organizer, error) {
	var organizers []organizer.Organizer
	err := r.db.Find(&organizers).Error
	return &organizers, err
}

// Find an organizer by ID. Returns an error if the organizer does not exist or if retrieval fails.
// The organizer is returned as a pointer to an organizer.Organizer struct.
func (r *organizerRepository) FindById(id uint) (*organizer.Organizer, error) {
	var organizer organizer.Organizer
	err := r.db.First(&organizer, id).Error
	return &organizer, err
}

// Update an organizer by ID. Returns an error if deletion fails.
// The organizer is updated in the database if deletion is successful.
// The updated organizer is not returned.
func (r *organizerRepository) Update(o *organizer.Organizer) error {
	return r.db.Model(o).Updates(o).Error
}

// Delete an organizer by ID. Returns an error if deletion fails.
// The organizer is deleted from the database if deletion is successful.
// The deleted organizer is not returned.
func (r *organizerRepository) Delete(id uint) error {
	return r.db.Delete(&organizer.Organizer{}, id).Error
}
