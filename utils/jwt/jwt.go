package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	cfg "github.com/spf13/viper"
	"reflect"
	"time"
)

var (
	defaultTokenLookup = "header:" + fiber.HeaderAuthorization
)

// New ...
func New(config ...Config) fiber.Handler {
	cfg := makeCfg(config)

	extractors := cfg.getExtractors()

	// Return middleware handler
	return func(c *fiber.Ctx) error {
		// Filter request to skip middleware
		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}
		var auth string
		var err error

		for _, extractor := range extractors {
			auth, err = extractor(c)
			if auth != "" && err == nil {
				break
			}
		}
		if err != nil {
			return cfg.ErrorHandler(c, err)
		}
		var token *jwt.Token

		if _, ok := cfg.Claims.(jwt.MapClaims); ok {
			token, err = jwt.Parse(auth, cfg.KeyFunc)
		} else {
			t := reflect.ValueOf(cfg.Claims).Type().Elem()
			claims := reflect.New(t).Interface().(jwt.Claims)
			token, err = jwt.ParseWithClaims(auth, claims, cfg.KeyFunc)
		}
		if err == nil && token.Valid {
			// Store user information from token into context.
			c.Locals(cfg.ContextKey, token)
			return cfg.SuccessHandler(c)
		}
		return cfg.ErrorHandler(c, err)
	}
}

/*
	GenerateToken generates a JWT token with the given claims and expiration time.

	Parameters:
	- claims: A map of claims to include in the token.
	- expTime: The expiration time for the token (Unix timestamp).

	Returns:
	- A signed JWT token as a string.
	- An error if the token generation fails.

Example payload

	claims := jwt.MapClaims{
		"user_id":    12345,
		"parent_id":  67890,
		"type":       1,
		"title":      "User Title",
		"has_org":    true,
		"nda_signed": false,
	}
*/
func GenerateToken(claims jwt.MapClaims, expTime int64) (string, error) {
	signingKey := []byte(cfg.GetString("Server.APP_SECRET"))

	claims["iat"] = time.Now().Unix()
	claims["exp"] = expTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}
