# golefunt

Go port of the ELEFUNT elementary function test suite by W.J. Cody for testing the accuracy of standard math library functions.

## Structure

- `machar/` - Machine parameter detection (floating-point characteristics)
- `random/` - Random number generator for test inputs
- `asin/`, `atan/`, `exp/`, `log/`, `power/`, `sincos/`, `sinh/`, `sqrt/`, `tan/`, `tanh/` - Individual test programs for each math function

## Building and Running

```bash
make build          # Build all test programs with version/SHA
make test           # Build and run all tests
make test-asin      # Run a specific test
make clean          # Remove build artifacts
```

From the repository root:
```bash
make build          # Build both Fortran and Go
make test           # Run both test suites
make test-go        # Run Go tests only
```

## Project Files

- `VERSION` - Version number
- `LICENSE` - BSD-style license with attribution
- `Makefile` - Build with version/SHA injection via ldflags

## Comparison with Fortran

The Go and Fortran versions produce comparable results:
- Both detect IEEE 754 double precision (53 base-2 significant digits)
- Both test the same mathematical identities
- Maximum relative errors are similar (~2^-51 to 2^-52)
- Both lose only 1-2 bits of precision in worst case

The Go `math` package achieves the same accuracy as Fortran intrinsics.

## Testing Methodology

ELEFUNT tests use ratio-based identity testing:
- Sample 2000 random arguments from test intervals
- Compare function results against mathematically equivalent expressions
- Report maximum relative error (MRE) and root mean square (RMS) error
- Express errors as powers of 2 to show bits of precision lost

Example identity for sin: `sin(x)` vs `3*sin(x/3) - 4*sin³(x/3)`

## Origin

Ported from the original Fortran ELEFUNT package (Elementary Function Testing) developed by W.J. Cody at Argonne National Laboratory.

See [BIBLIOGRAPHY.md](../BIBLIOGRAPHY.md) for references including CELEFUNT (complex functions).
