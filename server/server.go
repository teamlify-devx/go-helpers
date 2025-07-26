package server

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/teamlify-devx/go-helpers/logger"
)

const (
	certFile        = "ssl/server.crt"
	keyFile         = "ssl/server.pem"
	maxHeaderBytes  = 1 << 20
	gzipLevel       = 5
	stackSize       = 1 << 10 // 1 KB
	csrfTokenHeader = "X-CSRF-Token"
	bodyLimit       = "2M"
)

// server
type server struct {
	ctx    *context.Context
	logger logger.Logger
	fiber  *fiber.App
}

// NewServer constructor
func NewServer(ctx *context.Context, logger logger.Logger) *server {
	return &server{
		ctx:    ctx,
		logger: logger,
	}
}
