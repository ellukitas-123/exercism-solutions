// Package weather uses current conditions to forecast weather.
package weather

// CurrentCondition stores the actual weather condition.
var CurrentCondition string

// CurrentLocation takes the actual location.
var CurrentLocation string

// Forecast takes the location and the condition and returns a string describing.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
