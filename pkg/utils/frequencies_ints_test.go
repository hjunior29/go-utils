package utils

import (
	"reflect"
	"testing"
)

func TestFrequenciesIntsBasic(t *testing.T) {
	input := []int{1, 2, 2, 3, 3, 3}
	expected := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	actual := FrequenciesInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestFrequenciesIntsEmpty(t *testing.T) {
	input := []int{}
	expected := map[int]int{}
	actual := FrequenciesInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestFrequenciesIntsNegatives(t *testing.T) {
	input := []int{-1, -1, 0, 5, 5, 5}
	expected := map[int]int{
		-1: 2,
		0:  1,
		5:  3,
	}
	actual := FrequenciesInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
