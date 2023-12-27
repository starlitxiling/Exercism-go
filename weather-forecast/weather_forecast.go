// Package weather provides weather forecast information.
package weather

// CurrentCondition represents weather condition.
var CurrentCondition string

// CurrentLocation represents weather address.
var CurrentLocation string

// Forecast returns string value representing weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
