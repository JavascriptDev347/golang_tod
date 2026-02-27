package models

import "time"

type TodoStatus string
type TodoPriority string

const (
	StatusPending    TodoStatus = "pending"
	StatusInProgress TodoStatus = "in_progress"
	StatusCompleted  TodoStatus = "completed"
)

const (
	PriorityLow    TodoPriority = "low"
	PriorityMedium TodoPriority = "medium"
	PriorityHigh   TodoPriority = "high"
)

type Todo struct {
	ID          uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string       `gorm:"type:varchar(255);not null" json:"title"`
	Description string       `gorm:"type:text" json:"description"`
	Status      TodoStatus   `gorm:"type:todo_status;not null;default:'pending'" json:"status"`
	Priority    TodoPriority `gorm:"type:todo_priority;not null;default:'medium'" json:"priority"`
	DueDate     *time.Time   `gorm:"default:null" json:"due_date"`
	UserID      uint         `gorm:"not null" json:"user_id"`
	CategoryID  *uint        `gorm:"default:null" json:"category_id"`
	CreatedAt   time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time   `gorm:"default:null" json:"deleted_at,omitempty"`

	// Relations
	User     User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Category *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}
