// Package machar provides machine characteristic parameters for floating-point arithmetic.
// This is a Go port of the MACHAR subroutine by W.J. Cody, Argonne National Laboratory.
package machar

import "math"

// Params holds the machine characteristic parameters for floating-point arithmetic.
type Params struct {
	IBeta  int     // Radix of the floating-point representation
	IT     int     // Number of base IBeta digits in the significand
	IRnd   int     // Rounding mode (0=chop, 1=round, 2=IEEE round, 3-5=with partial underflow)
	NGrd   int     // Number of guard digits for multiplication
	MachEp int     // Largest negative int such that 1.0+IBeta^MachEp != 1.0
	NegEp  int     // Largest negative int such that 1.0-IBeta^NegEp != 1.0
	IExp   int     // Number of bits reserved for exponent
	MinExp int     // Largest magnitude negative int such that IBeta^MinExp is normalized
	MaxExp int     // Smallest positive power of IBeta that overflows
	Eps    float64 // Smallest positive number such that 1.0+Eps != 1.0
	EpsNeg float64 // Smallest positive number such that 1.0-EpsNeg != 1.0
	XMin   float64 // Smallest non-vanishing normalized floating-point power of the radix
	XMax   float64 // Largest finite floating-point number
}

// Float64 returns machine parameters for float64 (double precision).
// Uses Go's math package constants which reflect IEEE 754 double precision.
func Float64() Params {
	return Params{
		IBeta:  2,
		IT:     53,                    // 53 bits in mantissa (including implicit bit)
		IRnd:   5,                     // IEEE rounding with partial underflow (denormals)
		NGrd:   0,                     // No guard digits needed with rounding
		MachEp: -52,                   // 2^-52
		NegEp:  -53,                   // 2^-53
		IExp:   11,                    // 11 bits for exponent
		MinExp: -1021,                 // Minimum exponent
		MaxExp: 1024,                  // Maximum exponent
		Eps:    math.Pow(2, -52),      // ~2.22e-16
		EpsNeg: math.Pow(2, -53),      // ~1.11e-16
		XMin:   math.SmallestNonzeroFloat64 * math.Pow(2, 52), // ~2.23e-308
		XMax:   math.MaxFloat64,       // ~1.80e+308
	}
}

// Float32 returns machine parameters for float32 (single precision).
func Float32() Params {
	return Params{
		IBeta:  2,
		IT:     24,                         // 24 bits in mantissa
		IRnd:   5,                          // IEEE rounding with partial underflow
		NGrd:   0,                          // No guard digits needed
		MachEp: -23,                        // 2^-23
		NegEp:  -24,                        // 2^-24
		IExp:   8,                          // 8 bits for exponent
		MinExp: -125,                       // Minimum exponent
		MaxExp: 128,                        // Maximum exponent
		Eps:    float64(math.Pow(2, -23)),  // ~1.19e-7
		EpsNeg: float64(math.Pow(2, -24)),  // ~5.96e-8
		XMin:   float64(math.SmallestNonzeroFloat32) * math.Pow(2, 23),
		XMax:   float64(math.MaxFloat32),
	}
}

