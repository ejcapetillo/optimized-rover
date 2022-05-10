package model

type RoverName string

const (
	Curiosity   RoverName = "curiosity"
	Opportunity RoverName = "opportunity"
	Spirit      RoverName = "spirit"
)

func GetRovers() []RoverName {
	return []RoverName{Curiosity, Opportunity, Spirit}
}

type Rover struct {
	Id     int64     `json:"id"`
	Name   RoverName `json:"name"`
	Status string    `json:"status"`
}
