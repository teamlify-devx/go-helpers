package jwt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"main/config"
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

type TokenClaim struct {
	UserID      int64  `json:"id" validate:"required"`
	ParentID    int64  `json:"parent_id" validate:"required"`
	UserType    int    `json:"type" validate:"required"`
	Title       string `json:"title" validate:"required"`
	HasOrg      bool   `json:"has_org" validate:"required"`
	NdaSigned   bool   `json:"nda_signed" validate:"required"`
	AccessToken string `json:"access_token,omitempty"`
	Exp         int64  `json:"exp,omitempty"`
	Iat         int64  `json:"iat,omitempty"`
	jwt.RegisteredClaims
}

func GenerateToken(cfg *config.Config, claims TokenClaim, expTime int64) (string, error) {
	signingKey := []byte(cfg.Server.APP_SECRET)
	claims.Iat = time.Now().Unix()
	claims.Exp = expTime
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}
