package handlers

import (
	"github.com/funukonta/management_toko/internal/api/dtos"
	"github.com/funukonta/management_toko/internal/api/usecases"
	"github.com/funukonta/management_toko/pkg/common"
	"github.com/gofiber/fiber/v2"
)

type HanlderProduct struct {
	ProductUc *usecases.UsecaseProduct
}

func NewHandlerProduct(Productuc *usecases.UsecaseProduct) *HanlderProduct {
	return &HanlderProduct{ProductUc: Productuc}
}

func (h *HanlderProduct) AddProduct(c *fiber.Ctx) error {
	var req dtos.AddProductReq
	err := c.BodyParser(&req)
	if err != nil {
		return common.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	resp, err := h.ProductUc.AddProduct(req)
	if err != nil {
		return common.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	c.Status(fiber.StatusOK).JSON(resp)
	return nil
}

func (h *HanlderProduct) GetAllProduct(c *fiber.Ctx) error {
	resp, err := h.ProductUc.GetAllProduct()
	if err != nil {
		return common.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	c.Status(fiber.StatusOK).JSON(resp)
	return nil
}

func (h *HanlderProduct) GetDetailProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.ProductUc.GetDetailProduct(id)
	if err != nil {
		return common.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	c.Status(fiber.StatusOK).JSON(resp)
	return nil
}

func (h *HanlderProduct) UpdateProduct(c *fiber.Ctx) error {
	var req dtos.AddProductReq
	err := c.BodyParser(&req)
	if err != nil {
		return common.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	id := c.Params("id")

	resp, err := h.ProductUc.UpdateProduct(id, req)
	if err != nil {
		return common.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	c.Status(fiber.StatusOK).JSON(resp)
	return nil
}

func (h *HanlderProduct) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	resp, err := h.ProductUc.DeleteProduct(id)
	if err != nil {
		return common.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	c.Status(fiber.StatusOK).JSON(resp)
	return nil
}
