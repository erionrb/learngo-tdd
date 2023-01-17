package integers

import "testing"

func TestAdder(t *testing.T) {

	t.Run("Sum positive numbers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})

	t.Run("Sum negative numbers", func(t *testing.T) {
		sum := Add(-2, -2)
		expected := -4

		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})

	t.Run("Sum negative and positive numbers", func(t *testing.T) {
		sum := Add(-2, 4)
		expected := 2

		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})
}
