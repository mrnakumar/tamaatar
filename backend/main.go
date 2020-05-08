package main

import (
	"github.com/labstack/echo/v4"
	"mrnakumar.com/tamaatar/handlers"
)

func main() {
	e := echo.New()
	e.Static("/","ui/dist/ui/index.html")
	e.POST("/sprint", handlers.SprintPutHandler).Name = "record a successful completion of a sprint"
	e.Logger.Fatal(e.Start(":8080"))
}
