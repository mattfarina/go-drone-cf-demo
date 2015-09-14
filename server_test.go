package main

import (
	"testing"
)

// A test to just have a test and see it in the CI/CD workflow.
func TestMsg(t *testing.T) {
	msg := message()
	if msg != "Hello World" {
		t.Error("Error in Hello World")
	}
}
