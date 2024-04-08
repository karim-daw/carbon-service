package model

// Epd represents an environmental product declaration.
type Epd struct {
	// The unique identifier of the EPD.
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Type    string  `json:"type"`
	Unit    string  `json:"unit"`
	Value   float64 `json:"value"`
	Carbon  float64 `json:"carbon"`
	Area    float64 `json:"area"`
	Volume  float64 `json:"volume"`
	Energy  float64 `json:"energy"`
	Mass    float64 `json:"mass"`
	Density float64 `json:"density"`
}
