package algorithms

import "testing"

func areEqualArrays(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}

	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}

	return true
}

func TestPancakesort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		wait []int
	}{
		{"usual test no duplicates", args{[]int{4, 2, 7, 5, 8, 1, 6}}, []int{1, 2, 4, 5, 6, 7, 8}},
		{"usual test with duplicates", args{[]int{4, 2, 2, 5, 2, 1, 6}}, []int{1, 2, 2, 2, 4, 5, 6}},
		{"test already sorted ascending", args{[]int{2, 4, 5, 6, 8, 10}}, []int{2, 4, 5, 6, 8, 10}},
		{"test already sorted descending", args{[]int{10, 8, 6, 5, 4, 2}}, []int{2, 4, 5, 6, 8, 10}},
		{"test equal elements", args{[]int{2, 2, 2, 2, 2}}, []int{2, 2, 2, 2, 2}},
		{"test one element", args{[]int{2}}, []int{2}},
		{"test many elements", args{[]int{4, 2, 7, 5, 8, 10, 6, 21, 7, 3, 11, 9, 1}}, []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10, 11, 21}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pancakesort(tt.args.arr)
			if !areEqualArrays(tt.args.arr, tt.wait) {
				t.Errorf("Wait: %v\nGot: %v", tt.wait, tt.args.arr)
			}
		})
	}
}

func TestQuicksort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		wait []int
	}{
		{"usual test no duplicates", args{[]int{4, 2, 7, 5, 8, 1, 6}}, []int{1, 2, 4, 5, 6, 7, 8}},
		{"usual test with duplicates", args{[]int{4, 2, 2, 5, 2, 1, 6}}, []int{1, 2, 2, 2, 4, 5, 6}},
		{"test already sorted ascending", args{[]int{2, 4, 5, 6, 8, 10}}, []int{2, 4, 5, 6, 8, 10}},
		{"test already sorted descending", args{[]int{10, 8, 6, 5, 4, 2}}, []int{2, 4, 5, 6, 8, 10}},
		{"test equal elements", args{[]int{2, 2, 2, 2, 2}}, []int{2, 2, 2, 2, 2}},
		{"test one element", args{[]int{2}}, []int{2}},
		{"test many elements", args{[]int{4, 2, 7, 5, 8, 10, 6, 21, 7, 3, 11, 9, 1}}, []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10, 11, 21}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Quicksort(tt.args.arr)
			if !areEqualArrays(tt.args.arr, tt.wait) {
				t.Errorf("Wait: %v\nGot: %v", tt.wait, tt.args.arr)
			}
		})
	}
}

func TestBeadsort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		wait []int
	}{
		{"usual test no duplicates", args{[]int{4, 2, 7, 5, 8, 1, 6}}, []int{1, 2, 4, 5, 6, 7, 8}},
		{"usual test with duplicates", args{[]int{4, 2, 2, 5, 2, 1, 6}}, []int{1, 2, 2, 2, 4, 5, 6}},
		{"test already sorted ascending", args{[]int{2, 4, 5, 6, 8, 10}}, []int{2, 4, 5, 6, 8, 10}},
		{"test already sorted descending", args{[]int{10, 8, 6, 5, 4, 2}}, []int{2, 4, 5, 6, 8, 10}},
		{"test equal elements", args{[]int{2, 2, 2, 2, 2}}, []int{2, 2, 2, 2, 2}},
		{"test one element", args{[]int{2}}, []int{2}},
		{"test many elements", args{[]int{4, 2, 7, 5, 8, 10, 6, 21, 7, 3, 11, 9, 1}}, []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10, 11, 21}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Beadsort(tt.args.arr)
			if !areEqualArrays(tt.args.arr, tt.wait) {
				t.Errorf("Wait: %v\nGot: %v", tt.wait, tt.args.arr)
			}
		})
	}
}
