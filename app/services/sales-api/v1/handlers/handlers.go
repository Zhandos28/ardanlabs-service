package handlers

import (
	"github.com/Zhandos28/ardanlabs-service/app/services/sales-api/v1/handlers/hackgrp"
	v1 "github.com/Zhandos28/ardanlabs-service/business/web/v1"
	"github.com/Zhandos28/ardanlabs-service/foundation/web"
)

type Routes struct{}

// Add implements the RouterAdder interface to add all routes.
func (Routes) Add(app *web.App, apiCfg v1.APIMuxConfig) {
	hackgrp.Routes(app, apiCfg)
}