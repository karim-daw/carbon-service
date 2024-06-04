package model

// ProfileType enum for type safety
type ProfileType string

const (
	Embodied    ProfileType = "Embodied"
	Operational ProfileType = "Operational"
)

// CarbonProfile struct for carbon profiles (Single Responsibility)
type CarbonProfile struct {
	ProfileType  ProfileType
	EmissionRate float64
}
