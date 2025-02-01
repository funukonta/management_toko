package repositories

import (
	"github.com/funukonta/management_toko/internal/api/models"
	"gorm.io/gorm"
)

type RepoProduct interface {
	CreateProduct(prod *models.Products) error
	UpdateProduct(prod *models.Products, where models.ProductWhere) error
	DeleteProduct(id string) error
	GetProduct(cond models.ProductCondition) (*models.Products, error)
	GetProducts(cond models.ProductCondition) ([]models.Products, error)
}

type repoProduct struct {
	*gorm.DB
}

func NewRepoProduct(db *gorm.DB) RepoProduct {
	return &repoProduct{DB: db}
}

func (r *repoProduct) CreateProduct(prod *models.Products) error {
	return r.DB.Create(prod).Error
}

func (r *repoProduct) UpdateProduct(prod *models.Products, where models.ProductWhere) error {
	q := r.DB.Model(prod)
	where.ImplementWhere(q)
	q.Updates(prod)
	return q.Error
}

func (r *repoProduct) DeleteProduct(id string) error {
	q := r.DB.Delete(&models.Products{}, "id = ?", id)
	return q.Error
}

func (r *repoProduct) GetProduct(cond models.ProductCondition) (*models.Products, error) {
	product := &models.Products{}
	q := r.DB.Model(product)
	cond.Where.ImplementWhere(q)
	cond.Preload.ImplementPreload(q)

	err := q.First(product).Error

	return product, err
}

func (r *repoProduct) GetProducts(cond models.ProductCondition) ([]models.Products, error) {
	product := []models.Products{}
	q := r.DB.Model(product)
	cond.Where.ImplementWhere(q)
	cond.Preload.ImplementPreload(q)

	err := q.Find(product).Error

	return product, err
}
