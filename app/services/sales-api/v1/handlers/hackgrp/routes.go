package hackgrp

import (
	v1 "github.com/Zhandos28/ardanlabs-service/business/web/v1"
	"github.com/Zhandos28/ardanlabs-service/foundation/web"
	"net/http"
)

func Routes(app *web.App, apiCfg v1.APIMuxConfig) {
	//log := mid.Logger(apiCfg.Log)
	app.Handle(http.MethodGet, "", "/hack", Hack)
}
