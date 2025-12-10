// Program to test Sqrt
// Port of the elefunt sqrt.f and dsqrt.f test programs by W.J. Cody
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
	zero := 0.0

	a := one / beta
	b := one
	n := 2000
	xn := float64(n)

	// Random argument accuracy tests
	for j := 1; j <= 2; j++ {
		k1 := 0
		k3 := 0
		x1 := zero
		r6 := zero
		r7 := zero
		del := (b - a) / xn
		xl := a

		for i := 1; i <= n; i++ {
			x := del*rng.Float64() + xl

			// Test SQRT(X) vs X/SQRT(X)
			y := math.Sqrt(x)
			z := x / y
			w := one
			if y != zero {
				w = (y - z) / y
			}

			if w > zero {
				k1++
			}
			if w < zero {
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

		fmt.Println("\nTEST OF SQRT(X) VS X/SQRT(X)")
		fmt.Println()
		fmt.Printf("%7d RANDOM ARGUMENTS WERE TESTED FROM THE INTERVAL\n", n)
		fmt.Printf("      (%.4E, %.4E)\n\n", a, b)
		fmt.Printf(" SQRT(X) WAS LARGER %6d TIMES,\n", k1)
		fmt.Printf("            AGREED %6d TIMES, AND\n", k2)
		fmt.Printf("        WAS SMALLER %6d TIMES.\n\n", k3)
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

		a = one
		b = beta
	}

	// Special tests
	fmt.Println("\nSPECIAL TESTS")
	fmt.Println()
	fmt.Println(" THE IDENTITY  SQRT(X)*SQRT(X) = X  WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         F(X)*F(X) - X")

	for i := 1; i <= 5; i++ {
		x := rng.Float64()
		y := math.Sqrt(x)
		z := y*y - x
		fmt.Printf("  %.7E  %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" TEST OF SPECIAL ARGUMENTS")
	fmt.Println()

	x := mp.XMin
	y := math.Sqrt(x)
	fmt.Printf(" SQRT(XMIN) = SQRT(%.6E) = %.6E\n", x, y)

	x = one - mp.EpsNeg
	y = math.Sqrt(x)
	fmt.Printf(" SQRT(1-EPSNEG) = SQRT(%.17E) = %.17E\n", x, y)

	x = one
	y = math.Sqrt(x)
	fmt.Printf(" SQRT(1.0) = %.17E\n", y)

	x = one + mp.Eps
	y = math.Sqrt(x)
	fmt.Printf(" SQRT(1+EPS) = SQRT(%.17E) = %.17E\n", x, y)

	x = mp.XMax
	y = math.Sqrt(x)
	fmt.Printf(" SQRT(XMAX) = SQRT(%.6E) = %.6E\n", x, y)

	// Test of error returns
	fmt.Println()
	fmt.Println("TEST OF ERROR RETURNS")
	fmt.Println()

	x = zero
	fmt.Printf(" SQRT WILL BE CALLED WITH THE ARGUMENT %.4E\n", x)
	y = math.Sqrt(x)
	fmt.Printf(" SQRT RETURNED THE VALUE %.4E\n\n", y)

	x = -one
	fmt.Printf(" SQRT WILL BE CALLED WITH THE ARGUMENT %.4E\n", x)
	fmt.Println(" THIS SHOULD RETURN NaN")
	fmt.Println()
	y = math.Sqrt(x)
	fmt.Printf(" SQRT RETURNED THE VALUE %v\n\n", y)

	fmt.Println(" THIS CONCLUDES THE TESTS")
}
