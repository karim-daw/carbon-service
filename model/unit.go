package model

import "gorm.io/gorm"

type UnitConverter interface {
	Convert(b bool) Unit
}

type Unit struct {
	gorm.Model
	Metric  bool   `gorm:"default:true;"`
	Area    string `gorm:"default:'m2';"`
	Volume  string `gorm:"default:'m3';"`
	Energy  string `gorm:"default:'kWh';"`
	Mass    string `gorm:"default:'kg';"`
	Density string `gorm:"default:'kg/m3';"`
	Carbon  string `gorm:"default:'kgco2';"` // tco2, kgco2/m2, kgco2/m2/year
}

func (u *Unit) Convert(isMetric bool) {
	if isMetric {
		u.Metric = false
		u.Area = "ft2"
		u.Volume = "ft3"
		u.Energy = "kBtu"
		u.Mass = "lb"
		u.Density = "lb/ft3"
		u.Carbon = "lbco2"

	} else {
		u.Metric = true
		u.Area = "m2"
		u.Volume = "m3"
		u.Energy = "kWh"
		u.Mass = "kg"
		u.Density = "kg/m3"
		u.Carbon = "kgco2"
	}
}
