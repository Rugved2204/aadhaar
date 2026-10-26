package aadhaar

import "testing"

func TestVerhoeffKnownVector(t *testing.T) {
	// Classic example from the algorithm's reference description.
	if got := verhoeffCheckDigit("236"); got != '3' {
		t.Fatalf("check digit for 236 = %c, want 3", got)
	}
	if !verhoeffValid("2363") {
		t.Fatal("2363 should be valid")
	}
	if verhoeffValid("2364") {
		t.Fatal("2364 should be invalid")
	}
}
