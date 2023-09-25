package pipeline

import (
	"lab_05/internal/document"
	"lab_05/internal/rule"
	"time"
)

func LaunchLinear(docs []document.Document, rules []rule.Rule, flag int) []document.Document {
	result := make([]document.Document, 0)
	tasks := make([]Task, 0)

	timeStart := time.Now()

	for _, v := range docs {
		input := make(chan Task, len(docs))
		tokenized := make(chan Task, len(docs))
		ruled := make(chan Task, len(docs))
		sorted := make(chan Task, len(docs))

		task := Task{
			Doc:       v,
			TimeFirst: timeStart,
		}

		input <- task
		close(input)

		tokeniser(input, tokenized)
		ruler(tokenized, ruled, rules)
		sorter(ruled, sorted)

		res := <-sorted

		result = append(result, res.Doc)
		tasks = append(tasks, res)
	}

	if flag == 1 {
		LogTasks(tasks)
	}

	return result
}
