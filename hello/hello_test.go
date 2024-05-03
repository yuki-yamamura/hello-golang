package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Nana")
	want := "Hello, Nana"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}