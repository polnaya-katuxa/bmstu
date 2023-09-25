package query

import (
	"testing"
)

func TestIsCorrectQuery(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name    string
		args    args
		want1   int
		want2   int
		wantErr bool
	}{
		{
			name:    "usual test 1",
			args:    args{[]string{"вывести", "котик", "не", "пушистый"}},
			want1:   0,
			want2:   7500,
			wantErr: false,
		},
		{
			name:    "fail test 1",
			args:    args{[]string{"вывести", "котик", "котёнок", "не", "пушистый"}},
			want1:   -1,
			want2:   -1,
			wantErr: true,
		},
		{
			name:    "fail test 2",
			args:    args{[]string{"быстро", "котик", "не", "пушистый"}},
			want1:   -1,
			want2:   -1,
			wantErr: true,
		},
		{
			name:    "usual test 2",
			args:    args{[]string{"дать", "кошечка", "очень", "пушистый"}},
			want1:   11668,
			want2:   17500,
			wantErr: false,
		},
		{
			name:    "usual test 3",
			args:    args{[]string{"дать", "очень", "пушистый", "кошечка"}},
			want1:   11668,
			want2:   17500,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2, got3 := IsCorrectQuery(tt.args.tokens)
			if got1 != tt.want1 {
				t.Errorf("IsCorrectQuery() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("IsCorrectQuery() got2 = %v, want %v", got2, tt.want2)
			}
			if (got3 != nil) != tt.wantErr {
				t.Errorf("IsCorrectQuery() got3 = %v, want %v", got3, tt.wantErr)
			}
		})
	}
}
