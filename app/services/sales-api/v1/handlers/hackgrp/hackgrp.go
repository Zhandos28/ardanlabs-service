package hackgrp

import (
	"context"
	"encoding/json"
	"github.com/Zhandos28/ardanlabs-service/foundation/web"
	"net/http"
)

func Hack(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "OK",
	}

	if err := json.NewEncoder(w).Encode(status); err != nil {
		return err
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
