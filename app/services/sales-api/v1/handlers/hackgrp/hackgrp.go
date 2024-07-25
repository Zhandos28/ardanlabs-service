package hackgrp

import (
	"context"
	"errors"
	"github.com/Zhandos28/ardanlabs-service/business/web/v1/response"
	"github.com/Zhandos28/ardanlabs-service/foundation/web"
	"math/rand"
	"net/http"
)

func Hack(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Int() % 4; n == 0 {
		return response.NewError(errors.New("TRUST ERROR"), http.StatusBadRequest)
	}
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
