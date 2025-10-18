package main

import "testing"

func TestGetGreeting(t *testing.T) {
	expected := "Hello, World!"
	actual := getGreeting()

	if actual != expected {
		t.Errorf("getGreeting() = %q, want %q", actual, expected)
	}
}

func BenchmarkGetGreeting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getGreeting()
	}
}
