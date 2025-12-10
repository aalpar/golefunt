// Program to test Log
// Port of the elefunt alog.f and dlog.f test programs by W.J. Cody
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
	eight := 8.0

	// For log test: test interval is [1/sqrt(2), sqrt(2)]
	a := one / math.Sqrt(2.0)
	b := math.Sqrt(2.0)
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
			if j == 1 {
				// Test LOG(X) vs LOG(17X/16) - LOG(17/16)
				y := x - half
				y = (y + half) - half
				x = y + y/16.0
				z = math.Log(x)
				zz = math.Log(y) + math.Log(17.0/16.0)
			} else if j == 2 {
				// Test LOG(X) vs LOG(11X/10) - LOG(11/10)
				y := x - half
				y = (y + half) - half
				x = y + y/10.0
				z = math.Log(x)
				zz = math.Log(y) + math.Log(1.1)
			} else if j == 3 {
				// Test LOG(X*X) vs 2*LOG(X)
				z = math.Log(x * x)
				zz = 2.0 * math.Log(x)
			} else {
				// Test LOG10(X) vs LOG(X)/LOG(10)
				z = math.Log10(x)
				zz = math.Log(x) / math.Log(10.0)
			}

			w = one
			if z != zero {
				w = (z - zz) / z
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

		if j == 1 {
			fmt.Println("\nTEST OF LOG(X) VS LOG(17X/16) - LOG(17/16)")
		} else if j == 2 {
			fmt.Println("\nTEST OF LOG(X) VS LOG(11X/10) - LOG(11/10)")
		} else if j == 3 {
			fmt.Println("\nTEST OF LOG(X*X) VS 2*LOG(X)")
		} else {
			fmt.Println("\nTEST OF LOG10(X) VS LOG(X)/LOG(10)")
		}
		fmt.Println()
		fmt.Printf("%7d RANDOM ARGUMENTS WERE TESTED FROM THE INTERVAL\n", n)
		fmt.Printf("      (%.4E, %.4E)\n\n", a, b)
		fmt.Printf(" LOG(X) WAS LARGER %6d TIMES,\n", k1)
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

		if j == 2 {
			a = math.Sqrt(half)
			b = 15.0 / 16.0
		}
	}

	// Special tests
	fmt.Println("\nSPECIAL TESTS")
	fmt.Println()
	fmt.Println(" THE IDENTITY  LOG(X) = -LOG(1/X)  WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X           F(X) + F(1/X)")

	for i := 1; i <= 5; i++ {
		x := rng.Float64()
		x = x + x + 15.0/16.0
		z := math.Log(x) + math.Log(one/x)
		fmt.Printf("  %.7E    %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" TEST OF SPECIAL ARGUMENTS")
	fmt.Println()

	x := one
	y := math.Log(x)
	fmt.Printf(" LOG(1.0) = %.7E\n", y)

	x = mp.XMin
	y = math.Log(x)
	fmt.Printf(" LOG(XMIN) = LOG(%.6E) = %.6E\n", x, y)

	x = mp.XMax
	y = math.Log(x)
	fmt.Printf(" LOG(XMAX) = LOG(%.6E) = %.6E\n", x, y)

	// Test of error returns
	fmt.Println()
	fmt.Println("TEST OF ERROR RETURNS")
	fmt.Println()

	x = -2.0
	fmt.Printf(" LOG WILL BE CALLED WITH THE ARGUMENT %.4E\n", x)
	fmt.Println(" THIS SHOULD RETURN NaN")
	fmt.Println()
	y = math.Log(x)
	fmt.Printf(" LOG RETURNED THE VALUE %.4E\n\n", y)

	x = zero
	fmt.Printf(" LOG WILL BE CALLED WITH THE ARGUMENT %.4E\n", x)
	fmt.Println(" THIS SHOULD RETURN -Inf")
	fmt.Println()
	y = math.Log(x)
	fmt.Printf(" LOG RETURNED THE VALUE %v\n\n", y)

	_ = eight // unused in this version
	fmt.Println(" THIS CONCLUDES THE TESTS")
}
