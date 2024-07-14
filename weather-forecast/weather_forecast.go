// Package weather provides tools to forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition represents a current weather condition.
var CurrentCondition string

// CurrentLocation represents a current city.
var CurrentLocation string

// Forecast returns current weather condition for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
