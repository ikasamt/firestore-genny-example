package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//go:generate genny -in=ft/something_crud.go -out=ft/z-something_crud.go gen "Something=User,Article"
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/user/:userID", UserShowHandler())
	e.Start(":1981")
}
