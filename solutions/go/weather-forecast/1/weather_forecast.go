// Package weather imports the package responsible of getting the weather conditions.
package weather

// CurrentCondition storage the weather condition of a location.
var CurrentCondition string
// CurrentLocation storage the location.
var CurrentLocation string

// Forecast return the current location with it specific current condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
