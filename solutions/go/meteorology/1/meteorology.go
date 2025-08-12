package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (t TemperatureUnit) String() string {
	if t == Celsius {
		return "°C"
	}
	return "°F"
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	if t.unit == Celsius {
		return fmt.Sprintf("%d °C", t.degree)
	}
	return fmt.Sprintf("%d °F", t.degree)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (s SpeedUnit) String() string {
	if s == KmPerHour {
		return "km/h"
	}
	return "mph"
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	if s.unit == KmPerHour {
		return fmt.Sprintf("%d km/h", s.magnitude)
	}
	return fmt.Sprintf("%d mph", s.magnitude)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%s Humidity", m.location, m.temperature, m.windDirection, m.windSpeed, m.humidity, "%")
}
