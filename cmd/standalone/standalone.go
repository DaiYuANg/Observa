package main

import (
	"context"
	"github/DaiYuANg/Observa/modules/nats"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
)

type AppContainer struct {
	Apps []*fx.App
}

func NewAppContainer(apps ...*fx.App) *AppContainer {
	return &AppContainer{Apps: apps}
}

func (c *AppContainer) Run() error {
	fx.New(nats.ServerModule).Start(context.Background())
	// 顺序启动
	for _, app := range c.Apps {
		go app.Run() // fx.App 自己会阻塞，所以这里用 goroutine
	}
	// 等待退出信号
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	return nil
}
