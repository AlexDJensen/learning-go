package greetings

import (
	"regexp"
	"testing"
)

// This test calls greetings with a name and checks return value.
func TestHelloName(t *testing.T) {
	name := "Testerman"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Testerman")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Testerman") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty tests hello function without giving an actual name.
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
