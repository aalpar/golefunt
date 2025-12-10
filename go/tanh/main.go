// Program to test Tanh
// Port of the elefunt tanh.f and dtanh.f test programs by W.J. Cody
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
	two := 2.0

	a := zero
	b := 0.5
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

			// Test TANH(X) using identity
			// TANH(2X) = 2*TANH(X)/(1+TANH(X)^2)
			y := x / two
			z := math.Tanh(x)
			zz := math.Tanh(y)
			var w float64
			if z != zero {
				computed := two * zz / (one + zz*zz)
				w = (z - computed) / z
			} else {
				w = one
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

		fmt.Println("\nTEST OF TANH(X) VS 2*TANH(X/2)/(1+TANH(X/2)**2)")
		fmt.Println()
		fmt.Printf("%7d RANDOM ARGUMENTS WERE TESTED FROM THE INTERVAL\n", n)
		fmt.Printf("      (%.4E, %.4E)\n\n", a, b)
		fmt.Printf(" TANH(X) WAS LARGER %6d TIMES,\n", k1)
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

		a = 0.5
		b = 5.0
		if j == 2 {
			a = 5.0
			b = 20.0
		}
	}

	// Special tests
	fmt.Println("\nSPECIAL TESTS")
	fmt.Println()
	fmt.Println(" THE IDENTITY   TANH(-X) = -TANH(X)   WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         F(X) + F(-X)")

	for i := 1; i <= 5; i++ {
		x := rng.Float64() * 5.0
		z := math.Tanh(x) + math.Tanh(-x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" THE IDENTITY TANH(X) = X , X SMALL, WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         X - F(X)")

	betap := math.Pow(beta, float64(mp.IT))
	x := rng.Float64() / betap

	for i := 1; i <= 5; i++ {
		z := x - math.Tanh(x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
		x = x / beta
	}

	fmt.Println()
	fmt.Println(" TEST OF SPECIAL ARGUMENTS")
	fmt.Println()

	x = zero
	y := math.Tanh(x)
	fmt.Printf(" TANH(0.0) = %.7E\n", y)

	// TANH should approach ±1 for large arguments
	x = 20.0
	y = math.Tanh(x)
	fmt.Printf(" TANH(20.0) = %.17E (should be very close to 1.0)\n", y)

	x = -20.0
	y = math.Tanh(x)
	fmt.Printf(" TANH(-20.0) = %.17E (should be very close to -1.0)\n", y)

	// Test of error returns
	fmt.Println()
	fmt.Println("TEST OF ERROR RETURNS")
	fmt.Println()

	x = mp.XMax
	fmt.Printf(" TANH WILL BE CALLED WITH THE ARGUMENT %.4E\n", x)
	fmt.Println(" THIS SHOULD RETURN 1.0 (NO OVERFLOW)")
	fmt.Println()
	y = math.Tanh(x)
	fmt.Printf(" TANH RETURNED THE VALUE %.17E\n\n", y)

	fmt.Println(" THIS CONCLUDES THE TESTS")
}
