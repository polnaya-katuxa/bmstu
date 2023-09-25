package algorithms

import (
	"lab_02/internal/matrix"
	"reflect"
	"testing"
)

func TestWinogradMulMatrix(t *testing.T) {
	type args struct {
		m1 matrix.Matrix
		m2 matrix.Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    matrix.Matrix
		wantErr bool
	}{
		{
			name: "test one elem",
			args: args{
				m1: matrix.Matrix{M: 1, N: 1, Data: [][]int{{1}}},
				m2: matrix.Matrix{M: 1, N: 1, Data: [][]int{{1}}},
			},
			want:    matrix.Matrix{M: 1, N: 1, Data: [][]int{{1}}},
			wantErr: false,
		},
		{
			name: "test one elem zero",
			args: args{
				m1: matrix.Matrix{M: 1, N: 1, Data: [][]int{{1}}},
				m2: matrix.Matrix{M: 1, N: 1, Data: [][]int{{0}}},
			},
			want:    matrix.Matrix{M: 1, N: 1, Data: [][]int{{0}}},
			wantErr: false,
		},
		{
			name: "test only zeros",
			args: args{
				m1: matrix.Matrix{M: 2, N: 3, Data: [][]int{{0, 0, 0}, {0, 0, 0}}},
				m2: matrix.Matrix{M: 3, N: 2, Data: [][]int{{0, 0}, {0, 0}, {0, 0}}},
			},
			want:    matrix.Matrix{M: 2, N: 2, Data: [][]int{{0, 0}, {0, 0}}},
			wantErr: false,
		},
		{
			name: "fail sizes",
			args: args{
				m1: matrix.Matrix{M: 2, N: 3, Data: [][]int{{0, 0, 0}, {0, 0, 0}}},
				m2: matrix.Matrix{M: 2, N: 2, Data: [][]int{{0, 0}, {0, 0}}},
			},
			want:    matrix.Matrix{},
			wantErr: true,
		},
		{
			name: "usual test even size",
			args: args{
				m1: matrix.Matrix{M: 4, N: 2, Data: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}},
				m2: matrix.Matrix{M: 2, N: 2, Data: [][]int{{1, 2}, {3, 4}}},
			},
			want:    matrix.Matrix{M: 4, N: 2, Data: [][]int{{7, 10}, {15, 22}, {23, 34}, {31, 46}}},
			wantErr: false,
		},
		{
			name: "usual test odd size",
			args: args{
				m1: matrix.Matrix{M: 3, N: 3, Data: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
				m2: matrix.Matrix{M: 3, N: 1, Data: [][]int{{1}, {2}, {3}}},
			},
			want:    matrix.Matrix{M: 3, N: 1, Data: [][]int{{14}, {32}, {50}}},
			wantErr: false,
		},
		{
			name: "usual test",
			args: args{
				m1: matrix.Matrix{M: 5, N: 4, Data: [][]int{{5, 0, 0, 1}, {9, 7, 0, 8}, {3, 6, 4, 7}, {0, 0, 0, 0}, {2, 2, 8, 2}}},
				m2: matrix.Matrix{M: 4, N: 3, Data: [][]int{{0, 7, 5}, {0, 0, 0}, {0, 9, 5}, {0, 1, 0}}},
			},
			want:    matrix.Matrix{M: 5, N: 3, Data: [][]int{{0, 36, 25}, {0, 71, 45}, {0, 64, 35}, {0, 0, 0}, {0, 88, 50}}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WinogradMulMatrix(tt.args.m1, tt.args.m2)
			if (err != nil) != tt.wantErr {
				t.Errorf("WinogradMulMatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WinogradMulMatrix() got = %v, want %v", got, tt.want)
			}
		})
	}
}
