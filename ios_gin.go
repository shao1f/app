package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shao1f/app/handler"
)

func main() {
	route := gin.Default()
	route.Any("/testing", handler.StartPage)
	route.Any("/weather", handler.GetWeather)
	route.GET("/search", handler.Search)
	route.Run(":8080")
}
