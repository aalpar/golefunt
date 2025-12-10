# ELEFUNT

Go port of the ELEFUNT elementary function test suite by W.J. Cody for testing the accuracy of standard math library functions.

## Overview

ELEFUNT (Elementary Function Testing) is a classic test suite developed at Argonne National Laboratory for evaluating the accuracy of implementations of elementary mathematical functions (sin, cos, exp, log, sqrt, etc.).

This repository contains:
- The original Fortran test programs
- A Go port of the test suite

## Building

```bash
make build          # Build both Fortran and Go versions
make build-fortran  # Build only Fortran
make build-go       # Build only Go
```

## Running Tests

```bash
make test           # Run both Fortran and Go test suites
make test-fortran   # Run Fortran tests only
make test-go        # Run Go tests only
```

## Project Structure

```
elefunt/
├── fortran/        # Original Fortran test programs
│   ├── *.f         # Single and double precision tests
│   └── Makefile
├── go/             # Go port
│   ├── machar/     # Machine parameter detection
│   ├── random/     # Random number generator
│   ├── asin/       # Asin/Acos test
│   ├── atan/       # Atan/Atan2 test
│   ├── exp/        # Exp test
│   ├── log/        # Log test
│   ├── power/      # Power (x^y) test
│   ├── sincos/     # Sin/Cos test
│   ├── sinh/       # Sinh/Cosh test
│   ├── sqrt/       # Sqrt test
│   ├── tan/        # Tan test
│   ├── tanh/       # Tanh test
│   └── Makefile
└── Makefile
```

## Attribution

### Original Fortran Code

W.J. Cody, Argonne National Laboratory

Reference: W.J. Cody and W. Waite, "Software Manual for the Elementary Functions", Prentice-Hall, Englewood Cliffs, NJ, 1980.

The original ELEFUNT package is available from [Netlib](https://www.netlib.org/elefunt/).

### Go Port

The Go port was created by Aaron Alpar with assistance from [Claude Code](https://claude.ai/code).

## License

See [LICENSE](go/LICENSE) for details. The original Fortran code is in the public domain. The Go port is released under a BSD-style license.
