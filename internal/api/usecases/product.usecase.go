package usecases

import (
	"github.com/funukonta/management_toko/internal/api/constants"
	"github.com/funukonta/management_toko/internal/api/dtos"
	"github.com/funukonta/management_toko/internal/api/models"
	"github.com/funukonta/management_toko/internal/api/repositories"
	"github.com/funukonta/management_toko/pkg/common"
)

type UsecaseProduct struct {
	repo repositories.RepoProduct
}

func NewUsecaseProduct(repo repositories.RepoProduct) *UsecaseProduct {
	return &UsecaseProduct{repo: repo}
}

func (uc *UsecaseProduct) AddProduct(req dtos.AddProductReq) (*common.RespData, error) {

	if constants.MapToLabelProductType()[req.Type] == "" {
		return nil, &common.CustomError{Message: "Type not valid"}
	}

	product := &models.Products{
		Name:     req.Name,
		Type:     req.Type,
		Price:    req.Price,
		Quantity: req.Quantity,
	}

	err := uc.repo.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return &common.RespData{
		Message: "Product Created",
		Data:    product,
	}, nil

}

func (uc *UsecaseProduct) GetAllProduct() (*common.RespData, error) {
	products, err := uc.repo.GetProducts(models.ProductCondition{})
	if err != nil {
		return nil, err
	}

	return &common.RespData{
		Message: "Products List",
		Data:    products,
	}, nil
}

func (uc *UsecaseProduct) GetDetailProduct(id string) (*common.RespData, error) {
	product, err := uc.repo.GetProduct(models.ProductCondition{
		Where: models.ProductWhere{
			ID: id,
		},
		Preload: models.ProductPreload{
			Type: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &common.RespData{
		Message: "Products List",
		Data:    product,
	}, nil
}

func (uc *UsecaseProduct) UpdateProduct(id string, req dtos.AddProductReq) (*common.RespData, error) {
	updatedProduct := &models.Products{
		Name:     req.Name,
		Type:     req.Type,
		Price:    req.Price,
		Quantity: req.Quantity,
	}
	err := uc.repo.UpdateProduct(updatedProduct, models.ProductWhere{
		ID: id,
	})
	if err != nil {
		return nil, err
	}

	return &common.RespData{
		Message: "Product Updated",
		Data:    updatedProduct,
	}, nil
}

func (uc *UsecaseProduct) DeleteProduct(id string) (*common.RespData, error) {
	err := uc.repo.DeleteProduct(id)
	if err != nil {
		return nil, err
	}

	resp := &common.RespData{
		Message: "Product Deleted",
	}

	return resp, nil
}
