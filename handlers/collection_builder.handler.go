package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/michaelStillwell/turtle-cms/views/pages"
)

type (
	CollectionBuilderRepo interface {
	}
)

func NewCollectionBuilderHandler(repo CollectionBuilderRepo) *CollectionBuilderHandler {
	return &CollectionBuilderHandler{
		repo: repo,
	}
}

type CollectionBuilderHandler struct {
	repo CollectionBuilderRepo
}

func (h *CollectionBuilderHandler) handleView(ctx echo.Context) error {
	tenant := ctx.Param("tenant")

	return page(ctx, "CollectionBuilder", pages.CollectionBuilder(tenant))
}
