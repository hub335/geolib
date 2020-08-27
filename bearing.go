package geolib

import (
	"math"
)

// Bearing {from, to} a point
// Result:
//  - Type: float64 (initial bearing)
//  - Type: float64 (final bearing)
//  - Metric: Degress from North
func Bearing(φ1, λ1, φ2, λ2 float64) (float64, float64) {
	return initial(φ1, λ1, φ2, λ2), final(φ2, λ2, φ1, λ1)
}

func initial(φ1, λ1, φ2, λ2 float64) float64 {
	φ1 = Deg2Rad(φ1)
	φ2 = Deg2Rad(φ2)

	Δλ := Deg2Rad(λ2 - λ1)

	bearing := Rad2Deg(
		math.Atan2(
			math.Sin(
				Δλ)*math.Cos(φ2), math.Cos(φ1)*math.Sin(φ2)-math.Sin(φ1)*math.Cos(φ2)*math.Cos(Δλ)))

	return math.Mod(bearing+360, 360)
}

func final(φ1, λ1, φ2, λ2 float64) float64 {
	φ1 = Deg2Rad(φ1)
	φ2 = Deg2Rad(φ2)

	Δλ := Deg2Rad(λ2 - λ1)

	bearing := Rad2Deg(
		math.Atan2(
			math.Sin(
				Δλ)*math.Cos(φ2), math.Cos(φ1)*math.Sin(φ2)-math.Sin(φ1)*math.Cos(φ2)*math.Cos(Δλ)))

	return math.Mod(bearing+180, 360)
}
