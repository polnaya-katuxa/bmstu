package query

import (
	"reflect"
	"testing"
)

func TestNormalize(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "объекты",
			args: args{[]string{"кошечка", "котёночек", "котик", "коты", "котята", "плешивая"}},
			want: []string{"кошечка", "котёночек", "котик", "кот", "котёнок", "плешивый"},
		},
		{
			name: "запросы",
			args: args{[]string{"выведи", "дай", "показывай", "покажи", "какие", "выдай", "хочу", "найти", "найди"}},
			want: []string{"вывести", "дать", "показывать", "показать", "какой", "выдать", "хотеть", "найти", "найти"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Normalize(tt.args.tokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenize(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "usual test",
			args: args{"здравствуйте, как дела? что делаете?    какая-то собака!!!"},
			want: []string{"здравствуйте", "как", "дела", "что", "делаете", "какая-то", "собака"},
		},
		{
			name: "1 token",
			args: args{"здравствуйте!"},
			want: []string{"здравствуйте"},
		},
		{
			name: "no tokens",
			args: args{""},
			want: []string{},
		},
		{
			name: "no tokens only sym",
			args: args{"!!!  .  - ?"},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Tokenize(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
