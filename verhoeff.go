package aadhaar

// Verhoeff tables: d is dihedral group D5 multiplication, p is the position
// permutation, inv is the group inverse.
var d = [10][10]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {1, 2, 3, 4, 0, 6, 7, 8, 9, 5},
	{2, 3, 4, 0, 1, 7, 8, 9, 5, 6}, {3, 4, 0, 1, 2, 8, 9, 5, 6, 7},
	{4, 0, 1, 2, 3, 9, 5, 6, 7, 8}, {5, 9, 8, 7, 6, 0, 4, 3, 2, 1},
	{6, 5, 9, 8, 7, 1, 0, 4, 3, 2}, {7, 6, 5, 9, 8, 2, 1, 0, 4, 3},
	{8, 7, 6, 5, 9, 3, 2, 1, 0, 4}, {9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
}

var p = [8][10]int{
	{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {1, 5, 7, 6, 2, 8, 3, 0, 9, 4},
	{5, 8, 0, 3, 7, 9, 6, 1, 4, 2}, {8, 9, 1, 6, 0, 4, 3, 5, 2, 7},
	{9, 4, 5, 3, 1, 2, 6, 8, 7, 0}, {4, 2, 8, 6, 5, 7, 3, 9, 0, 1},
	{2, 7, 9, 3, 8, 0, 6, 4, 1, 5}, {7, 0, 4, 6, 9, 1, 3, 2, 5, 8},
}

var inv = [10]int{0, 4, 3, 2, 1, 5, 6, 7, 8, 9}

// verhoeffValid reports whether s, consisting only of ASCII digits, passes the
// Verhoeff check (its last digit is the check digit).
func verhoeffValid(s string) bool {
	c := 0
	for i := 0; i < len(s); i++ {
		c = d[c][p[i%8][s[len(s)-1-i]-'0']]
	}
	return c == 0
}

// verhoeffCheckDigit returns the check digit to append to s, which must
// consist only of ASCII digits.
func verhoeffCheckDigit(s string) byte {
	c := 0
	for i := 0; i < len(s); i++ {
		c = d[c][p[(i+1)%8][s[len(s)-1-i]-'0']]
	}
	return byte('0' + inv[c])
}
