// Package weather provides tools and functions to forecast the weather.
package weather


var (
    // CurrentCondition is the string which refers to the current weather condition.
	CurrentCondition string
   // CurrentLocation refers to current location the forecast refers to.
	CurrentLocation  string
)

// Forecast returns a string that merge currentCondition, another string and CurrentLocation.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
