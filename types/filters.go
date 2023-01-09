package types

type Filters struct {
	Parameters []Parameters `json:"parameters"`
}

type Parameters struct {
	DistanceInKm int     `json:"DistanceInKm"`
	Timestamp    string  `json:"Timestamp"`
	Depth        int     `json:"Depth"`
	Magnitude    float64 `json:"Magnitude"`
}
