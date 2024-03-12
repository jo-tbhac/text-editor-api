package main

import (
	"github.com/jo-tbhac/text-editor-api/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/users", handler.Create)
	e.Logger.Fatal(e.Start(":8080"))
}
