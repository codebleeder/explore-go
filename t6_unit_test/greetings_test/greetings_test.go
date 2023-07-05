package greetings_test

import (
	"greetings"
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := greetings.Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := greetings.Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
