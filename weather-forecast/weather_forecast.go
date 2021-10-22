// Package weather is to help a weather station manage their weather forecasting program.
package weather

// CurrentCondition is holds current weather condition.
var CurrentCondition string
// CurrentLocation is golds current location information.
var CurrentLocation string

// Forecast is takings 2 integers that city and condition and returns weather condition on city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
