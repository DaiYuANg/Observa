package scheduler

import (
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("scheduler", fx.Provide(newScheduler), fx.Invoke(lifecycle))

func newScheduler() (gocron.Scheduler, error) {
	return gocron.NewScheduler()
}

func lifecycle(lc fx.Lifecycle, scheduler gocron.Scheduler) {
	lc.Append(fx.StartHook(func() {
		go scheduler.Start()
	}))
}
