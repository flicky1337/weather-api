package clients

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"weather-api/internal/models"

	"github.com/joho/godotenv"
)

type WeatherClient struct {
	apiKey string
}

// load api key
func NewWeatherClient() *WeatherClient {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	key := os.Getenv("API_KEY")
	return &WeatherClient{apiKey: key}
}

func (wc *WeatherClient) GetCurrent(city string) (*models.WeatherResponse, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", wc.apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weather models.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

func (wc *WeatherClient) GetForecast(city string, days int) ([]models.DayForecast, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=%d", wc.apiKey, city, days)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var forecast models.ForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&forecast); err != nil {
		return nil, err
	}

	var result []models.DayForecast
	for _, day := range forecast.Forecast.Forecastday {
		result = append(result, models.DayForecast{
			Date:      day.Date,
			TempC:     day.Day.AvgtempC,
			Condition: day.Day.Condition.Text,
		})
	}

	return result, nil
}
