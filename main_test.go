package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false", 0)
	}

	if msg != "0 is not prime by definition" {
		t.Error("wrong message returned:", msg)
	}

	result, msg = isPrime(7)
	if !result {
		t.Errorf("with %d as test parameter, got false, but expected true", 7)
	}

	if msg != "7 is a prime number" {
		t.Error("wrong message returned:", msg)
	}
}

func Test_isPrime_withTable(t *testing.T) {
	testCases := []struct {
		n              int
		expectedResult bool
		expectedMsg    string
		description    string
	}{
		{n: 0, expectedResult: false, expectedMsg: "0 is not prime by definition", description: "testing 0 not prime by definition"},
		{n: 1, expectedResult: false, expectedMsg: "1 is not prime by definition", description: "testing 1 not prime by definition"},
		{n: 2, expectedResult: true, expectedMsg: "2 is a prime number", description: "testing 2 is a prime"},
		{n: 7, expectedResult: true, expectedMsg: "7 is a prime number", description: "testing 7 is a prime"},
		{n: 8, expectedResult: false, expectedMsg: "8 is not a prime number because it is divisible by 2", description: "testing 8 not prime because it is divisible"},
		{n: -7, expectedResult: false, expectedMsg: "Negative numbers are not prime by definition", description: "testing -7 not prime because it is negative"},
	}

	for _, e := range testCases {
		t.Run(e.description, func(t *testing.T) {
			actualResult, actualMsg := isPrime(e.n)
			if e.expectedResult != actualResult {
				t.Errorf("%s: for number %d got: %t but wanted: %t", e.description, e.n, actualResult, e.expectedResult)
			}
			if e.expectedMsg != actualMsg {
				t.Errorf("%s: for number %d got: %s but wanted: %s", e.description, e.n, actualMsg, e.expectedMsg)
			}
		})
	}
}

func Test_prompt(t *testing.T) {
	stdOutOld := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	prompt()
	w.Close()
	b, _ := io.ReadAll(r)
	wanted := "-> "
	if string(b) != wanted {
		t.Errorf("wanted %s but got %s", wanted, string(b))
	}
	os.Stdout = stdOutOld
}

func Test_intro(t *testing.T) {
	stdOutOld := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	intro()
	w.Close()
	b, _ := io.ReadAll(r)
	wanted := "Enter a whole number"
	if !strings.Contains(string(b), wanted) {
		t.Errorf("wanted %s in the output but got %s", wanted, string(b))
	}
	os.Stdout = stdOutOld
}
