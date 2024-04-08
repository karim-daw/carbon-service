package converter

// UnitConverter defines the contract for objects that can perform unit conversion.
type UnitConversionService interface {
	ToMetric(value float64) float64
	ToImperial(value float64) float64
}

// UnitConversionFactory provides factory methods to create unit converters.
type unitConversionService struct{}

// NewUnitConversionFactory creates a new instance of UnitConversionFactory.
func NewUnitConversionService() *unitConversionService {
	return &unitConversionService{}
}

// GetConverter returns the appropriate converter based on the unit type.
func (f *unitConversionService) GetConverter(unitType string) UnitConversionService {
	switch unitType {
	case "carbon":
		return &CarbonConverter{}
	case "area":
		return &AreaConverter{}
	case "volume":
		return &VolumeConverter{}
	case "energy":
		return &EnergyConverter{}
	case "mass":
		return &MassConverter{}
	case "density":
		return &DensityConverter{}
	default:
		return nil
	}
}

// CarbonConverter converts carbon values.
type CarbonConverter struct{}

// converts the lbco2 to kgco2
func (c *CarbonConverter) ToMetric(value float64) float64 {
	// Implement conversion from imperial to metric for carbon values
	return value * LBCO2_TO_KGC02
}

// converts the kgco2 to lbco2
func (c *CarbonConverter) ToImperial(value float64) float64 {
	// Implement conversion from metric to imperial for carbon values
	return value * KGC02_TO_LBCO2
}

// AreaConverter converts area values.
type AreaConverter struct{}

// converts the ft2 to m2
func (c *AreaConverter) ToMetric(value float64) float64 {
	// Implement conversion from imperial to metric for area values
	return value * FT2_TO_M2
}

// converts the m2 to ft2
func (c *AreaConverter) ToImperial(value float64) float64 {
	// Implement conversion from metric to imperial for area values
	return value * M2_TO_FT2
}

// VolumeConverter converts volume values.
type VolumeConverter struct{}

// converts the ft3 to m3
func (c *VolumeConverter) ToMetric(value float64) float64 {
	// Implement conversion from imperial to metric for volume values
	return value * FT3_TO_M3
}

// converts the m3 to ft3
func (c *VolumeConverter) ToImperial(value float64) float64 {
	// Implement conversion from metric to imperial for volume values
	return value * M3_TO_FT3
}

// EnergyConverter converts energy values.
type EnergyConverter struct{}

// converts the kbtu to kwh
func (c *EnergyConverter) ToMetric(value float64) float64 {
	// Implement conversion from imperial to metric for energy values
	return value * KBTU_TO_KWH
}

// converts the kwh to kbtu
func (c *EnergyConverter) ToImperial(value float64) float64 {
	// Implement conversion from metric to imperial for energy values
	return value * KWH_TO_KBTU
}

// MassConverter converts mass values.
type MassConverter struct{}

// converts the lb to kg
func (c *MassConverter) ToMetric(value float64) float64 {
	// Implement conversion from imperial to metric for mass values
	return value * LB_TO_KG
}

// converts the kg to lb
func (c *MassConverter) ToImperial(value float64) float64 {
	// Implement conversion from metric to imperial for mass values
	return value * KG_TO_LB
}

// DensityConverter converts density values.
type DensityConverter struct{}

// converts the lb/ft3 to kg/m3
func (c *DensityConverter) ToMetric(value float64) float64 {
	// Implement conversion from imperial to metric for density values
	return value * LB_FT3_TO_KG_M3
}

// converts the kg/m3 to lb/ft3
func (c *DensityConverter) ToImperial(value float64) float64 {
	// Implement conversion from metric to imperial for density values
	return value * KG_M3_TO_LB_FT3
}
