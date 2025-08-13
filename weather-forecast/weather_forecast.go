// Package weather:
// This package provides tools for meteorological forecasting.
package weather

// CurrentCondition: indicates current condition.
var CurrentCondition string

// CurrentLocation: indicates current location.
var CurrentLocation string

// Forecast: Returns the current condition of the city and the current weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
