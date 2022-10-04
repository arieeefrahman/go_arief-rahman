package calculation

import "testing"

func TestAddition(t *testing.T) {
	var expected int = 3
	var actual int = Addition(1, 2)
	
	if expected != actual {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
}

func TestSubtraction(t *testing.T) {
	var expected int = -1
	var actual int = Subtraction(1, 2)
	
	if expected != actual {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
}

func TestDivision(t *testing.T) {
	var expected int = 2
	var actual int = Division(8, 4)
	
	if expected != actual {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
}

func TestMultiplication(t *testing.T) {
	var expected int = 10
	var actual int = Multiplication(5, 2)
	
	if expected != actual {
		t.Errorf("Expected: %d, got: %d", expected, actual)
	}
}