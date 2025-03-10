package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 4}},
		{[]int{0, 5, 10}, []int{0, 10}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		got := sliceExample(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("sliceExample(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	want := []int{1, 2, 3, 4}
	got := addElements(input, 4)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("addElements(%v, 4) = %v, want %v", input, got, want)
	}
}

func TestCopySlice(t *testing.T) {
	original := []int{1, 2, 3}
	copied := copySlice(original)
	original[0] = 100
	if copied[0] == 100 {
		t.Error("Copied slice should not be affected by original changes")
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		input []int
		index int
		want  []int
		err   string
	}{
		{[]int{1, 2, 3, 4}, 1, []int{1, 3, 4}, ""},
		{[]int{5}, 0, []int{}, ""},
		{[]int{1, 2}, 2, nil, "index out of range"},
	}

	for _, tt := range tests {
		got, err := removeElement(tt.input, tt.index)
		if tt.err != "" {
			if err == nil || err.Error() != tt.err {
				t.Errorf("removeElement(%v, %d) error = %v, want %v", tt.input, tt.index, err, tt.err)
			}
		} else {
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElement(%v, %d) = %v, want %v", tt.input, tt.index, got, tt.want)
			}
		}
	}
}
