package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/michaelStillwell/turtle-cms/views/pages"
)

type (
	SettingsRepo interface {
	}
)

func NewSettingsHandler(repo SettingsRepo) *SettingsHandler {
	return &SettingsHandler{
		repo: repo,
	}
}

type SettingsHandler struct {
	repo SettingsRepo
}

func (h *SettingsHandler) handleView(ctx echo.Context) error {
	return page(ctx, "Settings Home", pages.Settings())
}
