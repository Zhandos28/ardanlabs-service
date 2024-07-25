package checkgrp

import (
	"github.com/Zhandos28/ardanlabs-service/foundation/logger"
	"github.com/Zhandos28/ardanlabs-service/foundation/web"
	"net/http"

	"github.com/jmoiron/sqlx"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
	DB    *sqlx.DB
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	hdl := New(cfg.Build, cfg.Log, cfg.DB)
	app.HandleNoMiddleware(http.MethodGet, version, "/readiness", hdl.Readiness)
	app.HandleNoMiddleware(http.MethodGet, version, "/liveness", hdl.Liveness)
}
