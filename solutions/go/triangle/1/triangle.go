// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle


// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
    // Pick values for the following identifiers used by the test program.
    NaT Kind = "not a triangle"// not a triangle
    Equ Kind = "equilateral"// equilateral
    Iso Kind = "isosceles"// isosceles
    Sca Kind = "scalene"// scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	k = NaT
    if a == 0 || b == 0 || c == 0 {
        return k
    }
    
	if a+b >= c && b+c >= a && c+a >= b {
		switch {
		case a == b && c == a:
			k = Equ
		case a == b || b == c || c == a:
			k = Iso
		case a != b && b != c && c != a:
			k = Sca
		}
	}

	return k
}
