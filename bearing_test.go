package geolib

import (
	"fmt"
	"testing"
)

func bearing(t *testing.T, expectInitial string, expectFinal string, φ1, λ1, φ2, λ2 float64) {
	initialBearing, finalBearing := Bearing(φ1, λ1, φ2, λ2)
	if result := fmt.Sprintf("%.2f", initialBearing); expectInitial != result {
		t.Errorf("Initial Bearing = expect %s; result %s", expectInitial, result)
	}
	if result := fmt.Sprintf("%.2f", finalBearing); expectFinal != result {
		t.Errorf("Final Bearing = expect %s; result %s", expectFinal, result)
	}
}

func TestBearing1(t *testing.T) {
	bearing(t, "48.93", "52.60", 50.116667, 8.683333, 52.516667, 13.383333)
}

func TestBearing2(t *testing.T) {
	bearing(t, "232.60", "228.93", 52.516667, 13.383333, 50.116667, 8.683333)
}

func TestBearing3(t *testing.T) {
	bearing(t, "146.93", "32.86", -41.32, 174.81, 40.96, -5.50)
}

func TestBearing4(t *testing.T) {
	bearing(t, "179.11", "179.23", 51.5104, 7.3256, 43.778889, 7.491)
}

func TestBearing5(t *testing.T) {
	bearing(t, "19.79", "21.12", 39.778889, -104.9825, 43.778889, -102.9825)
}
