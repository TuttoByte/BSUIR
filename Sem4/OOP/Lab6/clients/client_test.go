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

func TestNewOpenWeatherClient_Api(t *testing.T) {

}

func TestGoogleWetherClient(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t,
			"/currentConditions:lookup?key=testkey&location.latitude=55.7558&location.longitude=37.6173", r.URL.String())

		resp := googleWeatherResponse{
			Temperature: struct {
				Degrees decimal.Decimal `json:"degrees"`
			}{Degrees: decimal.NewFromFloat(20.5)}}

		json.NewEncoder(w).Encode(resp)
	}))

	cli := NewGoogleWetherClient("testkey", ts.URL+"/currentConditions")
	temp, err := cli.LocationCurrentTemperature(decimal.NewFromFloat(55.7558), decimal.NewFromFloat(37.6173))
	require.NoError(t, err)
	assert.Equal(t, decimal.NewFromFloat(20.5), temp)
}

func TestOpenWeatherClient_Coordinates(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/geo/1.0/direct?q=London,&limit=1&appid=testkey", r.URL.String())

		resp := newOpenWeatherCoordResponce([]openWeatherCityInfo{
			{
				Name:    "London",
				Lat:     decimal.NewFromFloat(51.5073219),
				Lon:     decimal.NewFromFloat(-0.1276474),
				Country: "GB",
			},
		})

		json.NewEncoder(w).Encode(resp)
	}))

	cli := NewOpenWeatherCoords("testkey", ts.URL+"/geo/1.0/direct")
	lat, lan, err := cli.GetCurrentLocation("London")
	assert.Nil(t, err)
	assert.Equal(t, decimal.NewFromFloat(51.5073219), lat)
	assert.Equal(t, decimal.NewFromFloat(-0.1276474), lan)

}
