package handlers

import (
	"strconv"
	"weather-api/internal/clients"

	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	client *clients.WeatherClient
}

func NewWeatherHandler(client *clients.WeatherClient) *WeatherHandler {
	return &WeatherHandler{client: client}
}

func (h *WeatherHandler) ShowCurrent(c *gin.Context) {
	city := c.DefaultQuery("q", "London")
	weather, err := h.client.GetCurrent(city)
	if err != nil {
		c.JSON(500, gin.H{"error": "weather request failed"})
		return
	}
	c.JSON(200, weather)
}

func (h *WeatherHandler) ShowForecast(c *gin.Context) {
	city := c.DefaultQuery("q", "London")
	daysStr := c.DefaultQuery("days", "3")
	days, _ := strconv.Atoi(daysStr)
	forecast, err := h.client.GetForecast(city, days)
	if err != nil {
		c.JSON(500, gin.H{"error": "weather request failed"})
		return
	}
	c.JSON(200, forecast)
}
