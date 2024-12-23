package handlers

import (
	"github.com/funukonta/management_toko/internal/api/dtos"
	"github.com/funukonta/management_toko/internal/api/usecases"
	"github.com/funukonta/management_toko/pkg/common"
	"github.com/gofiber/fiber/v2"
)

type HanlderUser struct {
	userUc *usecases.UsecaseUser
}

func NewHandlerUser(useruc *usecases.UsecaseUser) *HanlderUser {
	return &HanlderUser{userUc: useruc}
}

func (h *HanlderUser) GetUser(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "GetUser endpoint is working!",
	})
}

func (h *HanlderUser) Login(c *fiber.Ctx) error {
	var req dtos.LoginReq
	err := c.BodyParser(&req)
	if err != nil {
		return common.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	resp, err := h.userUc.Login(req)
	if err != nil {
		return common.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	c.Status(fiber.StatusOK).JSON(resp)
	return nil
}

func (h *HanlderUser) Register(c *fiber.Ctx) error {
	var req dtos.RegisterReq
	err := c.BodyParser(&req)
	if err != nil {
		return common.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	resp, err := h.userUc.Register(req)
	if err != nil {
		return common.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	c.Status(fiber.StatusOK).JSON(resp)
	return nil
}
