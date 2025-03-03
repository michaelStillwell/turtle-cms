package handlers

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/michaelStillwell/turtle-cms/db"
	"github.com/michaelStillwell/turtle-cms/pkg"
)

func New(store *db.Store) *Handlers {
	return &Handlers{
		auth:              NewAuthHandler(store.Repo),
		collectionBuilder: NewCollectionBuilderHandler(store.Repo),
		collectionEditor:  NewCollectionEditorHandler(store.Repo),
	}
}

type Handlers struct {
	auth              *AuthHandler
	settings          *SettingsHandler
	collectionBuilder *CollectionBuilderHandler
	collectionEditor  *CollectionEditorHandler
}

func Setup(e *echo.Echo, h *Handlers) {
	var (
		config = echojwt.Config{
			Skipper: func(c echo.Context) bool {
				return strings.HasPrefix(c.Path(), "/auth") || strings.HasPrefix(c.Path(), "/api/register")
			},
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(pkg.JWTClaims)
			},
			SigningKey:  pkg.GetJwtKey(),
			TokenLookup: fmt.Sprintf("cookie:%s", pkg.GetCookieName()),
			ErrorHandler: func(c echo.Context, err error) error {
				return redirectLogin(c)
			},
		}
	)

	e.Use(echojwt.WithConfig(config))

	tenant := e.Group("/:tenant")
	tenant.GET("/collection-editor", h.collectionEditor.handleView)
	tenant.GET("/collection-builder", h.collectionBuilder.handleView)
	tenant.GET("/settings", h.settings.handleView)

	// admin := e.Group("/admin") // , WithAdmin)
	// e.GET("/", h.admin.View)

	auth := e.Group("/auth")
	auth.GET("/login", h.auth.handleView)
	auth.POST("/login", h.auth.handleLogin)
	auth.GET("/logout", h.auth.handleLogout)

	api := e.Group("/api")
	api.POST("/register", h.auth.handleRegister)
}
