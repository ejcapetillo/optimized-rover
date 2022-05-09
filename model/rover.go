package model

type Rover string

const (
	Curiosity   Rover = "curiosity"
	Opportunity Rover = "opportunity"
	Spirit      Rover = "spirit"
)

func GetRovers() []Rover {
	return []Rover{Curiosity, Opportunity, Spirit}
}
