package service

import (
	"be-menu-tree-system/internal/dto"
	"be-menu-tree-system/internal/model"
	"be-menu-tree-system/internal/repository"
	"errors"

	"github.com/google/uuid"
)

type MenuService interface {
	CreateMenu(req dto.CreateMenuRequest) (*dto.MenuResponse, error)
	GetMenuTree() ([]dto.MenuResponse, error)
	GetMenuByID(id uuid.UUID) (*dto.MenuResponse, error)
	UpdateMenu(id uuid.UUID, req dto.UpdateMenuRequest) (*dto.MenuResponse, error)
	DeleteMenu(id uuid.UUID) error
	MoveMenu(id uuid.UUID, parentID *uuid.UUID) error
	ReorderMenu(id uuid.UUID, order int) error
}

type menuService struct {
	repo repository.MenuRepository
}

func NewMenuService(repo repository.MenuRepository) MenuService {
	return &menuService{repo}
}

func (s *menuService) CreateMenu(req dto.CreateMenuRequest) (*dto.MenuResponse, error) {
	if req.ParentID != nil {
		exists, err := s.repo.Exists(*req.ParentID)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, errors.New("parent menu not found")
		}
	}

	menu := &model.Menu{
		Name:     req.Name,
		ParentID: req.ParentID,
		Order:    req.Order,
	}

	if err := s.repo.Create(menu); err != nil {
		return nil, err
	}

	return s.toResponse(menu), nil
}

func (s *menuService) GetMenuTree() ([]dto.MenuResponse, error) {
	menus, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return buildTree(menus), nil
}

func (s *menuService) GetMenuByID(id uuid.UUID) (*dto.MenuResponse, error) {
	menu, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return s.toResponse(menu), nil
}

func (s *menuService) UpdateMenu(id uuid.UUID, req dto.UpdateMenuRequest) (*dto.MenuResponse, error) {
	menu, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		menu.Name = *req.Name
	}
	if req.Order != nil {
		menu.Order = *req.Order
	}
	if req.ParentID != nil {
		if err := s.checkCircularDependency(id, *req.ParentID); err != nil {
			return nil, err
		}
		menu.ParentID = req.ParentID
	}

	if err := s.repo.Update(menu); err != nil {
		return nil, err
	}

	return s.toResponse(menu), nil
}

func (s *menuService) DeleteMenu(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *menuService) MoveMenu(id uuid.UUID, parentID *uuid.UUID) error {
	if parentID != nil {
		if id == *parentID {
			return errors.New("cannot move menu to itself")
		}

		exists, err := s.repo.Exists(*parentID)
		if err != nil {
			return err
		}
		if !exists {
			return errors.New("target parent menu not found")
		}

		if err := s.checkCircularDependency(id, *parentID); err != nil {
			return err
		}
	}

	return s.repo.Move(id, parentID)
}

func (s *menuService) ReorderMenu(id uuid.UUID, order int) error {
	return s.repo.UpdateOrder(id, order)
}

// checkCircularDependency ensures a menu isn't moved into its own subtree
func (s *menuService) checkCircularDependency(id uuid.UUID, newParentID uuid.UUID) error {
	if id == newParentID {
		return errors.New("menu cannot be its own parent")
	}

	descendants, err := s.repo.GetDescendants(id)
	if err != nil {
		return err
	}

	for _, d := range descendants {
		if d.ID == newParentID {
			return errors.New("cannot move menu into its own subtree")
		}
	}

	return nil
}

func (s *menuService) toResponse(m *model.Menu) *dto.MenuResponse {
	return &dto.MenuResponse{
		ID:        m.ID,
		Name:      m.Name,
		ParentID:  m.ParentID,
		Order:     m.Order,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		Children:  []dto.MenuResponse{},
	}
}

// buildTree converts flat menu list into hierarchical structure
func buildTree(menus []model.Menu) []dto.MenuResponse {
	// Map menus by their parent ID for quick lookup
	parentMap := make(map[uuid.UUID][]model.Menu)
	var rootMenus []model.Menu

	for _, m := range menus {
		if m.ParentID == nil {
			rootMenus = append(rootMenus, m)
		} else {
			parentMap[*m.ParentID] = append(parentMap[*m.ParentID], m)
		}
	}

	// Recursive function to assemble nodes
	var assemble func(m model.Menu) dto.MenuResponse
	assemble = func(m model.Menu) dto.MenuResponse {
		node := dto.MenuResponse{
			ID:        m.ID,
			Name:      m.Name,
			ParentID:  m.ParentID,
			Order:     m.Order,
			CreatedAt: m.CreatedAt,
			UpdatedAt: m.UpdatedAt,
			Children:  []dto.MenuResponse{},
		}

		// Fill children recursively
		for _, child := range parentMap[m.ID] {
			node.Children = append(node.Children, assemble(child))
		}
		return node
	}

	var finalTree []dto.MenuResponse
	for _, root := range rootMenus {
		finalTree = append(finalTree, assemble(root))
	}

	return finalTree
}



