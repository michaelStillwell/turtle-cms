package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/michaelStillwell/turtle-cms/views/partials"
)

func redirectLogin(ctx echo.Context) error {
	return ctx.Redirect(http.StatusTemporaryRedirect, "/auth/login")
}

func redirectNotFound(ctx echo.Context) error {
	return ctx.Redirect(http.StatusTemporaryRedirect, "/not-found")
}

func view(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func page(c echo.Context, title string, cmp templ.Component) error {
	tenant := c.Param("tenant")
	fmt.Println("tenant", tenant)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	ctx := templ.WithChildren(c.Request().Context(), cmp)
	return partials.Layout(title, tenant).Render(ctx, c.Response().Writer)
}
