package model

import (
	"errors"
	"fyne.io/fyne/v2"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"lab_05/pkg/sdk"
)

type Result struct {
	DeclinePercent float64
}

var (
	ErrCantGetEvent = errors.New("cant get event from queue")
)

type Model struct {
	clientCreator      *ClientCreator
	operatorCreators   []*OperatorCreator
	computerCreators   []*ComputerCreator
	operatorToComputer map[int]int
}

func NewModel(clientCreator *ClientCreator, operatorCreators []*OperatorCreator, computerCreators []*ComputerCreator, operatorToComputer map[int]int) *Model {
	return &Model{clientCreator: clientCreator, operatorCreators: operatorCreators, computerCreators: computerCreators, operatorToComputer: operatorToComputer}
}

func (m *Model) Widget() fyne.CanvasObject {
	grid := sdk.Grid()

	grid.Add(m.clientCreator.Widget())

	for _, o := range m.operatorCreators {
		grid.Add(o.Widget())
	}

	for _, c := range m.computerCreators {
		grid.Add(c.Widget())
	}

	return grid
}

func (m *Model) initComponents() (*client, []*operator, error) {
	client, err := m.clientCreator.create()
	if err != nil {
		return nil, nil, err
	}

	operators := make([]*operator, len(m.operatorCreators))
	for i := range m.operatorCreators {
		operators[i], err = m.operatorCreators[i].create()
		if err != nil {
			return nil, nil, err
		}
	}

	computers := make([]*computer, len(m.computerCreators))
	for i := range m.computerCreators {
		computers[i], err = m.computerCreators[i].create()
		if err != nil {
			return nil, nil, err
		}
	}

	for i := range operators {
		operators[i].toComputer = computers[m.operatorToComputer[i]]
	}

	return client, operators, nil
}

func (m *Model) Compute(limit int) (*Result, error) {
	events := priorityqueue.NewWith(byTimestamp)

	client, operators, err := m.initComponents()
	if err != nil {
		return nil, err
	}

	events.Enqueue(client.planEventClientArrived(baseEvent{ts: 0}))

	generatedMessages, processedMessages, declinedMessages := 0, 0, 0
	for processedMessages+declinedMessages < limit {
		rawEvent, ok := events.Dequeue()
		if !ok {
			return nil, ErrCantGetEvent
		}

		switch e := rawEvent.(type) {
		case eventClientArrived:
			generatedMessages++

			// Клиент пришел обслуживаться.
			// Если есть свободные операторы, то выдаем задачу свободному оператору с наибольшей
			// производительностью, если таковых нет — отклоняем заявку.

			var mostPerformance *operator
			for i := range operators {
				// Если оператор свободен, то сравниваем его среднее время обработки с минимальным.
				// Нужно найти оператора с минимальным временем обработки — это самый производительный оператор.
				if operators[i].busy {
					continue
				}

				if mostPerformance == nil || operators[i].averageTime < mostPerformance.averageTime {
					mostPerformance = operators[i]
				}
			}

			if mostPerformance != nil {
				// Если нашли свободного оператора, надо отдать заявку ему в обработку.
				mostPerformance.busy = true

				// Дали заявку в обработку значит, что через время своей обработки
				// он завершит обработку заявки.
				events.Enqueue(mostPerformance.planEventOperatorProcessed(e))
			} else {
				// Если не нашли свободного оператора, то отклоняем заявку.
				declinedMessages++
			}

			// Если уже сгенерировали заявок больше, чем лимит, то больше не генерируем.
			if generatedMessages >= limit {
				continue
			}

			events.Enqueue(client.planEventClientArrived(e))
		case eventOperatorProcessed:
			// Оператор обработал заявку. Теперь он отправляет ее в накопитель (очередь),
			// из которой заявки будут уходить в компьютеры на обработку.

			e.operator.busy = false

			if e.toComputer.busy {
				// Если компьютер, которому предназначена заявка, занят, то просто кладем в накопитель.
				e.toComputer.queue++
				continue
			}

			// Если свободен, то можем сразу отдать заявку ему в обработку.
			e.toComputer.busy = true
			events.Enqueue(e.toComputer.planEventComputerProcessed(e))
		case eventComputerProcessed:
			processedMessages++

			e.computer.busy = false

			// Компьютер обработал заявку. В этом случае она считается успешно завершенной.
			// Если накопитель к этому компьютеру не пуст, то он берет новую заявку из очереди на обработку.
			// Если пуст, то не берет.
			if e.computer.queue <= 0 {
				continue
			}

			// Берем новую заявку из накопителя.
			e.computer.queue--
			e.computer.busy = true

			// Взяли в обработку, значит компьютер через некоторое время обработает.
			events.Enqueue(e.computer.planEventComputerProcessed(e))
		}
	}

	return &Result{
		DeclinePercent: float64(declinedMessages) / float64(declinedMessages+processedMessages) * 100.0,
	}, nil
}
