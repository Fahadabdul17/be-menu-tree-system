package repository

import (
	"be-menu-tree-system/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu *model.Menu) error
	GetAll() ([]model.Menu, error)
	GetByID(id uuid.UUID) (*model.Menu, error)
	Update(menu *model.Menu) error
	Delete(id uuid.UUID) error
	GetChildren(parentID uuid.UUID) ([]model.Menu, error)
	Move(id uuid.UUID, parentID *uuid.UUID) error
	UpdateOrder(id uuid.UUID, order int) error
	Exists(id uuid.UUID) (bool, error)
	GetDescendants(id uuid.UUID) ([]model.Menu, error)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db}
}

// Create saves a new menu
func (r *menuRepository) Create(menu *model.Menu) error {
	return r.db.Create(menu).Error
}

// GetAll returns all menus ordered by 'order'
func (r *menuRepository) GetAll() ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Order("\"order\" ASC").Find(&menus).Error
	return menus, err
}

// GetByID finds a menu by its ID
func (r *menuRepository) GetByID(id uuid.UUID) (*model.Menu, error) {
	var menu model.Menu
	err := r.db.First(&menu, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuRepository) Update(menu *model.Menu) error {
	return r.db.Save(menu).Error
}

// Delete removes a menu and all descendants using CTE
func (r *menuRepository) Delete(id uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var ids []uuid.UUID
		query := `
			WITH RECURSIVE descendants AS (
				SELECT id FROM menus WHERE id = ?
				UNION ALL
				SELECT m.id FROM menus m
				JOIN descendants d ON m.parent_id = d.id
			)
			SELECT id FROM descendants
		`
		if err := tx.Raw(query, id).Scan(&ids).Error; err != nil {
			return err
		}

		if len(ids) == 0 {
			return gorm.ErrRecordNotFound
		}

		return tx.Delete(&model.Menu{}, "id IN ?", ids).Error
	})
}

func (r *menuRepository) GetChildren(parentID uuid.UUID) ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Where("parent_id = ?", parentID).Order("\"order\" ASC").Find(&menus).Error
	return menus, err
}

func (r *menuRepository) Move(id uuid.UUID, parentID *uuid.UUID) error {
	return r.db.Model(&model.Menu{}).Where("id = ?", id).Update("parent_id", parentID).Error
}

func (r *menuRepository) UpdateOrder(id uuid.UUID, order int) error {
	return r.db.Model(&model.Menu{}).Where("id = ?", id).Update("order", order).Error
}

func (r *menuRepository) Exists(id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&model.Menu{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}

// GetDescendants returns all sub-menus recursively
func (r *menuRepository) GetDescendants(id uuid.UUID) ([]model.Menu, error) {
	var menus []model.Menu
	query := `
		WITH RECURSIVE descendants AS (
			SELECT * FROM menus WHERE id = ?
			UNION ALL
			SELECT m.* FROM menus m
			JOIN descendants d ON m.parent_id = d.id
		)
		SELECT * FROM descendants WHERE id != ?
	`
	err := r.db.Raw(query, id, id).Scan(&menus).Error
	return menus, err
}

