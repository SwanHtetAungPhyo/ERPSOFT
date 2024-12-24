package middlewares

import (
	"fmt"
	"time"

	"github.com/SwanHtetAungPhyo/esfor/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

var SecretKey = "p0EKBU4NvrdYGqnYgCNzvdZQHNrZiUjj4jQ7ZVrgRj5Ymg5S"

func JwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing auth header",
			})
		}

		tokenString := authHeader[len("Bearer "):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(SecretKey), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		c.Locals("claims", claims)

		return c.Next()
	}
}

func LoggingMiddleware() fiber.Handler {
	logger := utils.GetLogger()

	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()
		logger.Info("Request",
			zap.String("method", c.Method()),
			zap.String("url", c.OriginalURL()),
			zap.String("path", c.Path()),
			zap.String("protocol", c.Protocol()),
			zap.String("user_agent", c.Get("User-Agent")),
			zap.String("referer", c.Get("Referer")),
			zap.String("request_id", c.Get("X-Request-ID")),
			zap.String("remote_ip", c.IP()),
			zap.String("host", c.Hostname()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", time.Since(start)),
			zap.String("client_ip", c.IP()),
		)

		if err != nil {
			logger.Error("Error processing request", zap.Error(err))
			return err
		}

		return nil
	}
}
