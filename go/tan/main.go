// Program to test Tan
// Port of the elefunt tan.f and dtan.f test programs by W.J. Cody
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
	three := 3.0
	a := zero
	b := math.Pi / 4.0 // 0.785398163
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
			y := x / three
			y = (x + y) - x
			x = three * y

			var z, zz, w float64
			if j <= 2 {
				// Test TAN(X) vs TAN(X/3) identity
				z = math.Tan(x)
				zz = math.Tan(y)
				// TAN(3Y) = TAN(Y)*(3-TAN(Y)^2)/(1-3*TAN(Y)^2)
				w = one
				if z != zero {
					zz2 := zz * zz
					computed := zz * (three - zz2) / (one - three*zz2)
					w = (z - computed) / z
				}
			} else {
				// Test COT(X) = 1/TAN(X)
				z = math.Tan(x)
				if z != zero {
					zz = one / z
					cotx := one / math.Tan(x)
					w = (zz - cotx) / zz
				} else {
					w = one
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
			fmt.Println("\nTEST OF TAN(X) VS TAN(X/3) IDENTITY")
		} else {
			fmt.Println("\nTEST OF COT(X) = 1/TAN(X)")
		}
		fmt.Println()
		fmt.Printf("%7d RANDOM ARGUMENTS WERE TESTED FROM THE INTERVAL\n", n)
		fmt.Printf("      (%.4E, %.4E)\n\n", a, b)
		fmt.Printf(" TAN(X) WAS LARGER %6d TIMES,\n", k1)
		fmt.Printf("           AGREED %6d TIMES, AND\n", k2)
		fmt.Printf("       WAS SMALLER %6d TIMES.\n\n", k3)
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

		a = 6.0 * math.Pi
		b = a + math.Pi/4.0
	}

	// Special tests
	fmt.Println("\nSPECIAL TESTS")
	fmt.Println()
	fmt.Println(" THE IDENTITY   TAN(-X) = -TAN(X)   WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         F(X) + F(-X)")

	for i := 1; i <= 5; i++ {
		x := rng.Float64() * a
		z := math.Tan(x) + math.Tan(-x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" THE IDENTITY TAN(X) = X , X SMALL, WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         X - F(X)")

	betap := math.Pow(beta, float64(mp.IT))
	x := rng.Float64() / betap

	for i := 1; i <= 5; i++ {
		z := x - math.Tan(x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
		x = x / beta
	}

	// Test of error returns
	fmt.Println()
	fmt.Println("TEST OF ERROR RETURNS")
	fmt.Println()

	x = math.Pi / 2.0
	fmt.Printf(" TAN WILL BE CALLED WITH THE ARGUMENT %.16E\n", x)
	fmt.Println(" THIS SHOULD NOT CAUSE AN ERROR (TAN(PI/2) is large but finite in IEEE)")
	fmt.Println()
	y := math.Tan(x)
	fmt.Printf(" TAN RETURNED THE VALUE %.4E\n\n", y)

	fmt.Println(" THIS CONCLUDES THE TESTS")
}
