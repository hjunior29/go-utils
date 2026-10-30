package utils

import (
	"reflect"
	"testing"
)

func TestTakeIntsStandard(t *testing.T) {
	input := []int{10, 20, 30, 40, 50}
	got, err := TakeInts(input, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{10, 20, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TakeInts(%v, 3) = %v, want %v", input, got, want)
	}
}

func TestTakeIntsEmptyInput(t *testing.T) {
	input := []int{}
	got, err := TakeInts(input, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TakeInts(%v, 2) = %v, want %v", input, got, want)
	}
}

func TestTakeIntsBoundaryZero(t *testing.T) {
	input := []int{1, 2, 3}
	got, err := TakeInts(input, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TakeInts(%v, 0) = %v, want %v", input, got, want)
	}
}

func TestTakeIntsExceedsLength(t *testing.T) {
	input := []int{5, 6}
	got, err := TakeInts(input, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TakeInts(%v, 5) = %v, want %v", input, got, want)
	}
}

func TestTakeIntsNegativeN(t *testing.T) {
	input := []int{1, 2, 3}
	_, err := TakeInts(input, -1)
	if err == nil {
		t.Errorf("TakeInts(%v, -1) expected error, got nil", input)
	}
}
