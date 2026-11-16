package utils

import (
	"reflect"
	"testing"
)

func TestRotateIntsBasic(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	got, err := RotateInts(nums, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{3, 4, 5, 1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RotateInts(%v, 2) = %v, want %v", nums, got, want)
	}

	if !reflect.DeepEqual(nums, []int{1, 2, 3, 4, 5}) {
		t.Errorf("input was modified: %v", nums)
	}
}

func TestRotateIntsEmpty(t *testing.T) {
	nums := []int{}
	got, err := RotateInts(nums, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RotateInts(%v, 3) = %v, want %v", nums, got, want)
	}
}

func TestRotateIntsBoundary(t *testing.T) {
	nums := []int{10, 20, 30}
	got, err := RotateInts(nums, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{10, 20, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RotateInts(%v, 3) = %v, want %v", nums, got, want)
	}

	gotLarge, err := RotateInts(nums, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	wantLarge := []int{30, 10, 20}
	if !reflect.DeepEqual(gotLarge, wantLarge) {
		t.Errorf("RotateInts(%v, 5) = %v, want %v", nums, gotLarge, wantLarge)
	}
}

func TestRotateIntsInvalid(t *testing.T) {
	nums := []int{1, 2, 3}
	_, err := RotateInts(nums, -1)
	if err == nil {
		t.Errorf("expected error for negative offset, got nil")
	}
}
