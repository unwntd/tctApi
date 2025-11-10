package usecase

import (
	"fmt"
	"tctApi/internal/role"
	"tctApi/internal/role/repository"
)

type RoleUsecase interface {
	Create(input *role.RoleRequest) (*role.Role, error)
	FindAll() (*[]role.Role, error)
	FindById(id uint) (*role.Role, error)
	Update(id uint, input *role.RoleRequest) (*role.Role, error)
	Delete(id uint) error
}

type roleUsecase struct {
	repo repository.RoleRepository
}

// NewRoleUsecase returns a new instance of RoleUsecase with the given repository.
func NewRoleUsecase(repo repository.RoleRepository) RoleUsecase {
	return &roleUsecase{repo: repo}
}

// Create a new role with the given name. Returns an error if the role already exists or if creation fails.
func (r *roleUsecase) Create(input *role.RoleRequest) (*role.Role, error) {
	exists, err := r.repo.IsExistByName(0, input.Name)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("role with name %s already exists", input.Name)
	}

	newRole := &role.Role{Name: input.Name}
	err = r.repo.Create(newRole)
	if err != nil {
		return nil, fmt.Errorf("failed to create role: %w", err)
	}

	return newRole, nil
}

// FindAll retrieves all roles from database
// Returns an error if retrieval fails
func (r *roleUsecase) FindAll() (*[]role.Role, error) {
	roles, err := r.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to find all roles: %w", err)
	}

	return roles, nil
}

func (r *roleUsecase) FindById(id uint) (*role.Role, error) {
	role, err := r.repo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find role: %w", err)
	}

	return role, err
}

// Update a role by ID. Returns an error if the role does not exist or if deletion fails.
func (r *roleUsecase) Update(roleID uint, input *role.RoleRequest) (*role.Role, error) {
	existingRole, err := r.repo.FindById(roleID)
	if err != nil {
		return nil, err
	}

	exists, err := r.repo.IsExistByName(roleID, input.Name)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("role with name %s already exists", input.Name)
	}

	role := &role.Role{
		ID:        roleID,
		Name:      input.Name,
		CreatedAt: existingRole.CreatedAt,
		UpdatedAt: existingRole.UpdatedAt,
	}

	if err := r.repo.Update(role); err != nil {
		return nil, fmt.Errorf("failed to update role: %w", err)
	}

	return role, nil
}

// Delete a role by ID. Returns an error if the role does not exist or if deletion fails.
func (r *roleUsecase) Delete(roleID uint) error {
	_, err := r.repo.FindById(roleID)
	if err != nil {
		return fmt.Errorf("failed to delete role: %w", err)
	}

	return r.repo.Delete(roleID)
}
