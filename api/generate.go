package api

import (
	"autobullshiter/service"

	"github.com/labstack/echo"
)

// Generate 生成字符串
func Generate(c echo.Context) error {
	service := service.GenerateService{}
	err := c.Bind(&service)
	if err == nil {
		res := service.Generate()
		return c.JSON(200, res)
	}
	return c.JSON(200, err)
}
