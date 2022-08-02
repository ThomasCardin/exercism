// Package weather provides the forecast of Goblinocus cities.
package weather

// CurrentCondition of Goblinocus.
var CurrentCondition string

// CurrentLocation inside the city of Goblinocus.
var CurrentLocation string

// Forecast return the weather of Goblinocus cities.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
