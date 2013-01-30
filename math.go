package linAlg

import (
	"math"
)

// Returns the normal vector of the plane that intersects the three points p1,
// cp, and p2. Follows the 'right hand rule' of:
// vector(p1 - cp) cross vector(p2 - cp)
func FindTriangleNormal(p1, cp, p2 Point) Vector {
	v := p1.Sub(cp)
	w := p2.Sub(cp)
	n := v.Cross(w)
	return n.Unit()
}

// Returns the point on the surface of a sphere of size radius based on the
// vertical and horizontal angles. verAngle is the angle with respect to the
// XY plane, and horAngle is the angle on the XY plane with respect to the
// positive X axis. Both angles should be provided in radians.
func SpherePoint(radius, verAngle, horAngle float64) Point {
	temp := radius * math.Cos(verAngle)
	return Point{
		X: temp * math.Cos(horAngle),
		Y: temp * math.Sin(horAngle),
		Z: radius * math.Sin(verAngle),
	}
}

// Law Of Cosines. Returns the angle across from sideA in radians
func LOCangle(sideA, sideB, sideC float64) float64 {
	return math.Acos((sideB*sideB + sideC*sideC - sideA*sideA) / (2 * sideB * sideC))
}

// Law Of Cosines. Returns the length of the side across angleA, angleA must be in radians
func LOCside(angleA, sideB, sideC float64) float64 {
	return math.Sqrt(sideB*sideB + sideC*sideC - 2*sideB*sideC*math.Cos(angleA))
}

// Convert radians to degrees
func RtoD(radians float64) float64 {
	return radians * 180 / math.Pi
}

// Convert degrees to radians
func DtoR(degrees float64) float64 {
	return degrees * math.Pi / 180
}
