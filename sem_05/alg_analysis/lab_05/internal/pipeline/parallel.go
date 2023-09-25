package pipeline

import (
	"lab_05/internal/document"
	"lab_05/internal/rule"
	"sort"
	"time"
)

type Task struct {
	Doc document.Document

	TimeFirst time.Time

	TimeStart1 time.Time
	TimeEnd1   time.Time

	TimeStart2 time.Time
	TimeEnd2   time.Time

	TimeStart3 time.Time
	TimeEnd3   time.Time
}

func tokeniser(in <-chan Task, out chan<- Task) {
	for v := range in {
		v.TimeStart1 = time.Now()
		v.Doc.Tokenize()
		v.TimeEnd1 = time.Now()
		out <- v
	}
	close(out)
}

func ruler(in <-chan Task, out chan<- Task, rules []rule.Rule) {
	for v := range in {
		v.TimeStart2 = time.Now()
		v.Doc.ApplyRules(rules)
		v.TimeEnd2 = time.Now()
		out <- v
	}
	close(out)
}

func sorter(in <-chan Task, out chan<- Task) {
	for v := range in {
		v.TimeStart3 = time.Now()
		sort.Strings(v.Doc.Tokens)
		v.TimeEnd3 = time.Now()
		out <- v
	}
	close(out)
}

func LaunchParallel(docs []document.Document, rules []rule.Rule, flag int) []document.Document {
	result := make([]document.Document, 0)
	tasks := make([]Task, 0)

	timeStart := time.Now()

	input := make(chan Task, len(docs))
	tokenized := make(chan Task, len(docs))
	ruled := make(chan Task, len(docs))
	sorted := make(chan Task, len(docs))

	go tokeniser(input, tokenized)
	go ruler(tokenized, ruled, rules)
	go sorter(ruled, sorted)

	for _, v := range docs {
		task := Task{
			Doc:       v,
			TimeFirst: timeStart,
		}
		input <- task
	}

	close(input)

	for v := range sorted {
		result = append(result, v.Doc)
		tasks = append(tasks, v)
	}

	if flag == 1 {
		LogTasks(tasks)
	}

	return result
}
