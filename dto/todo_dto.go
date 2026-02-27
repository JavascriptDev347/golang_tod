package dto

import "time"

type CreateTodoRequest struct {
	Title       string     `json:"title" binding:"required,min=4,max=255"`
	Description string     `json:"description"`
	Priority    string     `json:"priority" binding:"omitempty,oneof=low medium high"`
	DueDate     *string `json:"due_date"`
	CategoryID  *uint      `json:"category_id"`
}

type UpdateTodoRequest struct {
	Title       string     `json:"title" binding:"omitempty,min=2,max=255"`
	Description string     `json:"description"`
	Status      string     `json:"status" binding:"omitempty,oneof=pending in_progress completed"`
	Priority    string     `json:"priority" binding:"omitempty,oneof=low medium high"`
	DueDate     *string `json:"due_date"`
	CategoryID  *uint      `json:"category_id"`
}

type TodoFilter struct {
	Search   string `form:"search"`
	Status   string `form:"status" binding:"omitempty,oneof=pending in_progress completed"`
	Priority string `form:"priority" binding:"omitempty,oneof=low medium high"`
	SortBy   string `form:"sort_by"`
	SortDir  string `form:"sort_dir"`
}

type TodoResponse struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      string            `json:"status"`
	Priority    string            `json:"priority"`
	DueDate     *time.Time        `json:"due_date"`
	UserID      uint              `json:"user_id"`
	CategoryID  *uint             `json:"category_id"`
	Category    *CategoryResponse `json:"category,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}
