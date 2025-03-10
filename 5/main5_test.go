package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	tests := []struct {
		name      string
		a         []int
		b         []int
		wantBool  bool
		wantSlice []int
	}{
		{
			name:      "Есть пересечения",
			a:         []int{65, 3, 58, 678, 64},
			b:         []int{64, 2, 3, 43},
			wantBool:  true,
			wantSlice: []int{64, 3},
		},
		{
			name:      "Нет пересечений",
			a:         []int{1, 2, 3},
			b:         []int{4, 5, 6},
			wantBool:  false,
			wantSlice: []int{},
		},
		{
			name:      "Первый слайс пуст",
			a:         []int{},
			b:         []int{1, 2},
			wantBool:  false,
			wantSlice: []int{},
		},
		{
			name:      "Второй слайс пуст",
			a:         []int{1},
			b:         []int{},
			wantBool:  false,
			wantSlice: []int{},
		},
		{
			name:      "Дубликаты в обоих слайсах",
			a:         []int{2, 2, 3},
			b:         []int{2, 3, 3},
			wantBool:  true,
			wantSlice: []int{2, 3},
		},
		{
			name:      "Оба слайса пусты",
			a:         []int{},
			b:         []int{},
			wantBool:  false,
			wantSlice: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool, gotSlice := FindIntersection(tt.a, tt.b)

			if gotBool != tt.wantBool {
				t.Errorf("FindIntersection() gotBool = %v, want %v", gotBool, tt.wantBool)
			}

			sort.Ints(gotSlice)
			sort.Ints(tt.wantSlice)

			if !reflect.DeepEqual(gotSlice, tt.wantSlice) {
				t.Errorf("FindIntersection() gotSlice = %v, want %v", gotSlice, tt.wantSlice)
			}
		})
	}
}
