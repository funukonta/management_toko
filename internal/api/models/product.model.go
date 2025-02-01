package models

import (
	"github.com/funukonta/management_toko/pkg/common"
	"gorm.io/gorm"
)

type Products struct {
	common.ModelID
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`

	ProductType ProductType `json:"product_type" gorm:"foreignKey:Type"`
}

func (p *Products) TableName() string {
	return "products"
}

func (p *Products) BeforeCreate(db *gorm.DB) error {
	p.SetUUID("prd")
	p.ModelID.BeforeCreate(db)
	return nil
}

type ProductType struct {
	common.ModelCode
}

func (p *ProductType) TableName() string {
	return "productypes"
}

type ProductCondition struct {
	Where   ProductWhere
	Preload ProductPreload
}

type ProductWhere struct {
	Types []string
	Type  string

	ID  string
	IDs []string
}

func (pw *ProductWhere) ImplementWhere(db *gorm.DB) {
	if len(pw.Types) > 0 {
		db.Where("", pw.Types)
	}
	if pw.Type != "" {
		db.Where("", pw.Type)
	}

	if pw.ID != "" {
		db.Where("id = ?", pw.ID)
	}
	if len(pw.IDs) > 0 {
		db.Where("id in ?", pw.IDs)
	}
}

type ProductPreload struct {
	Type bool
}

func (p *ProductPreload) ImplementPreload(db *gorm.DB) {
	if p.Type {
		db.Preload("ProductType")
	}
}
