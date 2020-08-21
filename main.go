package main

import (
	"autobullshiter/api"
	"autobullshiter/model"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	godotenv.Load()
	model.Init()
	e := echo.New()

	e.POST("/api/generate", api.Generate)

	e.Logger.Fatal(e.Start(":1323"))
}
