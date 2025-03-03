package handlers

// import (
// 	"context"

// 	"github.com/labstack/echo/v4"
// 	"github.com/michaelStillwell/turtle-cms/internal/repository"
// 	"github.com/michaelStillwell/turtle-cms/views/pages"
// )

// type IndexRepo interface {
// 	GetUser(context.Context, int64) (repository.User, error)
// }

// func NewIndexHandler(repo IndexRepo) *IndexHandler {
// 	return &IndexHandler{
// 		repo: repo,
// 	}
// }

// type IndexHandler struct {
// 	repo IndexRepo
// }

// func (h *IndexHandler) View(ctx echo.Context) error {
// 	userId, ok := ctx.Get(GetUserIDKey()).(int64)
// 	if !ok {
// 		return redirectNotFound(ctx)
// 	}

// 	user, err := h.repo.GetUser(ctx.Request().Context(), userId)
// 	if err != nil {
// 		return redirectNotFound(ctx)
// 	}

// 	return page(ctx, "Home", pages.Index(user.Email))
// }
