package hackgrp

import (
	"github.com/Zhandos28/ardanlabs-service/business/web/v1/auth"
	"github.com/Zhandos28/ardanlabs-service/business/web/v1/mid"
	"github.com/Zhandos28/ardanlabs-service/foundation/web"
	"net/http"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Auth *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Auth)
	ruleAdmin := mid.Authorize(cfg.Auth, auth.RuleAdminOnly)

	app.Handle(http.MethodGet, version, "/hack", Hack)
	app.Handle(http.MethodGet, version, "/hackauth", Hack, authen, ruleAdmin)
}
