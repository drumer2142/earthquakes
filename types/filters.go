package types

type Filters struct {
	Parameters []Parameters `json:"parameters"`
}

type Parameters struct {
	MaxDistanseInKM float64 `json:"MaxDistanseInKM"`
	Timestamp       string  `json:"Timestamp"`
	MinDepth        float64 `json:"MinDepth"`
	MinMagnitude    float64 `json:"MinMagnitude"`
}
