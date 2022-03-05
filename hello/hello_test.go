package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	t.Run("Greeting Enigma", func(t *testing.T) {
		got := Hello("Enigma")
		expected := "Hello Enigma!"

		if got != expected {
			t.Errorf("Got: %s, Expected: %s", got, expected)
		}
	})

	t.Run("Greeting Reza", func(t *testing.T) {
		got := Hello("Reza")
		expected := "Hello Reza!"

		if got != expected {
			t.Errorf("Got: %s, Expected: %s", got, expected)
		}
	})
}
