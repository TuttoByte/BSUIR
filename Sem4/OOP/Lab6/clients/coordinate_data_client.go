package clients

type CoordinateDataCleint interface {
	GetLocation(cityName string) ([]OpenWeatherCityInfo, error)
}
