// Package weather this is package weather.
package weather

// CurrentCondition this is var CurrentCondition.
var CurrentCondition string
// CurrentLocation this is var CurrentLocation.
var CurrentLocation string

// Forecast will return string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
