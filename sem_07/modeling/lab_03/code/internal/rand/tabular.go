package rand

import (
	_ "embed"
	"strconv"
	"time"
)

//go:embed table.txt
var table []byte

type Tabular struct {
	n        int
	position int
}

func NewTabular(n int) *Tabular {
	return &Tabular{
		n:        n,
		position: int(time.Now().UnixNano()) % len(table),
	}
}

func (t *Tabular) movePosition() {
	t.position = (t.position + 1) % len(table)
}

func (t *Tabular) Rand() int {
	// Если нужно однозначное число, то просто берем цифру из таблицы и двигаем позицию.
	if t.n == 1 {
		result := int(table[t.position] - '0')
		t.movePosition()
		return result
	}

	// Если больше, чем однозначное, то читаем кусок из таблицы, пропуская ведущие нули.
	resultBytes := make([]byte, t.n)
	// Пропускаем ведущие нули, пока первый символ 0 — идем дальше.
	for table[t.position] == '0' {
		t.movePosition()
	}
	resultBytes[0] = table[t.position]

	for i := 1; i < len(resultBytes); i++ {
		resultBytes[i] = table[t.position]
		t.movePosition()
	}

	result, _ := strconv.Atoi(string(resultBytes))

	return result
}
