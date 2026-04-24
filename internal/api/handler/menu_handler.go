package handler

import (
	"be-menu-tree-system/internal/dto"
	"be-menu-tree-system/internal/service"
	"be-menu-tree-system/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MenuHandler struct {
	service service.MenuService
}

func NewMenuHandler(service service.MenuService) *MenuHandler {
	return &MenuHandler{service}
}

// CreateMenu adds a new menu item
// @Summary Create menu
// @Tags menus
// @Param menu body dto.CreateMenuRequest true "payload"
// @Success 201 {object} response.Response
// @Router /api/menus [post]
func (h *MenuHandler) CreateMenu(c *gin.Context) {
	var req dto.CreateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request payload", err.Error())
		return
	}

	res, err := h.service.CreateMenu(req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusCreated, "Menu created", res)
}

// GetMenuTree returns hierarchical menu
// @Summary Get menu tree
// @Tags menus
// @Success 200 {object} response.Response
// @Router /api/menus [get]
func (h *MenuHandler) GetMenuTree(c *gin.Context) {
	res, err := h.service.GetMenuTree()
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Menus retrieved", res)
}

// GetMenuByID returns a single menu item
// @Summary Get menu by ID
// @Tags menus
// @Param id path string true "id"
// @Success 200 {object} response.Response
// @Router /api/menus/{id} [get]
func (h *MenuHandler) GetMenuByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID format", nil)
		return
	}

	res, err := h.service.GetMenuByID(id)
	if err != nil {
		response.NotFound(c, "Menu not found")
		return
	}

	response.Success(c, http.StatusOK, "Menu retrieved", res)
}

// UpdateMenu modifies a menu
// @Summary Update menu
// @Tags menus
// @Param id path string true "id"
// @Param menu body dto.UpdateMenuRequest true "payload"
// @Success 200 {object} response.Response
// @Router /api/menus/{id} [put]
func (h *MenuHandler) UpdateMenu(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID format", nil)
		return
	}

	var req dto.UpdateMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request payload", err.Error())
		return
	}

	res, err := h.service.UpdateMenu(id, req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Menu updated", res)
}

// DeleteMenu removes menu and descendants
// @Summary Delete menu
// @Tags menus
// @Param id path string true "id"
// @Success 200 {object} response.Response
// @Router /api/menus/{id} [delete]
func (h *MenuHandler) DeleteMenu(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID format", nil)
		return
	}

	if err := h.service.DeleteMenu(id); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Menu deleted", nil)
}

// MoveMenu changes a menu's parent
func (h *MenuHandler) MoveMenu(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID format", nil)
		return
	}

	var req dto.MoveMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request payload", err.Error())
		return
	}

	if err := h.service.MoveMenu(id, req.ParentID); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Menu moved", nil)
}

// ReorderMenu changes a menu's order within its level
func (h *MenuHandler) ReorderMenu(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid ID format", nil)
		return
	}

	var req dto.ReorderMenuRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request payload", err.Error())
		return
	}

	if err := h.service.ReorderMenu(id, req.Order); err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Success(c, http.StatusOK, "Menu reordered", nil)
}
