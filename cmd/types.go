package cmd

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	ID        string
	Title     string
	Done      bool
	CreatedAt time.Time
	Priority  int
}

func NewTodo(title string, priority int) *Todo {
	id := uuid.NewV4()
	return &Todo{
		ID:        id.String()[:6],
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
		Priority:  priority,
	}
}

func (t *Todo) CreatedTimeInWords() string {
	const (
		Day   = time.Hour * 24
		Month = Day * 30
		Year  = Month * 12
	)

	d := time.Since(t.CreatedAt)

	if d > Year {
		return fmt.Sprintf("%d anos atrás", int(d/Year))
	}
	if d > Month {
		return fmt.Sprintf("%d meses atrás", int(d/Month))
	}
	if d > 2*Day {
		return fmt.Sprintf("%d dias atrás", int(d/Day))
	}
	if d > Day {
		return "ontem"
	}
	return "hoje"
}
