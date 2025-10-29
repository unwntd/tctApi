package repository

import (
	"tctApi/internal/role"

	"gorm.io/gorm"
)

type RoleRepository interface {
	Create(r *role.Role) error
	FindAll() (*[]role.Role, error)
	FindById(id uint) (*role.Role, error)
	Update(r *role.Role) error
	Delete(id uint) error
	IsExistByName(roleID uint, name string) (bool, error)
}

type roleRepository struct {
	db *gorm.DB
}

// NewRoleRepository returns a new instance of RoleRepository with the given database.
// It is used to create a new RoleRepository for database operations.
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{db}
}

// Create a new role with the given information. Returns an error if creation fails.
func (r *roleRepository) Create(rl *role.Role) error {
	return r.db.Create(rl).Error
}

// FindAll retrieves all roles from database
// Returns an error if retrieval fails
func (r *roleRepository) FindAll() (*[]role.Role, error) {
	var roles []role.Role
	err := r.db.Find(&roles).Error
	return &roles, err
}

// Find a role by ID. Returns an error if the role does not exist or if retrieval fails.
// The role is returned as a pointer to a role.Role struct.
func (r *roleRepository) FindById(id uint) (*role.Role, error) {
	var role role.Role
	err := r.db.First(&role, id).Error
	return &role, err
}

// Update a role by ID. Returns an error if deletion fails.
func (r *roleRepository) Update(rl *role.Role) error {
	return r.db.Model(&role.Role{}).Where("id = ?", rl.ID).Updates(rl).Error
}

// Delete a role by ID. Returns an error if deletion fails.
func (r *roleRepository) Delete(id uint) error {
	return r.db.Delete(&role.Role{}, id).Error
}

// IsExistByName checks if a role with the given name already exists in the database.
// If isUpdate is true, it will ignore the role with the given ID.
// It returns true if the role exists, and an error if the query fails.
// If the role does not exist, it returns false and nil error.
func (r *roleRepository) IsExistByName(roleID uint, name string) (bool, error) {
	var count int64

	query := r.db.Model(&role.Role{}).Where("name = ?", name)
	if roleID > 0 {
		query = query.Where("id != ?", roleID)
	}

	if err := query.Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
