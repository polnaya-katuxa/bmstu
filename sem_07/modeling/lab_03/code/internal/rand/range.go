package rand

// Range диапазон целых чисел [Start; End).
type Range struct {
	Start, End int
}

// Ranges диапазоны чисел. В данном случае однозначные, двузначные, трехзначные числа.
var Ranges = map[int]Range{
	1: Range{
		Start: 0,
		End:   10,
	},
	2: Range{
		Start: 10,
		End:   100,
	},
	3: Range{
		Start: 100,
		End:   1000,
	},
}
