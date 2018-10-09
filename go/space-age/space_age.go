package space

type Planet string

// Age() returns the age on a given planet based on seconds provided
func Age(seconds float64, planet Planet) float64 {
	v := seconds / 31557600

	switch planet {
	case "Earth":
		return v
	case "Mercury":
		v /= 0.2408467
	case "Venus":
		v /= 0.61519726
	case "Mars":
		v /= 1.8808158
	case "Jupiter":
		v /= 11.862615
	case "Saturn":
		v /= 29.447498
	case "Uranus":
		v /= 84.016846
	case "Neptune":
		v /= 164.79132
	default:
		return 0
	}

	return v
}
