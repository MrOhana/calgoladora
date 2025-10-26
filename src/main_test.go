package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestValidateUserInput(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "Valid input 1", input: "1", want: true},
		{name: "Valid input 2", input: "2", want: true},
		{name: "Valid input 3", input: "3", want: true},
		{name: "Valid input 4", input: "4", want: true},
		{name: "Valid input 5", input: "5", want: true},
		{name: "Invalid input empty", input: "", want: false},
		{name: "Invalid input text", input: "abc", want: false},
		{name: "Invalid input number", input: "6", want: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidateUserInput(tc.input)
			if got != tc.want {
				t.Errorf("ValidateUserInput(%q) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestPerformOperation(t *testing.T) {
	testCases := []struct {
		name      string
		operation string
		num1      float64
		num2      float64
		want      float64
	}{
		{name: "Sum", operation: "1", num1: 10, num2: 5, want: 15},
		{name: "Subtraction", operation: "2", num1: 10, num2: 5, want: 5},
		{name: "Multiplication", operation: "3", num1: 10, num2: 5, want: 50},
		{name: "Division", operation: "4", num1: 10, num2: 5, want: 2},
		{name: "Division by zero", operation: "4", num1: 10, num2: 0, want: 0},
		{name: "Invalid operation", operation: "6", num1: 10, num2: 5, want: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := PerformOperation(tc.operation, tc.num1, tc.num2)
			if got != tc.want {
				t.Errorf("PerformOperation(%q, %f, %f) = %f; want %f", tc.operation, tc.num1, tc.num2, got, tc.want)
			}
		})
	}
}

func TestGetUserInputNumber(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  float64
	}{
		{name: "Valid number", input: "10", want: 10},
		{name: "Valid float number", input: "10.5", want: 10.5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockStdin(tc.input, func() {
				got := GetUserInputNumber("test message")
				if got != tc.want {
					t.Errorf("GetUserInputNumber() = %f; want %f", got, tc.want)
				}
			})
		})
	}
}

// mockStdin simulates user input for testing functions that read from os.Stdin.
func mockStdin(input string, f func()) {
	oldStdin := os.Stdin
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	// Write the input to the pipe
	go func() {
		defer w.Close()
		io.Copy(w, strings.NewReader(input))
	}()

	f()
}

func TestGetUserInput(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		message string
		want    string
	}{
		{name: "Simple input", input: "hello", message: "Enter something:", want: "hello"},
		{name: "Input with spaces", input: "hello world", message: "Enter something:", want: "hello world"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Redirect stdout to capture the message
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			mockStdin(tc.input, func() {
				got := GetUserInput(tc.message)
				if got != tc.want {
					t.Errorf("GetUserInput() = %q; want %q", got, tc.want)
				}
			})

			// Restore stdout and read the captured output
			w.Close()
			os.Stdout = oldStdout
			var buf bytes.Buffer
			io.Copy(&buf, r)
			// We don't check the output message in this test, just that the input is read correctly.
		})
	}
}