package clients

import (
	"encoding/json"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOpenWetherClient_Standart(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		assert.Equal(t, "/weather?lat=55.7558&lon=37.6173&appid=testkey&units=metric", r.URL.String())

		resp := openWeatherResponse{
			Main: struct {
				Temp decimal.Decimal `json:"temp"`
			}{Temp: decimal.NewFromFloat(20.5)},
		}

		json.NewEncoder(w).Encode(resp)
	}))

	cli := NewOpenWeatherClient("testkey", ts.URL+"/weather")
	temp, err := cli.LocationCurrentTemperature(decimal.NewFromFloat(55.7558), decimal.NewFromFloat(37.6173))
	require.NoError(t, err)
	assert.Equal(t, temp, decimal.NewFromFloat(20.5))
}

func TestGoogleWetherClient(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t,
			"/v1/currentConditions:lookup?key=testkey&location.latitude=55.7558&location.longitude=37.6173", r.URL.String())

		resp := googleWeatherResponse{
			Temperature: struct {
				Degrees decimal.Decimal `json:"degrees"`
			}{Degrees: decimal.NewFromFloat(20.5)}}

		json.NewEncoder(w).Encode(resp)
	}))

	cli := NewGoogleWetherClient("testkey", ts.URL)
	temp, err := cli.LocationCurrentTemperature(decimal.NewFromFloat(55.7558), decimal.NewFromFloat(37.6173))
	require.NoError(t, err)
	assert.Equal(t, decimal.NewFromFloat(20.5), temp)
}

func TestOpenWeatherClient_Coordinates(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/geo/1.0/direct?q=London&limit=5&appid=testkey", r.URL.String())

		resp := []OpenWeatherCityInfo{
			{
				Name:    "London",
				Lat:     decimal.NewFromFloat(51.5073219),
				Lon:     decimal.NewFromFloat(-0.1276474),
				Country: "GB",
			},
		}

		json.NewEncoder(w).Encode(resp)
	}))

	cli := NewOpenWeatherCoords("testkey", ts.URL)
	location, err := cli.GetLocation("London")
	assert.Nil(t, err)
	assert.Equal(t, decimal.NewFromFloat(51.5073219), location[0].Lat)
	assert.Equal(t, decimal.NewFromFloat(-0.1276474), location[0].Lon)
}

func TestGoogleWeatherClient_Forecast(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t,
			"/v1/currentConditions:lookup?key=testkey&location.latitude=55.7558&location.longitude=37.6173", r.URL.String())

		resp := googleForecastResponse{
			Temperature: struct {
				Degrees decimal.Decimal `json:"degrees"`
				Unit    string          `json:"unit"`
			}{Degrees: decimal.NewFromFloat(12.2), Unit: "CELSIUS"},
		}

		json.NewEncoder(w).Encode(resp)
	}))

	cli := NewGoogleWetherClient("testkey", ts.URL)
	temp, err := cli.LocationCurrentForcast(decimal.NewFromFloat(55.7558), decimal.NewFromFloat(37.6173))
	require.NoError(t, err)
	assert.Equal(t, decimal.NewFromFloat(12.2), temp.Temperature)
}

func TestOpenWeatherClient_Forecast(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t,
			"/v1/currentConditions:lookup?key=testkey&location.latitude=55.7558&location.longitude=37.6173", r.URL.String())

		resp := openForecastResponce{Main: struct {
			Temp      decimal.Decimal `json:"temp"`
			FeelsLike decimal.Decimal `json:"feels_like"`
			TempMin   decimal.Decimal `json:"temp_min"`
			TempMax   decimal.Decimal `json:"temp_max"`
			Pressure  decimal.Decimal `json:"pressure"`
			Humidity  int             `json:"humidity"`
			SeaLevel  int             `json:"sea_level"`
			GrndLevel int             `json:"grnd_level"`
		}{Temp: decimal.NewFromFloat(12.12), FeelsLike: decimal.NewFromFloat(12.12), TempMin: decimal.NewFromFloat(9.12), TempMax: decimal.NewFromFloat(20.12), Pressure: decimal.NewFromFloat(12.12), Humidity: 13, SeaLevel: 123, GrndLevel: 13}}

		json.NewEncoder(w).Encode(resp)
	}))

	cli := NewOpenWeatherClient("testkey", ts.URL)
	temp, err := cli.LocationCurrentForcast(decimal.NewFromFloat(55.7558), decimal.NewFromFloat(37.6173))
	require.NoError(t, err)
	assert.Equal(t, decimal.NewFromFloat(12.2), temp.Temperature)
}
