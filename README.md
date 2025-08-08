# Weather API Server

This is a simple Go server that uses [WeatherAPI](https://www.weatherapi.com/) to provide weather data.

## Prerequisites

- [Go](https://golang.org/dl/) installed
- An account on [WeatherAPI](https://www.weatherapi.com/) with a valid API key

## Setup

1. **Clone or download** this repository to your local machine.

2. **Register** at [WeatherAPI](https://www.weatherapi.com/) and obtain your unique API key.

3. **Create `.env` file** in the root of the project: ```
env
API_KEY=your_api_key_here```
4. **Open the folder with the files:** ```
cd /path/to/weather-api
go mod init weather-api```
5. **Install dependencies:** `go get github.com/gin-gonic/gin`
6. **Run the server** `go run main.go`

- **Working on:** [localhost](http://localhost:8080)
- **Supportable requests:** `/current.json?q=YourCity` and `/forecast.json?q=YourCity&days=YourDays`
