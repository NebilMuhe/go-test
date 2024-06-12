package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Nebil")
	want := "Hello, Nebil"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
