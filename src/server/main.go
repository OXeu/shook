package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.POST("/admin/:name", func(c echo.Context) error {
		resp := add(c.Param("name"), c.FormValue("pwd"), c.FormValue("shell"))
		println(resp)
		return c.String(200, resp)
	})
	e.PUT("/admin", func(c echo.Context) error {
		resp := initServer(c.FormValue("token"))
		println(resp)
		return c.String(200, resp)
	})
	e.DELETE("/admin/:name", func(c echo.Context) error {
		resp := del(c.Param("name"))
		println(resp)
		return c.String(200, resp)
	})
	e.GET("/admin", func(c echo.Context) error {
		resp := ls()
		println(resp)
		return c.String(200, resp)
	})
	e.Any("/:name", func(c echo.Context) error {
		resp := run(c.Param("name"))
		println(resp)
		return c.String(200, resp)
	})
	e.Use(middleware.KeyAuthWithConfig(
		middleware.KeyAuthConfig{
			Skipper: func(c echo.Context) bool {
				if c.Path() == "/:name" {
					return true
				}
				return false
			},
			Validator: func(key string, c echo.Context) (bool, error) {
				token, err := getToken()
				if len(err) != 0 {
					println(err)
					return false, c.String(500, err)
				}
				if len(token) == 0 {
					return true, nil
				}
				return key == token, nil
			},
		}))
	e.Logger.Fatal(e.Start(":2399"))
}
