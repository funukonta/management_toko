package models

import (
	"github.com/funukonta/management_toko/pkg/common"
	"gorm.io/gorm"
)

type Users struct {
	common.ModelID
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

func (u *Users) TableName() string {
	return "users"
}

func (u *Users) BeforeCreate(db *gorm.DB) error {
	u.SetUUID("usr")
	u.ModelID.BeforeCreate(db)
	return nil
}

type UserCondition struct {
	Where UserWhere
}

func (w *UserWhere) ImpelementWhere(db *gorm.DB) {
	if w.ID != "" {
		db.Where("id = ?", w.ID)
	}

	if w.Username != "" {
		db.Where("username = ?", w.Username)
	}
}

type UserWhere struct {
	ID       string
	Username string
}
