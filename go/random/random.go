// Package random provides a simple random number generator compatible with the
// original elefunt REN function (Algorithm 266 by Pike and Hill).
package random

// Generator is a simple linear congruential random number generator.
type Generator struct {
	iy int
}

// New creates a new random number generator with the default seed.
func New() *Generator {
	return &Generator{iy: 100001}
}

// NewWithSeed creates a new random number generator with a custom seed.
func NewWithSeed(seed int) *Generator {
	return &Generator{iy: seed}
}

// Float64 returns a random float64 uniformly distributed over (0, 1).
// This is the double precision version (equivalent to Fortran REN).
func (g *Generator) Float64() float64 {
	g.iy = g.iy * 125
	g.iy = g.iy - (g.iy/2796203)*2796203
	// Double precision version includes additional factor for better distribution
	return float64(g.iy) / 2796203.0 * (1.0 + 1.0e-6 + 1.0e-12)
}

// Float32 returns a random float32 uniformly distributed over (0, 1).
// This is the single precision version (equivalent to Fortran REN).
func (g *Generator) Float32() float32 {
	g.iy = g.iy * 125
	g.iy = g.iy - (g.iy/2796203)*2796203
	return float32(g.iy) / 2796203.0
}

// Reset resets the generator to its initial state.
func (g *Generator) Reset() {
	g.iy = 100001
}

// Seed sets a new seed value.
func (g *Generator) Seed(seed int) {
	g.iy = seed
}
