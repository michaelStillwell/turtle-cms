package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/michaelStillwell/turtle-cms/db"
	"github.com/michaelStillwell/turtle-cms/handlers"
)

func main() {
	var (
		u         = os.Getenv("DATABASE_URL")
		store     = db.MustNew(u)
		h         = handlers.New(store)
		logConfig = middleware.RequestLoggerConfig{
			LogStatus: true,
			LogURI:    true,
			BeforeNextFunc: func(c echo.Context) {
				// c.Set("customValueFromContext", 42)
			},
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				fmt.Printf("REQUEST: uri: %v, status: %v\n", v.URI, v.Status)
				return nil
			},
		}
	)

	logFile, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("cannot open file: %v", err))
	}
	defer logFile.Close()

	e := echo.New()
	e.Logger.SetOutput(io.MultiWriter(os.Stdout, logFile))
	e.Use(middleware.CORS())
	e.Use(middleware.RequestLoggerWithConfig(logConfig))

	handlers.Setup(e, h)
	// // e.GET("/swagger/*", echoSwagger.WrapHandler)

	if data, err := json.MarshalIndent(e.Routes(), "", "  "); err == nil {
		os.WriteFile("internal/routes.json", data, 0644)
	}

	e.Logger.Fatal(e.Start(":6969"))
}
