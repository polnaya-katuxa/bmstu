package model

import (
	"errors"
	"github.com/emirpasic/gods/queues/priorityqueue"
	mathRand "math/rand"
)

var (
	ErrCantGetEvent = errors.New("cant get event from queue")
)

type EventModel struct {
	*BaseModel
}

func NewEventModel(base *BaseModel) *EventModel {
	return &EventModel{
		BaseModel: base,
	}
}

func (m *EventModel) Compute(limit, percent int) (*Result, error) {
	q := newQueue()
	events := priorityqueue.NewWith(byTimestamp)

	generateTime, err := m.generatorTimer.Rand()
	if err != nil {
		return nil, err
	}

	events.Enqueue(event{
		eventType: eventGenerated,
		timestamp: generateTime,
	})

	processedMessages, returnedMessages := 0, 0
	for processedMessages < limit {
		rawEvent, ok := events.Dequeue()
		if !ok {
			return nil, ErrCantGetEvent
		}
		e := rawEvent.(event)

		switch e.eventType {
		case eventGenerated:
			// Настало время генерировать новое сообщение в очередь.
			wasEmpty := q.empty()
			q.push()

			generatePeriod, err := m.generatorTimer.Rand()
			if err != nil {
				return nil, err
			}

			events.Enqueue(event{
				eventType: eventGenerated,
				timestamp: e.timestamp + generatePeriod,
			})

			if !wasEmpty {
				continue
			}
		case eventProcessed:
			if q.empty() {
				continue
			}

			q.pop()
			processedMessages++

			if mathRand.Intn(100) < percent {
				q.push()
				returnedMessages++
			}
		}

		if !q.empty() {
			processPeriod, err := m.processorTimer.Rand()
			if err != nil {
				return nil, err
			}

			events.Enqueue(event{
				eventType: eventProcessed,
				timestamp: e.timestamp + processPeriod,
			})
		}
	}

	return &Result{QueueLen: q.maxLen()}, nil
}
