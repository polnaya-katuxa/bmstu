package algorithms

import "testing"

func TestDamerauLevenshtein(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty-empty", args{"", ""}, 0},
		{"usual test", args{"abc", "aba"}, 1},
		{"substitution test", args{"text", "tetx"}, 1},
		{"russian letters test", args{"скат", "кот"}, 2},
		{"empty first string test", args{"", "booterbrod"}, 10},
		{"empty second string test", args{"keelka", ""}, 6},
		{"russian letters empty first string test", args{"", "макароны"}, 8},
		{"russian letters empty second string test", args{"котлета", ""}, 7},
		{"russian letters 2 subs string test", args{"кеотон", "кетоно"}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DamerauLevenshtein(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("DamerauLevenshtein() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLevenshtein(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty-empty", args{"", ""}, 0},
		{"usual test", args{"abc", "aba"}, 1},
		{"substitution test", args{"text", "tetx"}, 2},
		{"russian letters test", args{"скат", "кот"}, 2},
		{"empty first string test", args{"", "ababab"}, 6},
		{"empty second string test", args{"acaca", ""}, 5},
		{"russian letters empty first string test", args{"", "пюрешка"}, 7},
		{"russian letters empty second string test", args{"сосиска", ""}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Levenshtein(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Levenshtein() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursiveDamerauLevenshtein(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty-empty", args{"", ""}, 0},
		{"usual test", args{"abc", "aba"}, 1},
		{"substitution test", args{"text", "tetx"}, 1},
		{"russian letters test", args{"скот", "кот"}, 1},
		{"empty first string test", args{"", "toosovochka"}, 11},
		{"empty second string test", args{"peevo", ""}, 5},
		{"russian letters empty first string test", args{"", "карбюратор"}, 10},
		{"russian letters empty second string test", args{"руль", ""}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecursiveDamerauLevenshtein(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("RecursiveDamerauLevenshtein() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecursiveDamerauLevenshteinCached(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty-empty", args{"", ""}, 0},
		{"usual test", args{"abc", "aba"}, 1},
		{"substitution test", args{"text", "tetx"}, 1},
		{"russian letters test", args{"скат", "кот"}, 2},
		{"empty first string test", args{"", "babooshka"}, 9},
		{"empty second string test", args{"papa", ""}, 4},
		{"russian letters empty first string test", args{"", "пельмени"}, 8},
		{"russian letters empty second string test", args{"лошадь", ""}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecursiveDamerauLevenshteinCached(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("RecursiveDamerauLevenshteinCached() = %v, want %v", got, tt.want)
			}
		})
	}
}
