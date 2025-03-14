package main

import (
	"time"

	"github.com/chhz0/gokit/pkg/server"
	"github.com/chhz0/gokit/pkg/server/engines"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

func useGin() *engines.GinEnginWrapper {
	g := gin.Default()
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return engines.Gin(g)
}

func useEcho() *engines.EchoEngineWrapper {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	return engines.Echo(e)
}

func main() {
	engine := useEcho()
	s := server.NewHttp(&server.HttpConfig{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}, engine)

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
