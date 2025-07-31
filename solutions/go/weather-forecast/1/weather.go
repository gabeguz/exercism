// Package weather gets the forecast for a city based on the parameters you send it
package weather

// The current weather condition
var CurrentCondition string

// The current weather location
var CurrentLocation string

// Forecast returns a formatted string with the location's current weather condition
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
