package main

import "testing"

func TestCalcRoundResultLoose(t *testing.T) {
	result := calcResult(100, 4, [3]int{1, 2, 3})
	expected := -100
	if result != expected {
		t.Errorf("CalcRoundResult(100, 4, [1,2,3]) = %d want %d", result, expected)
	}
}

func TestCalcRoundResultWinOneRight(t *testing.T) {
	result := calcResult(100, 4, [3]int{1, 2, 4})
	expected := 100
	if result != expected {
		t.Errorf("CalcRoundResult(100, 4, [1,2,4]) = %d want %d", result, expected)
	}
}
func TestCalcRoundResultWinThreeRight(t *testing.T) {
	result := calcResult(100, 4, [3]int{4, 4, 4})
	expected := 300
	if result != expected {
		t.Errorf("CalcRoundResult(100, 4, [4,4, 4]) = %d want %d", result, expected)
	}
}
