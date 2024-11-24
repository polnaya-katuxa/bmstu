package model

import (
	"math"
	mathRand "math/rand"
)

const step = 1e-3

var inf = math.Inf(1)

type TimeModel struct {
	*BaseModel
}

func NewTimeModel(base *BaseModel) *TimeModel {
	return &TimeModel{
		BaseModel: base,
	}
}

func (m *TimeModel) Compute(limit, percent int) (*Result, error) {
	q := newQueue()

	generateTime, err := m.generatorTimer.Rand()
	if err != nil {
		return nil, err
	}

	processedMessages, returnedMessages := 0, 0
	processTime := inf

	for currentTime := 0.0; processedMessages < limit; currentTime += step {
		if currentTime > generateTime {
			// Настало время генерировать новое сообщение в очередь.
			generateTime = inf
			wasEmpty := q.empty()

			q.push()
			generatePeriod, err := m.generatorTimer.Rand()
			if err != nil {
				return nil, err
			}
			generateTime = currentTime + generatePeriod

			if !wasEmpty {
				continue
			}
		} else if currentTime > processTime {
			// Обработка сообщения из очереди.
			processTime = inf

			if q.empty() {
				continue
			}

			q.pop()
			processedMessages++

			if mathRand.Intn(100) < percent {
				q.push()
				returnedMessages++
			}
		} else {
			continue
		}

		if !q.empty() {
			processPeriod, err := m.processorTimer.Rand()
			if err != nil {
				return nil, err
			}
			processTime = currentTime + processPeriod
		}
	}

	return &Result{QueueLen: q.maxLen()}, nil
}
