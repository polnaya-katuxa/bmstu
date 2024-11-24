package model

import "github.com/emirpasic/gods/utils"

type Result struct {
	QueueLen int
}

type queue struct {
	cur int
	max int
}

func newQueue() *queue {
	return &queue{}
}

func (q *queue) push() {
	q.cur++
	if q.cur > q.max {
		q.max = q.cur
	}
}

func (q *queue) pop() {
	if q.cur > 0 {
		q.cur--
	}
}

func (q *queue) empty() bool {
	return q.cur == 0
}

func (q *queue) maxLen() int {
	return q.max
}

type eventType int

const (
	eventGenerated eventType = iota + 1
	eventProcessed
)

type event struct {
	eventType eventType
	timestamp float64
}

func byTimestamp(a, b interface{}) int {
	return utils.Float64Comparator(a.(event).timestamp, b.(event).timestamp)
}
