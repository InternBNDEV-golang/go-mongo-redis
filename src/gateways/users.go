package gateways

import (
	"go-mongo-redis/domain/entities"

	"github.com/gofiber/fiber/v2"
)

func (h HTTPGateway) GetAllUserData(ctx *fiber.Ctx) error {

	data, err := h.UserService.GetAllUser()
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get all users data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}
func (h HTTPGateway) GetAllUserDataRedis(ctx *fiber.Ctx) error {

	data, err := h.UserService.GetAllUserRedis()
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot get all users data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "success", Data: data})
}

func (h HTTPGateway) SetAllUserDataRedis(ctx *fiber.Ctx) error {
	_, err := h.UserService.SetAllUserRedis()
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(entities.ResponseModel{Message: "cannot set all users data"})
	}
	return ctx.Status(fiber.StatusOK).JSON(entities.ResponseModel{Message: "set redis success"})
}
