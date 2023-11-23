package space

type Planet string

const earthSeconds = 31557600.0

func Age(seconds float64, planet Planet) float64 {
	earthYears := seconds / earthSeconds

	switch planet {
	case "Mercury":
		return earthYears / 0.2408467
	case "Venus":
		return earthYears / 0.61519726
	case "Earth":
		return earthYears / 1.0
	case "Mars":
		return earthYears / 1.8808158
	case "Jupiter":
		return earthYears / 11.862615
	case "Saturn":
		return earthYears / 29.447498
	case "Uranus":
		return earthYears / 84.016846
	case "Neptune":
		return earthYears / 164.79132
	default:
		return -1.0
	}
}
