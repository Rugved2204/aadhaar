// Package aadhaar validates Indian Aadhaar numbers using the Verhoeff checksum.
//
// A passing result means the number is well-formed: 12 digits, not starting
// with 0 or 1, with a correct Verhoeff check digit. It does NOT mean the
// number has been issued or belongs to anyone; only UIDAI can confirm that.
//
// The package performs no I/O and never stores or logs its input.
package aadhaar

import (
	"math/rand"
	"strings"
)

const length = 12

var normalizer = strings.NewReplacer(" ", "", "-", "")

// Normalize strips spaces and hyphens, e.g. "2345 6789 0123" becomes
// "234567890123".
func Normalize(s string) string {
	return normalizer.Replace(s)
}

// Validate reports whether s is a structurally valid Aadhaar number. Spaces
// and hyphens are ignored.
func Validate(s string) bool {
	s = Normalize(s)
	if len(s) != length || s[0] < '2' {
		return false
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return verhoeffValid(s)
}

// Mask returns the number in the form "XXXX XXXX 0123", or "" if s is not a
// valid Aadhaar number.
func Mask(s string) string {
	s = Normalize(s)
	if !Validate(s) {
		return ""
	}
	return "XXXX XXXX " + s[8:]
}

// GenerateTestNumber returns a random number that passes Validate. It is meant
// for tests and fixtures only: such numbers are not issued Aadhaar numbers.
// If r is nil, the global math/rand source is used.
func GenerateTestNumber(r *rand.Rand) string {
	intn := rand.Intn
	if r != nil {
		intn = r.Intn
	}
	b := make([]byte, length-1)
	b[0] = byte('2' + intn(8))
	for i := 1; i < len(b); i++ {
		b[i] = byte('0' + intn(10))
	}
	s := string(b)
	return s + string(verhoeffCheckDigit(s))
}
