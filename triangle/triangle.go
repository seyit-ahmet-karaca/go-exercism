// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	max := math.Max(a, math.Max(b, c))
	if a <= 0 || b <= 0 || c <= 0 || max >= a+b+c-max {
		return NaT
	}
	min := math.Min(a, math.Min(b, c))
	if max == min {
		return Equ
	}

	middle := (a + b + c) - (max + min)
	if max != middle && middle != min {
		return Sca
	}
	return Iso
}
