package sdk

import (
	"github.com/charmbracelet/log"
	"sync"
)

type changer interface {
	SetOnValidationChanged(callback func(error))
}

type ValidGroup struct {
	mu      sync.Mutex
	invalid int

	onValid   func()
	onInvalid func()
}

func NewValidGroup(onValid, onInvalid func()) *ValidGroup {
	return &ValidGroup{
		onValid:   onValid,
		onInvalid: onInvalid,
	}
}

func (g *ValidGroup) Add(c changer) {
	c.SetOnValidationChanged(g.onValidationChanged)
}

func (g *ValidGroup) onValidationChanged(err error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	log.Warn("changed validation state", "error", err)

	if err != nil {
		g.invalid++

		if g.invalid == 1 {
			g.onInvalid()
		}
	} else {
		g.invalid--

		if g.invalid < 0 {
			g.invalid = 0
			g.onValid()
		}
	}
}
