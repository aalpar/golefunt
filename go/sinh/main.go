// Program to test Sinh/Cosh
// Port of the elefunt sinh.f and dsinh.f test programs by W.J. Cody
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
	b := 0.5
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
				// Test SINH(X) vs identity
				// SINH(3X) = SINH(X)*(3+4*SINH(X)^2)
				y := x / three
				z = math.Sinh(x)
				zz = math.Sinh(y)
				w = one
				if z != zero {
					computed := zz * (three + 4.0*zz*zz)
					w = (z - computed) / z
				}
			} else {
				// Test COSH(X) vs identity
				// COSH(3X) = COSH(X)*(4*COSH(X)^2-3)
				y := x / three
				z = math.Cosh(x)
				zz = math.Cosh(y)
				w = one
				if z != zero {
					computed := zz * (4.0*zz*zz - three)
					w = (z - computed) / z
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
			fmt.Println("\nTEST OF SINH(X) VS 3*SINH(X/3)+4*SINH(X/3)**3")
		} else {
			fmt.Println("\nTEST OF COSH(X) VS 4*COSH(X/3)**3-3*COSH(X/3)")
		}
		fmt.Println()
		fmt.Printf("%7d RANDOM ARGUMENTS WERE TESTED FROM THE INTERVAL\n", n)
		fmt.Printf("      (%.4E, %.4E)\n\n", a, b)

		if j <= 2 {
			fmt.Printf(" SINH(X) WAS LARGER %6d TIMES,\n", k1)
		} else {
			fmt.Printf(" COSH(X) WAS LARGER %6d TIMES,\n", k1)
		}
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

		a = 3.0
		b = math.Log(mp.XMax) - math.Log(3.0)
	}

	// Special tests
	fmt.Println("\nSPECIAL TESTS")
	fmt.Println()
	fmt.Println(" THE IDENTITY   SINH(-X) = -SINH(X)   WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         F(X) + F(-X)")

	for i := 1; i <= 5; i++ {
		x := rng.Float64() * 5.0
		z := math.Sinh(x) + math.Sinh(-x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" THE IDENTITY SINH(X) = X , X SMALL, WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         X - F(X)")

	betap := math.Pow(beta, float64(mp.IT))
	x := rng.Float64() / betap

	for i := 1; i <= 5; i++ {
		z := x - math.Sinh(x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
		x = x / beta
	}

	fmt.Println()
	fmt.Println(" THE IDENTITY   COSH(-X) = COSH(X)   WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         F(X) - F(-X)")

	for i := 1; i <= 5; i++ {
		x := rng.Float64() * 5.0
		z := math.Cosh(x) - math.Cosh(-x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" TEST OF SPECIAL ARGUMENTS")
	fmt.Println()

	x = zero
	y := math.Sinh(x)
	fmt.Printf(" SINH(0.0) = %.7E\n", y)

	y = math.Cosh(zero)
	fmt.Printf(" COSH(0.0) = %.17E (should be 1.0)\n", y)

	// Test of error returns
	fmt.Println()
	fmt.Println("TEST OF ERROR RETURNS")
	fmt.Println()

	x = math.Log(mp.XMax) + 2.0
	fmt.Printf(" SINH WILL BE CALLED WITH THE ARGUMENT %.4E\n", x)
	fmt.Println(" THIS SHOULD OVERFLOW")
	fmt.Println()
	y = math.Sinh(x)
	fmt.Printf(" SINH RETURNED THE VALUE %v\n\n", y)

	fmt.Println(" THIS CONCLUDES THE TESTS")
}
