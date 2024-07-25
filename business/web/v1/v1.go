package v1

import (
	"github.com/Zhandos28/ardanlabs-service/business/web/v1/auth"
	"github.com/Zhandos28/ardanlabs-service/business/web/v1/mid"
	"github.com/Zhandos28/ardanlabs-service/foundation/logger"
	"github.com/Zhandos28/ardanlabs-service/foundation/web"
	"os"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *logger.Logger
	Auth     *auth.Auth
}

// RouteAdder defines behavior that sets the routes to bind for an instance
// of the service.
type RouteAdder interface {
	Add(app *web.App, cfg APIMuxConfig)
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig, routeAdder RouteAdder) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Metrics(), mid.Panics())
	routeAdder.Add(app, cfg)

	return app
}
