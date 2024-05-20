package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("mintu")
		want := "Hello, mintu"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("saying hello world when given param is an empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
