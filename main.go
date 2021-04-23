package main

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, status=${status}, uri=${uri}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
	}))

	api := e.Group("/api")
	Router(api)
	
	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize: 1048576,
		IdleTimeout: 10 * time.Second,
	}

	e.Logger.Fatal(e.StartH2CServer(":3000", s))
}