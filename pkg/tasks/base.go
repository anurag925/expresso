package tasks

import (
	"context"
	"expresso/configs/core"
	"expresso/configs/core/initializers"
	"fmt"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

type Task struct {
	app core.App
}

func New(a core.App) *Task {
	return &Task{a}
}

func (s *Task) Logger() initializers.Logger {
	return s.app.Logger().With("task_id", uuid.New())
}

func (s *Task) ScheduleTask() {
	s.app.Tasks().New("print_hello").Periodic("@every 30s").Save()
}

func (s *Task) RegisterTask() {
	s.app.Tasks().AddTask("print_hello", asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		fmt.Printf("hello")
		return nil
	}))
}
