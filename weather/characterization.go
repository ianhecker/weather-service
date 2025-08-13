package weather

type Characterization string

const (
	Cold     Characterization = "cold"
	Moderate                  = "moderate"
	Hot                       = "hot"
)

func MakeCharacterization(temperature int) Characterization {
	return TemperatureToCharacterization(temperature)
}

func TemperatureToCharacterization(temperature int) Characterization {
	switch {
	case temperature > 75:
		return Hot
	case temperature < 30:
		return Cold
	default:
		return Moderate
	}
}

func (c Characterization) String() string {
	return string(c)
}
