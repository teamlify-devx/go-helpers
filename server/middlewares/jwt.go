package middlewares

import (
	"github.com/gofiber/fiber/v2"
	cfg "github.com/spf13/viper"
	"github.com/teamlify-devx/go-helpers/utils/jwt"
)

// JwtProtect Middleware JWT function
func JwtProtect() func(*fiber.Ctx) error {
	return jwt.New(jwt.Config{
		SigningKey:   jwt.SigningKey{Key: []byte(cfg.GetString("APP_SECRET"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error":   true,
			"message": "Missing or malformed JWT",
			"data":    nil,
		})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error":   true,
			"message": "Invalid or expired JWT",
			"data":    nil,
		})
	}
}
