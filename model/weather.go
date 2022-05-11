package model

type SolWeather struct {
	Sol        int
	EarthDate  string
	MaxAirTemp int
	MinAirTemp int
	Pressure   int
}

func GetStaticCuriosityWeather() map[string]SolWeather {
	sol3463 := SolWeather{Sol: 3463, EarthDate: "2022-05-04", MaxAirTemp: 39, MinAirTemp: -90, Pressure: 825}
	sol3462 := SolWeather{Sol: 3462, EarthDate: "2022-05-03", MaxAirTemp: 43, MinAirTemp: -92, Pressure: 822}
	sol3461 := SolWeather{Sol: 3461, EarthDate: "2022-05-02", MaxAirTemp: 43, MinAirTemp: -96, Pressure: 822}
	sol3460 := SolWeather{Sol: 3460, EarthDate: "2022-05-01", MaxAirTemp: 34, MinAirTemp: -90, Pressure: 822}
	sol3459 := SolWeather{Sol: 3461, EarthDate: "2022-04-30", MaxAirTemp: 39, MinAirTemp: -92, Pressure: 820}
	sol3458 := SolWeather{Sol: 3461, EarthDate: "2022-04-29", MaxAirTemp: 41, MinAirTemp: -90, Pressure: 826}
	sol3457 := SolWeather{Sol: 3461, EarthDate: "2022-04-28", MaxAirTemp: 34, MinAirTemp: -90, Pressure: 816}
	return map[string]SolWeather{sol3463.EarthDate: sol3463, sol3462.EarthDate: sol3462, sol3461.EarthDate: sol3461,
		sol3460.EarthDate: sol3460, sol3459.EarthDate: sol3459, sol3458.EarthDate: sol3458, sol3457.EarthDate: sol3457}
}
