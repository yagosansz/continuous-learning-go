// Package weather can forecast the weather conditions
// for various different cities in Goblinocus.
package weather

// CurrentCondition holds the current weather condition of a location.
var CurrentCondition string

// CurrentLocation holds the current location that we want to forecast
// a weather condition for.
var CurrentLocation string

// Forecast takes in the city and condition as arguments and returns
// the current weather condition for the specified location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