// Machar dynamically determines machine parameters using the original algorithm.
// This is a direct port of Cody's MACHAR routine.
func Machar() Params {
	var p Params

	one := 1.0
	two := 2.0
	zero := 0.0

	// Determine IBeta, Beta ala Malcolm
	a := one
	for {
		a = a + a
		temp := a + one
		temp1 := temp - a
		if temp1-one == zero {
			continue
		}
		break
	}

	b := one
	for {
		b = b + b
		temp := a + b
		itemp := int(temp - a)
		if itemp == 0 {
			continue
		}
		p.IBeta = itemp
		break
	}
	beta := float64(p.IBeta)

	// Determine IT, IRnd
	p.IT = 0
	b = one
	for {
		p.IT = p.IT + 1
		b = b * beta
		temp := b + one
		temp1 := temp - b
		if temp1-one == zero {
			continue
		}
		break
	}

	p.IRnd = 0
	betah := beta / two
	temp := a + betah
	if temp-a != zero {
		p.IRnd = 1
	}
	tempa := a + beta
	temp = tempa + betah
	if p.IRnd == 0 && temp-tempa != zero {
		p.IRnd = 2
	}

	// Determine NegEp, EpsNeg
	negep := p.IT + 3
	betain := one / beta
	a = one
	for i := 1; i <= negep; i++ {
		a = a * betain
	}
	b = a

	for {
		temp := one - a
		if temp-one != zero {
			break
		}
		a = a * beta
		negep = negep - 1
	}
	p.NegEp = -negep
	p.EpsNeg = a

	// Determine MachEp, Eps
	machep := -p.IT - 3
	a = b
	for {
		temp := one + a
		if temp-one != zero {
			break
		}
		a = a * beta
		machep = machep + 1
	}
	p.MachEp = machep
	p.Eps = a

	// Determine NGrd
	p.NGrd = 0
	temp = one + p.Eps
	if p.IRnd == 0 && temp*one-one != zero {
		p.NGrd = 1
	}

	// Determine IExp, MinExp, XMin
	i := 0
	k := 1
	z := betain
	t := one + p.Eps
	nxres := 0
	y := zero

	for {
		y = z
		z = y * y
		a = z * one
		temp = z * t
		if a+a == zero || math.Abs(z) >= y {
			break
		}
		temp1 := temp * betain
		if temp1*beta == z {
			break
		}
		i = i + 1
		k = k + k
	}

	if p.IBeta != 10 {
		p.IExp = i + 1
		mx := k + k
		p.MinExp, p.XMin, nxres = determineMinExp(y, betain, t, beta, k, mx, nxres)
		p.MaxExp = mx + p.MinExp
	} else {
		p.IExp = 2
		iz := p.IBeta
		for k >= iz {
			iz = iz * p.IBeta
			p.IExp = p.IExp + 1
		}
		mx := iz + iz - 1
		p.MinExp, p.XMin, nxres = determineMinExp(y, betain, t, beta, k, mx, nxres)
		p.MaxExp = mx + p.MinExp
	}

	// Adjust IRnd for partial underflow
	p.IRnd = p.IRnd + nxres

	// Adjust for IEEE-style machines
	if p.IRnd >= 2 {
		p.MaxExp = p.MaxExp - 2
	}

	// Adjust for implicit leading bit
	isum := p.MaxExp + p.MinExp
	if p.IBeta == 2 && isum == 0 {
		p.MaxExp = p.MaxExp - 1
	}
	if isum > 20 {
		p.MaxExp = p.MaxExp - 1
	}
	if a != y {
		p.MaxExp = p.MaxExp - 2
	}

	// Determine XMax
	p.XMax = one - p.EpsNeg
	if p.XMax*one != p.XMax {
		p.XMax = one - beta*p.EpsNeg
	}
	p.XMax = p.XMax / (beta * beta * beta * p.XMin)
	isum = p.MaxExp + p.MinExp + 3
	for j := 1; j <= isum; j++ {
		if p.IBeta == 2 {
			p.XMax = p.XMax + p.XMax
		} else {
			p.XMax = p.XMax * beta
		}
	}

	return p
}

func determineMinExp(y, betain, t, beta float64, k, mx, nxres int) (minexp int, xmin float64, nxresOut int) {
	xmin = y
	for {
		y = y * betain
		a := y * 1.0
		temp := y * t
		if (a+a) == 0.0 || math.Abs(y) >= xmin {
			break
		}
		k = k + 1
		temp1 := temp * betain
		if temp1*beta != y || temp == y {
			xmin = y
			continue
		} else {
			nxres = 3
			xmin = y
		}
	}
	return -k, xmin, nxres
}
