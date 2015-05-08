package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	var greeting string
	greeting = greet("world")
	if greeting != "hello, world" {
		t.Error("Expected 'hello, world', got ", greeting)
	}
}
