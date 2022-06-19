// ==================================================
// test the kernel
// this file will grow ...
// ==================================================

// REF: 
//  create a test: https://tutorialedge.net/golang/intro-testing-in-go/
//  coverage: https://about.codecov.io/blog/getting-started-with-code-coverage-for-golang/

// User guide:  (run in the local src/Kernel directory)
// go test -covermode=atomic -coverprofile=coverage.out

package Kernel

import "testing"

// call testBoot that is defined in the Kernel
func TestLoad(t *testing.T) {
	Bootstrap()
}

// tests float functions 
func TestFloat(t *testing.T) {
	if 2 != INT(F_integer_I_float(2.3)) {t.Error("integer!(2.3)")}
	if 2 != INT(F_integer_I_float(2.9)) {t.Error("integer!(2.9)")}
	if 8.0 != F__exp_float(2.0,3.0) {t.Error("exp(2.0,3.0)")}
	if 2.0 != F_sqrt_float(4.0) {t.Error("sqrt(4.0)")}
}

