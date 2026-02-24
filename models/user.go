package models

import "time"

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type User struct {
	ID        uint       `gorm:"primaryKey;autoIncrement;" json:"id"`
	Name      string     `gorm:"type:varchar(100);not null;" json:"name"`
	Email     string     `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password  string     `gorm:"type:varchar(255);not null" json:"-"`
	Role      UserRole   `gorm:"type:user_role;not null;default:'user'" json:"role"`
	DeletedAt *time.Time `gorm:"default:null" json:"deleted_at"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
