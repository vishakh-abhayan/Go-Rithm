// Package weather provides utilities for weather forecasting in cities.
package weather

// CurrentCondition stores the current weather state for a location.
var CurrentCondition string
// CurrentLocation represents the city for which we are reporting the weather.
var CurrentLocation string


// Forecast returns the weather forecast for the given city and weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
