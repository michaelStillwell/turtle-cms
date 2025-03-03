package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/michaelStillwell/turtle-cms/views/pages"
)

type (
	CollectionEditorRepo interface {
	}
)

func NewCollectionEditorHandler(repo CollectionEditorRepo) *CollectionEditorHandler {
	return &CollectionEditorHandler{
		repo: repo,
	}
}

type CollectionEditorHandler struct {
	repo CollectionEditorRepo
}

func (h *CollectionEditorHandler) handleView(ctx echo.Context) error {
	return page(ctx, "Collection Editor", pages.CollectionEditor())
}
