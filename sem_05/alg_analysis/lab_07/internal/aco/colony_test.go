package aco

import (
	"lab_07/internal/graph"
	"testing"
)

func TestTravellingSalesmanACO(t *testing.T) {
	type args struct {
		g     graph.Graph
		alpha float64
		beta  float64
		k     float64
		time  int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		want1   int
		wantErr bool
	}{
		{
			name: "1x1 nill matrix",
			args: args{
				g: graph.Graph{
					Size:       1,
					Connection: [][]int{{0}},
				},
				alpha: 0.5,
				beta:  0.5,
				k:     0.5,
				time:  10,
			},
			want:    []int{},
			want1:   0,
			wantErr: false,
		},
		{
			name: "2x2 matrix",
			args: args{
				g: graph.Graph{
					Size:       2,
					Connection: [][]int{{0, 34}, {34, 0}},
				},
				alpha: 0.5,
				beta:  0.5,
				k:     0.5,
				time:  10,
			},
			want:    []int{0, 1},
			want1:   68,
			wantErr: false,
		},
		{
			name: "3x3 matrix",
			args: args{
				g: graph.Graph{
					Size:       3,
					Connection: [][]int{{0, 24, 6}, {24, 0, 13}, {6, 13, 0}},
				},
				alpha: 0.5,
				beta:  0.5,
				k:     0.5,
				time:  10,
			},
			want:    []int{0, 1, 2},
			want1:   43,
			wantErr: false,
		},
		{
			name: "5x5 matrix",
			args: args{
				g: graph.Graph{
					Size:       5,
					Connection: [][]int{{0, 12, 5, 23, 56}, {12, 0, 31, 4, 13}, {5, 31, 0, 8, 2}, {23, 4, 8, 0, 11}, {56, 13, 2, 11, 0}},
				},
				alpha: 0.5,
				beta:  0.5,
				k:     0.5,
				time:  10,
			},
			want:    []int{3, 1, 0, 2, 4},
			want1:   34,
			wantErr: false,
		},
		{
			name: "8x8 matrix",
			args: args{
				g: graph.Graph{
					Size: 8,
					Connection: [][]int{{0, 45, 12, 67, 88, 22, 14, 4}, {45, 0, 5, 7, 89, 34, 121, 7},
						{12, 5, 0, 23, 45, 32, 43, 12}, {67, 7, 23, 0, 44, 44, 32, 2}, {88, 89, 45, 44, 0, 4, 56, 21},
						{22, 34, 32, 44, 4, 0, 22, 47}, {14, 121, 43, 32, 56, 22, 0, 9}, {4, 7, 12, 2, 21, 47, 9, 0}},
				},
				alpha: 0.5,
				beta:  0.5,
				k:     0.5,
				time:  10,
			},
			want:    []int{5, 6, 7, 3, 1, 4, 0, 6},
			want1:   87,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1, err := TravellingSalesmanACO(tt.args.g, tt.args.alpha, tt.args.beta, tt.args.k, tt.args.time)
			if (err != nil) != tt.wantErr {
				t.Errorf("TravellingSalesmanACO() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("TravellingSalesmanACO() got = %v, want %v", got, tt.want)
			//}
			if got1 != tt.want1 {
				t.Errorf("TravellingSalesmanACO() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
