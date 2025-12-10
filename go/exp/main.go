// Program to test Exp
// Port of the elefunt exp.f and dexp.f test programs by W.J. Cody
package main

import (
	"fmt"
	"math"

	"golefunt/machar"
	"golefunt/random"
)

var (
	Version = "dev"
	GitSHA  = "unknown"
)

func main() {
	// Get machine parameters
	mp := machar.Float64()
	rng := random.New()

	beta := float64(mp.IBeta)
	albeta := math.Log(beta)
	ait := float64(mp.IT)
	one := 1.0
	two := 2.0
	ten := 10.0
	zero := 0.0
	v := 0.0625
	a := two
	b := math.Log(a) * 0.5
	a = -b + v
	d := math.Log(0.9 * mp.XMax)
	n := 2000
	xn := float64(n)

	// Random argument accuracy tests
	for j := 1; j <= 3; j++ {
		k1 := 0
		k3 := 0
		x1 := zero
		r6 := zero
		r7 := zero
		del := (b - a) / xn
		xl := a

		for i := 1; i <= n; i++ {
			x := del*rng.Float64() + xl

			// Purify arguments
			y := x - v
			if y < zero {
				x = y + v
			}
			z := math.Exp(x)
			zz := math.Exp(y)

			if j == 1 {
				z = z - z*6.058693718652421388e-2
			} else {
				if mp.IBeta != 10 {
					z = z*0.0625 - z*2.4453321046920570389e-3
				} else {
					z = z*6.0e-2 + z*5.466789530794296106e-5
				}
			}

			w := one
			if zz != zero {
				w = (z - zz) / zz
			}
			if w < zero {
				k1++
			}
			if w > zero {
				k3++
			}
			w = math.Abs(w)
			if w > r6 {
				r6 = w
				x1 = x
			}
			r7 = r7 + w*w
			xl = xl + del
		}

		k2 := n - k3 - k1
		r7 = math.Sqrt(r7 / xn)

		fmt.Printf("\nTEST OF EXP(X-%.4f) VS EXP(X)/EXP(%.4f)\n\n", v, v)
		fmt.Printf("%7d RANDOM ARGUMENTS WERE TESTED FROM THE INTERVAL\n", n)
		fmt.Printf("      (%.4E, %.4E)\n\n", a, b)
		fmt.Printf(" EXP(X-V) WAS LARGER %6d TIMES,\n", k1)
		fmt.Printf("             AGREED %6d TIMES, AND\n", k2)
		fmt.Printf("         WAS SMALLER %6d TIMES.\n\n", k3)
		fmt.Printf(" THERE ARE %4d BASE %4d SIGNIFICANT DIGITS IN A FLOATING-POINT NUMBER\n\n", mp.IT, mp.IBeta)

		w := -999.0
		if r6 != zero {
			w = math.Log(math.Abs(r6)) / albeta
		}
		fmt.Printf(" THE MAXIMUM RELATIVE ERROR OF %.4E = %4d ** %7.2f\n", r6, mp.IBeta, w)
		fmt.Printf("    OCCURRED FOR X = %.6E\n", x1)

		wmax := math.Max(ait+w, zero)
		fmt.Printf(" THE ESTIMATED LOSS OF BASE %4d SIGNIFICANT DIGITS IS %7.2f\n\n", mp.IBeta, wmax)

		w = -999.0
		if r7 != zero {
			w = math.Log(math.Abs(r7)) / albeta
		}
		fmt.Printf(" THE ROOT MEAN SQUARE RELATIVE ERROR WAS %.4E = %4d ** %7.2f\n", r7, mp.IBeta, w)
		wmax = math.Max(ait+w, zero)
		fmt.Printf(" THE ESTIMATED LOSS OF BASE %4d SIGNIFICANT DIGITS IS %7.2f\n\n", mp.IBeta, wmax)

		if j != 2 {
			v = 45.0 / 16.0
			a = -ten * b
			b = 4.0 * mp.XMin * math.Pow(beta, float64(mp.IT))
			b = math.Log(b)
		} else {
			a = -two * a
			b = ten * a
			if b < d {
				b = d
			}
		}
	}

	// Special tests
	fmt.Println("\nSPECIAL TESTS")
	fmt.Println()
	fmt.Println(" THE IDENTITY  EXP(X)*EXP(-X) = 1.0  WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         F(X)*F(-X) - 1")

	for i := 1; i <= 5; i++ {
		x := rng.Float64() * beta
		y := -x
		z := math.Exp(x)*math.Exp(y) - one
		fmt.Printf("  %.7E  %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" TEST OF SPECIAL ARGUMENTS")
	fmt.Println()

	x := zero
	y := math.Exp(x) - one
	fmt.Printf(" EXP(0.0) - 1.0 = %.7E\n", y)

	x = math.Floor(math.Log(mp.XMin))
	y = math.Exp(x)
	fmt.Printf(" EXP(%.6E) = %.6E\n", x, y)

	x = math.Floor(math.Log(mp.XMax))
	y = math.Exp(x)
	fmt.Printf(" EXP(%.6E) = %.6E\n", x, y)

	x = x / two
	v = x / two
	y = math.Exp(x)
	z := math.Exp(v)
	z = z * z
	fmt.Printf("\n IF EXP(%.6E) = %.6E IS NOT ABOUT\n", x, y)
	fmt.Printf(" EXP(%.6E)**2 = %.6E THERE IS AN ARG RED ERROR\n", v, z)

	// Test of error returns
	fmt.Println()
	fmt.Println("TEST OF ERROR RETURNS")
	fmt.Println()

	x = -one / math.Sqrt(mp.XMin)
	fmt.Printf(" EXP WILL BE CALLED WITH THE ARGUMENT %.4E\n", x)
	fmt.Println(" THIS SHOULD UNDERFLOW")
	fmt.Println()
	y = math.Exp(x)
	fmt.Printf(" EXP RETURNED THE VALUE %.4E\n\n", y)

	x = -x
	fmt.Printf(" EXP WILL BE CALLED WITH THE ARGUMENT %.4E\n", x)
	fmt.Println(" THIS SHOULD OVERFLOW")
	fmt.Println()
	y = math.Exp(x)
	fmt.Printf(" EXP RETURNED THE VALUE %.4E\n\n", y)

	fmt.Println(" THIS CONCLUDES THE TESTS")
}
