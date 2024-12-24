package handlers

import (
	"github.com/SwanHtetAungPhyo/esfor/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func BaseHandle(action func(*fiber.Ctx) (interface{}, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := action(c)
		if err != nil {
			return utils.JsonResp(c, fiber.StatusInternalServerError, err.Error())
		}
		return utils.JsonResp(c, fiber.StatusOK, "Success", data)
	}
}
