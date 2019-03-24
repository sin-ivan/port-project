package port

// Port is used for port info storage
type Port struct {
	ID          string    `json:"identifier"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Coordinates []float32 `json:"coordinates"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}
