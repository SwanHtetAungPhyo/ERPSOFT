package utils

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type (
	EnvVars struct {
		PORT string
		DSN  string
	}

	ApiResponse struct {
		StatusCode int    `json:"status"`
		Message    string `json:"message"`
		Body       any    `json:"body,omitempty"`
		MetaData   any    `json:"meta_data,omitempty"`
	}
)

func NewEnvConfig() *EnvVars {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	return &EnvVars{
		PORT: os.Getenv("PORT"),
		DSN:  os.Getenv("DSN"),
	}
}

func JsonResp(c *fiber.Ctx, status int, message string, data ...any) error {
	response := &ApiResponse{
		StatusCode: status,
		Message:    message,
	}

	if len(data) > 0 {
		response.Body = data[0]
	}
	if len(data) > 1 {
		response.MetaData = data[1]
	}

	return c.Status(status).JSON(response)
}
