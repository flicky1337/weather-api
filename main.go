package main

import (
	"weather-api/internal/clients"
	"weather-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	client := clients.NewWeatherClient()
	handler := handlers.NewWeatherHandler(client)

	r.GET("/current.json", handler.ShowCurrent)
	r.GET("/forecast.json", handler.ShowForecast)

	r.Run(":8080")
}
