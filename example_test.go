package aadhaar_test

import (
	"fmt"
	"math/rand"

	"github.com/Rugved2204/aadhaar"
)

func ExampleValidate() {
	n := aadhaar.GenerateTestNumber(rand.New(rand.NewSource(42)))
	fmt.Println(aadhaar.Validate(n))
	fmt.Println(aadhaar.Validate("123456789012"))
	// Output:
	// true
	// false
}

func ExampleMask() {
	n := aadhaar.GenerateTestNumber(rand.New(rand.NewSource(42)))
	fmt.Println(aadhaar.Mask(n)[:10])
	// Output: XXXX XXXX
}
