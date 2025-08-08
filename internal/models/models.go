package models

// структура  запроса на данный момент
type WeatherResponse struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

// структура запроса прогноза
type ForecastResponse struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`

	Forecast struct {
		Forecastday []struct {
			Date string `json:"date"`
			Day  struct {
				AvgtempC  float64 `json:"avgtemp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"day"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

// для вывода
type DayForecast struct {
	Date      string  `json:"date"`
	TempC     float64 `json:"temp_c"`
	Condition string  `json:"condition"`
}
