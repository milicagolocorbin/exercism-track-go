// Package weather provides tools to forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents a current weather condition.
var CurrentCondition string

// CurrentLocation represents a current city location.
var CurrentLocation string

// Forecast returns weather conditions for a current city location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
