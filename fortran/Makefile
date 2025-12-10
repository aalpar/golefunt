LIB=elefunt
FC=gfortran
FFLAGS=-O -std=legacy -w

# Source files for tests (each is a standalone program that needs machar.f)
SINGLE_TESTS = asin atan exp alog sin sinh sqrt tan tanh power
DOUBLE_TESTS = dasin datan dexp dlog dsin dsinh dsqrt dtan dtanh dpower

all: $(SINGLE_TESTS) $(DOUBLE_TESTS)

# Preprocess machar.f for single precision (uncomment CS lines)
machar_single.f: machar.f
	sed 's/^CS/  /' machar.f > machar_single.f

# Single precision tests (need preprocessed machar_single.f)
asin: asin.f machar_single.f
	$(FC) $(FFLAGS) -o $@ $^

atan: atan.f machar_single.f
	$(FC) $(FFLAGS) -o $@ $^

exp: exp.f machar_single.f
	$(FC) $(FFLAGS) -o $@ $^

alog: alog.f machar_single.f
	$(FC) $(FFLAGS) -o $@ $^

sin: sin.f machar_single.f
	$(FC) $(FFLAGS) -o $@ $^

sinh: sinh.f machar_single.f
	$(FC) $(FFLAGS) -o $@ $^

sqrt: sqrt.f machar_single.f
	$(FC) $(FFLAGS) -o $@ $^

tan: tan.f machar_single.f
	$(FC) $(FFLAGS) -o $@ $^

tanh: tanh.f machar_single.f
	$(FC) $(FFLAGS) -o $@ $^

power: power.f machar_single.f
	$(FC) $(FFLAGS) -o $@ $^

# Double precision tests (need dmachar.f)
dasin: dasin.f dmachar.f
	$(FC) $(FFLAGS) -o $@ $^

datan: datan.f dmachar.f
	$(FC) $(FFLAGS) -o $@ $^

dexp: dexp.f dmachar.f
	$(FC) $(FFLAGS) -o $@ $^

dlog: dlog.f dmachar.f
	$(FC) $(FFLAGS) -o $@ $^

dsin: dsin.f dmachar.f
	$(FC) $(FFLAGS) -o $@ $^

dsinh: dsinh.f dmachar.f
	$(FC) $(FFLAGS) -o $@ $^

dsqrt: dsqrt.f dmachar.f
	$(FC) $(FFLAGS) -o $@ $^

dtan: dtan.f dmachar.f
	$(FC) $(FFLAGS) -o $@ $^

dtanh: dtanh.f dmachar.f
	$(FC) $(FFLAGS) -o $@ $^

dpower: dpower.f dmachar.f
	$(FC) $(FFLAGS) -o $@ $^

# Run all tests
test: all
	@echo "Running single precision tests..."
	@for t in $(SINGLE_TESTS); do echo "=== $$t ===" && ./$$t; done
	@echo "Running double precision tests..."
	@for t in $(DOUBLE_TESTS); do echo "=== $$t ===" && ./$$t; done

clean:
	rm -f $(SINGLE_TESTS) $(DOUBLE_TESTS) *.o machar_single.f

.PHONY: all test clean
