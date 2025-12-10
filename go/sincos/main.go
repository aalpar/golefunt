// Program to test Sin/Cos
// Port of the elefunt sin.f and dsin.f test programs by W.J. Cody
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
	b := math.Pi / 2.0 // 1.570796327
	c := b
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
			y := x / three
			y = (x + y) - x
			x = three * y

			var z, zz, w float64
			if j != 3 {
				z = math.Sin(x)
				zz = math.Sin(y)
				w = one
				if z != zero {
					w = (z - zz*(three-4.0*zz*zz)) / z
				}
			} else {
				z = math.Cos(x)
				zz = math.Cos(y)
				w = one
				if z != zero {
					w = (z + zz*(three-4.0*zz*zz)) / z
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

		if j != 3 {
			fmt.Println("\nTEST OF SIN(X) VS 3*SIN(X/3)-4*SIN(X/3)**3")
		} else {
			fmt.Println("\nTEST OF COS(X) VS 4*COS(X/3)**3-3*COS(X/3)")
		}
		fmt.Println()
		fmt.Printf("%7d RANDOM ARGUMENTS WERE TESTED FROM THE INTERVAL\n", n)
		fmt.Printf("      (%.4E, %.4E)\n\n", a, b)

		if j != 3 {
			fmt.Printf(" SIN(X) WAS LARGER %6d TIMES,\n", k1)
		} else {
			fmt.Printf(" COS(X) WAS LARGER %6d TIMES,\n", k1)
		}
		fmt.Printf("           AGREED %6d TIMES, AND\n", k2)
		fmt.Printf("       WAS SMALLER %6d TIMES.\n\n", k3)

		fmt.Printf(" THERE ARE %4d BASE %4d SIGNIFICANT DIGITS IN A FLOATING-POINT NUMBER\n\n", mp.IT, mp.IBeta)

		w := -999.0
		if r6 != zero {
			w = math.Log(math.Abs(r6)) / albeta
		}
		fmt.Printf(" THE MAXIMUM RELATIVE ERROR OF %.4E = %4d ** %7.2f\n", r6, mp.IBeta, w)
		fmt.Printf("    OCCURRED FOR X = %.6E\n", x1)

		w = math.Max(ait+w, zero)
		fmt.Printf(" THE ESTIMATED LOSS OF BASE %4d SIGNIFICANT DIGITS IS %7.2f\n\n", mp.IBeta, w)

		w = -999.0
		if r7 != zero {
			w = math.Log(math.Abs(r7)) / albeta
		}
		fmt.Printf(" THE ROOT MEAN SQUARE RELATIVE ERROR WAS %.4E = %4d ** %7.2f\n", r7, mp.IBeta, w)
		w = math.Max(ait+w, zero)
		fmt.Printf(" THE ESTIMATED LOSS OF BASE %4d SIGNIFICANT DIGITS IS %7.2f\n\n", mp.IBeta, w)

		a = 6.0 * math.Pi // 18.84955592
		if j == 2 {
			a = b + c
		}
		b = a + c
	}

	// Special tests
	fmt.Println("\nSPECIAL TESTS")
	fmt.Println()

	c = one / math.Pow(beta, float64(mp.IT/2))
	z := (math.Sin(a+c) - math.Sin(a-c)) / (c + c)
	fmt.Printf(" IF %.6E IS NOT ALMOST 1.0,    SIN HAS THE WRONG PERIOD.\n\n", z)

	fmt.Println(" THE IDENTITY   SIN(-X) = -SIN(X)   WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         F(X) + F(-X)")

	for i := 1; i <= 5; i++ {
		x := rng.Float64() * a
		z := math.Sin(x) + math.Sin(-x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" THE IDENTITY SIN(X) = X , X SMALL, WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         X - F(X)")

	betap := math.Pow(beta, float64(mp.IT))
	x := rng.Float64() / betap

	for i := 1; i <= 5; i++ {
		z := x - math.Sin(x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
		x = x / beta
	}

	fmt.Println()
	fmt.Println(" THE IDENTITY   COS(-X) = COS(X)   WILL BE TESTED.")
	fmt.Println()
	fmt.Println("        X         F(X) - F(-X)")

	for i := 1; i <= 5; i++ {
		x := rng.Float64() * a
		z := math.Cos(x) - math.Cos(-x)
		fmt.Printf("  %.7E  %.7E\n", x, z)
	}

	fmt.Println()
	fmt.Println(" TEST OF UNDERFLOW FOR VERY SMALL ARGUMENT.")
	expon := float64(mp.MinExp) * 0.75
	x = math.Pow(beta, expon)
	y := math.Sin(x)
	fmt.Printf("\n      SIN(%.6E) = %.6E\n", x, y)

	fmt.Println()
	fmt.Println(" THE FOLLOWING THREE LINES ILLUSTRATE THE LOSS IN SIGNIFICANCE")
	fmt.Println(" FOR LARGE ARGUMENTS.  THE ARGUMENTS ARE CONSECUTIVE.")

	z = math.Sqrt(betap)
	x = z * (one - mp.EpsNeg)
	y = math.Sin(x)
	fmt.Printf("\n      SIN(%.6E) = %.6E\n", x, y)
	y = math.Sin(z)
	fmt.Printf("\n      SIN(%.6E) = %.6E\n", z, y)
	x = z * (one + mp.Eps)
	y = math.Sin(x)
	fmt.Printf("\n      SIN(%.6E) = %.6E\n", x, y)

	// Test of error returns
	fmt.Println()
	fmt.Println("TEST OF ERROR RETURNS")
	fmt.Println()
	x = betap
	fmt.Printf(" SIN WILL BE CALLED WITH THE ARGUMENT %.4E\n", x)
	fmt.Println(" THIS SHOULD NOT TRIGGER AN ERROR IN GO (NO ARGRED)")
	fmt.Println()
	y = math.Sin(x)
	fmt.Printf(" SIN RETURNED THE VALUE %.4E\n\n", y)

	fmt.Println(" THIS CONCLUDES THE TESTS")
}
