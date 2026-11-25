package utils

import (
	"reflect"
	"testing"
)

func TestFlattenIntsStandard(t *testing.T) {
	input := [][]int{
		{1, 2},
		{3},
		{4, 5, 6},
	}
	expected := []int{1, 2, 3, 4, 5, 6}
	actual := FlattenInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestFlattenIntsEmptyAndNil(t *testing.T) {
	input := [][]int{}
	expected := []int{}
	actual := FlattenInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestFlattenIntsWithEmptyInner(t *testing.T) {
	input := [][]int{
		{},
		{1, 2},
		{},
		{3},
	}
	expected := []int{1, 2, 3}
	actual := FlattenInts(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
