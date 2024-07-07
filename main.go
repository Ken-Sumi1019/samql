package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/Ken-Sumi1019/samql/routers"
	"github.com/Ken-Sumi1019/samql/templates"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = templates.Init()
	routers.SetRouting(e)

	e.Logger.Fatal(e.Start(":1323"))
}
