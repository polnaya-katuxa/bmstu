package model

import (
	"github.com/emirpasic/gods/utils"
)

type event interface {
	timestamp() float64
}

type baseEvent struct {
	ts float64
}

func (e baseEvent) timestamp() float64 {
	return e.ts
}

type eventClientArrived struct {
	baseEvent
}

type eventOperatorProcessed struct {
	baseEvent
	operator   *operator
	toComputer *computer
}

type eventComputerProcessed struct {
	baseEvent
	computer *computer
}

func byTimestamp(a, b interface{}) int {
	return utils.Float64Comparator(a.(event).timestamp(), b.(event).timestamp())
}
