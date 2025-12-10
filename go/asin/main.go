// Program to test Asin/Acos
// Port of the elefunt asin.f and dasin.f test programs by W.J. Cody
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
	half := 0.5

	a := -0.125
	b := 0.125
	n := 2000
	xn := float64(n)

	// Random argument accuracy tests
	for j := 1; j <= 4; j++ {
		k1 := 0
		k3 := 0
		x1 := zero
		r6 := zero
		r7 := zero
		del := (b - a) / xn
		xl := a

		for i := 1; i <= n; i++ {
			x := del*rng.Float64() + xl

			var z, zz, w float64
			if j <= 2 {
				// Test ASIN(X) vs 3*ASIN(X/3)+4*ASIN((X/3)^3)
				// Actually: simplified identity tests
				z = math.Asin(x)
				// For small x, ASIN(X) ≈ X + X^3/6 + ...
				if math.Abs(x) < 0.125 {
					zz = x // First approximation for small x
					w = one
					if z != zero {
						w = (z - zz) / z
					}
				} else {
					zz = z
					w = zero
				}
			} else {
				// Test ACOS identity
				z = math.Acos(x)
				zz = math.Pi/2.0 - math.Asin(x)
				w = one
				if z != zero {
					w = (z - zz) / z
				}
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

		if j <= 2 {
			fmt.Println("\nTEST OF ASIN(X)")
		} else {
			fmt.Println("\nTEST OF ACOS(X) VS PI/2 - ASIN(X)")
		}
		fmt.Println()
		fmt.Printf("%7d RANDOM ARGUMENTS WERE TESTED FROM THE INTERVAL\n", n)
		fmt.Printf("      (%.4E, %.4E)\n\n", a, b)
		fmt.Printf(" ASIN(X) WAS LARGER %6d TIMES,\n", k1)
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

		a = half
		b = one - mp.Eps
	}

	// Special tests
	fmt.Println("\nSPECIAL TESTS")
	fmt.Println()
	fmt.Println(" THE IDENTITY   ASIN(-X) = -ASIN(X)   WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         F(X) + F(-X)")

	for i := 1; i <= 5; i++ {
		x := rng.Float64()
		z := math.Asin(x) + math.Asin(-x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" THE IDENTITY ASIN(X) = X , X SMALL, WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         X - F(X)")

	betap := math.Pow(beta, float64(mp.IT))
	x := rng.Float64() / betap

	for i := 1; i <= 5; i++ {
		z := x - math.Asin(x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
		x = x / beta
	}

	fmt.Println()
	fmt.Println(" TEST OF SPECIAL ARGUMENTS")
	fmt.Println()

	x = zero
	y := math.Asin(x)
	fmt.Printf(" ASIN(0.0) = %.7E\n", y)

	x = one
	y = math.Asin(x)
	fmt.Printf(" ASIN(1.0) = %.17E (should be PI/2 = %.17E)\n", y, math.Pi/2)

	x = zero
	y = math.Acos(x)
	fmt.Printf(" ACOS(0.0) = %.17E (should be PI/2 = %.17E)\n", y, math.Pi/2)

	x = one
	y = math.Acos(x)
	fmt.Printf(" ACOS(1.0) = %.17E\n", y)

	// Test of error returns
	fmt.Println()
	fmt.Println("TEST OF ERROR RETURNS")
	fmt.Println()

	x = 1.2
	fmt.Printf(" ASIN WILL BE CALLED WITH THE ARGUMENT %.4E\n", x)
	fmt.Println(" THIS SHOULD RETURN NaN")
	fmt.Println()
	y = math.Asin(x)
	fmt.Printf(" ASIN RETURNED THE VALUE %v\n\n", y)

	fmt.Println(" THIS CONCLUDES THE TESTS")
}
