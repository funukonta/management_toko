package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID         string         `json:"id"`
	Username   string         `json:"username"`
	Password   string         `json:"password"`
	Name       string         `json:"name"`
	Role       string         `json:"role"`
	Status     string         `json:"status"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time      `json:"updated_at"`
	Deleted_at gorm.DeletedAt `json:"deleted_at"`
}
