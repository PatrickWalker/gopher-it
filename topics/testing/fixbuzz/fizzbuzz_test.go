package main

import "testing"

func TestFizzbuzz3(t *testing.T) {
	out := fizzbuzz(3)
	if out != "fizz" {
		// this is how we fail the test
		t.Errorf("Output was incorrect, got: %v, want: %v.", out, "fizz")
	}
}

func TestFizzbuzz5(t *testing.T) {
	out := fizzbuzz(5)
	if out != "buzz" {
		// this is how we fail the test
		t.Errorf("Output was incorrect, got: %v, want: %v.", out, "buzz")
	}
}

func TestFizzbuzz15(t *testing.T) {
	out := fizzbuzz(15)
	if out != "fizzbuzz" {
		// this is how we fail the test
		t.Errorf("Output was incorrect, got: %v, want: %v.", out, "fizzbuzz")
	}
}

func TestFizzbuzz1(t *testing.T) {
	out := fizzbuzz(1)
	if out != "1" {
		// this is how we fail the test
		t.Errorf("Output was incorrect, got: %v, want: %v.", out, "1")
	}
}
