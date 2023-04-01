package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/admin/:name", func(c echo.Context) error {
		resp := add(c.Param("name"), c.FormValue("pwd"), c.FormValue("shell"))
		println(resp)
		return c.String(200, resp+"\n")
	})
	e.DELETE("/admin/:name", func(c echo.Context) error {
		resp := del(c.Param("name"))
		println(resp)
		return c.String(200, resp+"\n")
	})
	e.GET("/admin", func(c echo.Context) error {
		resp := ls()
		println(resp)
		return c.String(200, resp+"\n")
	})
	e.GET("/:name", func(c echo.Context) error {
		resp := run(c.Param("name"))
		println(resp)
		return c.String(200, resp+"\n")
	})
	e.Logger.Fatal(e.Start(":2399"))
}
