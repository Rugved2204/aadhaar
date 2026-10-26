package aadhaar

import (
	"math/rand"
	"testing"
)

func newRand() *rand.Rand { return rand.New(rand.NewSource(1)) }

func TestRoundTrip(t *testing.T) {
	r := newRand()
	for i := 0; i < 10000; i++ {
		n := GenerateTestNumber(r)
		if !Validate(n) {
			t.Fatalf("generated number %q failed Validate", n)
		}
	}
}

func TestSingleDigitErrors(t *testing.T) {
	r := newRand()
	for i := 0; i < 500; i++ {
		n := []byte(GenerateTestNumber(r))
		for pos := 0; pos < len(n); pos++ {
			orig := n[pos]
			for c := byte('0'); c <= '9'; c++ {
				if c == orig {
					continue
				}
				n[pos] = c
				// Changing the first digit to 0/1 is rejected anyway.
				if Validate(string(n)) {
					t.Fatalf("single-digit change %q passed", n)
				}
			}
			n[pos] = orig
		}
	}
}

func TestAdjacentTranspositions(t *testing.T) {
	r := newRand()
	for i := 0; i < 500; i++ {
		n := []byte(GenerateTestNumber(r))
		for pos := 0; pos+1 < len(n); pos++ {
			if n[pos] == n[pos+1] {
				continue
			}
			n[pos], n[pos+1] = n[pos+1], n[pos]
			if Validate(string(n)) {
				t.Fatalf("transposition %q passed", n)
			}
			n[pos], n[pos+1] = n[pos+1], n[pos]
		}
	}
}

func TestValidateRejects(t *testing.T) {
	valid := GenerateTestNumber(newRand())
	cases := map[string]string{
		"empty":        "",
		"short":        valid[:11],
		"long":         valid + "0",
		"letters":      "ABCDEFGHIJKL",
		"leading zero": "0" + valid[1:],
		"leading one":  "1" + valid[1:],
		"unicode":      "٢٣٤٥٦٧٨٩٠١٢٣",
		"nul":          valid[:11] + "\x00",
	}
	for name, in := range cases {
		if Validate(in) {
			t.Errorf("%s: Validate(%q) = true", name, in)
		}
	}
}

func TestNormalizeAndFormats(t *testing.T) {
	n := GenerateTestNumber(newRand())
	spaced := n[:4] + " " + n[4:8] + " " + n[8:]
	hyphen := n[:4] + "-" + n[4:8] + "-" + n[8:]
	for _, in := range []string{n, spaced, hyphen} {
		if !Validate(in) {
			t.Errorf("Validate(%q) = false", in)
		}
		if Normalize(in) != n {
			t.Errorf("Normalize(%q) = %q, want %q", in, Normalize(in), n)
		}
	}
}

func TestMask(t *testing.T) {
	n := GenerateTestNumber(newRand())
	want := "XXXX XXXX " + n[8:]
	if got := Mask(n); got != want {
		t.Errorf("Mask = %q, want %q", got, want)
	}
	if got := Mask("123"); got != "" {
		t.Errorf("Mask(invalid) = %q, want empty", got)
	}
}

func FuzzValidate(f *testing.F) {
	f.Add("")
	f.Add("234567890123")
	f.Add("2345 6789 0123")
	f.Fuzz(func(t *testing.T, s string) {
		if Validate(s) && Mask(s) == "" {
			t.Fatalf("Validate(%q) true but Mask empty", s)
		}
	})
}

func BenchmarkValidate(b *testing.B) {
	n := GenerateTestNumber(newRand())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Validate(n)
	}
}
