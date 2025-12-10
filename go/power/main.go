// Program to test Power (X**Y)
// Port of the elefunt power.f and dpower.f test programs by W.J. Cody
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

	// Test X**Y using identity: X**(2Y) = (X**Y)**2
	a := one / beta
	b := one
	n := 2000
	xn := float64(n)

	// Random argument accuracy tests
	for j := 1; j <= 4; j++ {
		k1 := 0
		k3 := 0
		x1 := zero
		y1 := zero
		r6 := zero
		r7 := zero
		del := (b - a) / xn
		xl := a

		for i := 1; i <= n; i++ {
			x := del*rng.Float64() + xl

			var y, z, zz, w float64
			if j <= 2 {
				// Test X**(2Y) vs (X**Y)**2
				y = rng.Float64() * 2.0
				z = math.Pow(x, two*y)
				zz = math.Pow(x, y)
				zz = zz * zz
			} else {
				// Test X**Y vs EXP(Y*LOG(X))
				y = rng.Float64() * 2.0
				z = math.Pow(x, y)
				zz = math.Exp(y * math.Log(x))
			}

			w = one
			if z != zero {
				w = (z - zz) / z
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
				y1 = y
			}
			r7 = r7 + w*w
			xl = xl + del
		}

		k2 := n - k3 - k1
		r7 = math.Sqrt(r7 / xn)

		if j <= 2 {
			fmt.Println("\nTEST OF X**(2Y) VS (X**Y)**2")
		} else {
			fmt.Println("\nTEST OF X**Y VS EXP(Y*LOG(X))")
		}
		fmt.Println()
		fmt.Printf("%7d RANDOM ARGUMENTS WERE TESTED FROM THE INTERVAL\n", n)
		fmt.Printf("      X IN (%.4E, %.4E), Y IN (0, 2)\n\n", a, b)
		fmt.Printf(" X**Y WAS LARGER %6d TIMES,\n", k1)
		fmt.Printf("          AGREED %6d TIMES, AND\n", k2)
		fmt.Printf("      WAS SMALLER %6d TIMES.\n\n", k3)
		fmt.Printf(" THERE ARE %4d BASE %4d SIGNIFICANT DIGITS IN A FLOATING-POINT NUMBER\n\n", mp.IT, mp.IBeta)

		w := -999.0
		if r6 != zero {
			w = math.Log(math.Abs(r6)) / albeta
		}
		fmt.Printf(" THE MAXIMUM RELATIVE ERROR OF %.4E = %4d ** %7.2f\n", r6, mp.IBeta, w)
		fmt.Printf("    OCCURRED FOR X = %.6E, Y = %.6E\n", x1, y1)

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
	fmt.Println(" THE IDENTITY  X**1 = X  WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         X**1 - X")

	for i := 1; i <= 5; i++ {
		x := rng.Float64() * 10.0
		z := math.Pow(x, one) - x
		fmt.Printf("  %.7E  %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" THE IDENTITY  X**0 = 1  WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         X**0 - 1")

	for i := 1; i <= 5; i++ {
		x := rng.Float64() * 10.0
		z := math.Pow(x, zero) - one
		fmt.Printf("  %.7E  %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" TEST OF SPECIAL ARGUMENTS")
	fmt.Println()

	x := one
	y := zero
	z := math.Pow(x, y)
	fmt.Printf(" 1**0 = %.7E\n", z)

	x = zero
	y = one
	z = math.Pow(x, y)
	fmt.Printf(" 0**1 = %.7E\n", z)

	x = two
	y = two
	z = math.Pow(x, y)
	fmt.Printf(" 2**2 = %.7E (should be 4.0)\n", z)

	x = two
	y = 10.0
	z = math.Pow(x, y)
	fmt.Printf(" 2**10 = %.7E (should be 1024.0)\n", z)

	x = 10.0
	y = two
	z = math.Pow(x, y)
	fmt.Printf(" 10**2 = %.7E (should be 100.0)\n", z)

	// Test of error returns
	fmt.Println()
	fmt.Println("TEST OF ERROR RETURNS")
	fmt.Println()

	x = zero
	y = zero
	fmt.Printf(" 0**0 WILL BE COMPUTED\n")
	z = math.Pow(x, y)
	fmt.Printf(" 0**0 = %v\n\n", z)

	x = -two
	y = 3.5
	fmt.Printf(" (-2)**3.5 WILL BE COMPUTED\n")
	fmt.Println(" THIS SHOULD RETURN NaN")
	z = math.Pow(x, y)
	fmt.Printf(" (-2)**3.5 = %v\n\n", z)

	x = mp.XMax
	y = two
	fmt.Printf(" XMAX**2 WILL BE COMPUTED\n")
	fmt.Println(" THIS SHOULD OVERFLOW")
	z = math.Pow(x, y)
	fmt.Printf(" XMAX**2 = %v\n\n", z)

	fmt.Println(" THIS CONCLUDES THE TESTS")
}
