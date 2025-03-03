package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/michaelStillwell/turtle-cms/internal/repository"
	"github.com/michaelStillwell/turtle-cms/pkg"
	"github.com/michaelStillwell/turtle-cms/views/pages"
)

type (
	loginRequest struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	registerRequest struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	AuthRepo interface {
		GetUserByEmail(context.Context, string) (repository.User, error)
		CreateUser(context.Context, repository.CreateUserParams) (repository.User, error)
	}
)

func NewAuthHandler(repo AuthRepo) *AuthHandler {
	return &AuthHandler{
		repo: repo,
	}
}

type AuthHandler struct {
	repo AuthRepo
}

func (h *AuthHandler) handleView(ctx echo.Context) error {
	c, _ := ctx.Cookie(pkg.GetCookieName())
	if c.Valid() == nil {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return view(ctx, pages.Login())
}

func (h *AuthHandler) handleLogin(ctx echo.Context) error {
	req := &loginRequest{}
	if err := ctx.Bind(req); err != nil {
		return ctx.HTML(http.StatusBadRequest, "<h1>youre dumb, wrong params</h1>")
	}

	if len(req.Email) <= 0 || len(req.Password) <= 0 {
		return ctx.NoContent(http.StatusBadRequest)
	}

	user, err := h.repo.GetUserByEmail(ctx.Request().Context(), req.Email)
	if err != nil {
		return view(ctx, pages.LoginErrorMsg("not found"))
	}

	if !pkg.VerifyPassword(req.Password, user.Password) {
		return view(ctx, pages.LoginErrorMsg("not found"))
	}

	// TODO: set cookie
	if err := pkg.GenerateTokenAndCookie(ctx, int(user.UserID), "testing-daddy"); err != nil {
		fmt.Println(err)
		return ctx.NoContent(http.StatusInternalServerError)
	}

	ctx.Response().Header().Set("HX-Redirect", "/")
	return ctx.NoContent(http.StatusOK)
}

func (h *AuthHandler) handleLogout(ctx echo.Context) error {
	pkg.RemoveCookie(ctx)
	return ctx.Redirect(http.StatusTemporaryRedirect, "/auth/login")
}

func (h *AuthHandler) handleRegister(ctx echo.Context) error {
	req := &registerRequest{}
	if err := ctx.Bind(req); err != nil {
		return ctx.NoContent(http.StatusBadRequest)
	}

	if len(req.Email) <= 0 || len(req.Password) <= 0 {
		return ctx.NoContent(http.StatusBadRequest)
	}

	hash, err := pkg.HashPassword(req.Password)
	if err != nil {
		return ctx.NoContent(http.StatusInternalServerError)
	}

	if _, err := h.repo.CreateUser(ctx.Request().Context(), repository.CreateUserParams{
		Email:     req.Email,
		Password:  hash,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}); err != nil {
		return ctx.JSON(http.StatusInternalServerError, `{"message":"meow"}`)
	}

	return ctx.NoContent(http.StatusOK)
}
