// Package weather this is weather package, initialize from the very first.
package weather

var (
	// CurrentCondition variable as a string.
	CurrentCondition string
	// CurrentLocation variable as a string.
	CurrentLocation string
)

// Forecast is a function to run the weather apps.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
