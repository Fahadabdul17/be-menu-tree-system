package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateMenuRequest struct {
	Name     string     `json:"name" binding:"required"`
	ParentID *uuid.UUID `json:"parent_id"`
	Order    int        `json:"order"`
}

type UpdateMenuRequest struct {
	Name     *string    `json:"name"`
	ParentID *uuid.UUID `json:"parent_id"`
	Order    *int       `json:"order"`
}

type MoveMenuRequest struct {
	ParentID *uuid.UUID `json:"parent_id"`
}

type ReorderMenuRequest struct {
	Order int `json:"order"`
}

type MenuResponse struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	ParentID  *uuid.UUID     `json:"parent_id"`
	Order     int            `json:"order"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	Children  []MenuResponse `json:"children"`
}
