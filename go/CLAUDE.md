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

## Project Files

- `VERSION` - Version number
- `LICENSE` - BSD-style license with attribution
- `Makefile` - Build with version/SHA injection via ldflags

## Origin

Ported from the original Fortran ELEFUNT package (Elementary Function Testing) developed by W.J. Cody at Argonne National Laboratory.
