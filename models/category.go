package models

import "time"

type Category struct {
	ID          uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"type:varchar(100);not null;unique" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
	Color       string     `gorm:"type:varchar(50)" json:"color"`
	Image       string     `gorm:"type:varchar(255)" json:"image"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"default:null" json:"deleted_at,omitempty"`
}
