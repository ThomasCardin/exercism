package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (temperatureUnit TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[temperatureUnit]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temperature Temperature) String() string {
	return fmt.Sprintf("%v %v", temperature.degree, temperature.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speedUnit SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[speedUnit]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (speed Speed) String() string {
	return fmt.Sprintf("%v %v", speed.magnitude, speed.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Output: San Francisco: 57 °F, Wind NW at 19 mph, 60% Humidity
// Add a String method to MeteorologyData
func (meteorologyData MeteorologyData) String() string {
	location := meteorologyData.location
	temperature := meteorologyData.temperature.String()
	windDirection := meteorologyData.windDirection
	speed := meteorologyData.windSpeed.String()
	humidity := meteorologyData.humidity
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", location, temperature, windDirection, speed, humidity)
}
