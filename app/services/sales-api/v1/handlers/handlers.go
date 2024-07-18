package handlers

import (
	"github.com/Zhandos28/ardanlabs-service/app/services/sales-api/v1/handlers/hackgrp"
	v1 "github.com/Zhandos28/ardanlabs-service/business/web/v1"
	"github.com/dimfeld/httptreemux/v5"
)

type Routes struct{}

// Add implements the RouterAdder interface to add all routes.
func (Routes) Add(mux *httptreemux.ContextMux, apiCfg v1.APIMuxConfig) {
	hackgrp.Routes(mux)
}
