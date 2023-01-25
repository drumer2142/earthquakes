package types

type Filters struct {
	Parameters []Parameters `json:"parameters"`
}

type Parameters struct {
	DistanceInKm float64 `json:"DistanceInKm"`
	Timestamp    string  `json:"Timestamp"`
	Depth        float64 `json:"Depth"`
	Magnitude    float64 `json:"Magnitude"`
}
