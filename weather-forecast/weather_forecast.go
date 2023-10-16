// Package weather provides current weather forecast.
package weather

// CurrentCondition is the current weather condition.
var CurrentCondition string

// CurrentLocation is the current weather location.
var CurrentLocation string

// Forecast returns today's forecast based on the current weather condition and location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
