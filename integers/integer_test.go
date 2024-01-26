package integers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(1, 3)

	expected := 4

	if sum != expected {
		t.Errorf("Sum  %d is different than %d", sum, expected)
	}
}

func TestMultiply(t *testing.T) {

	product := Multiply(1, 3)

	expected := 3

	if product != expected {
		t.Errorf("product %d is different than %d", product, expected)
	}
}
