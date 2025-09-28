package main

import (
	"github/DaiYuANg/Observa/internal/gateway"
	"github/DaiYuANg/Observa/internal/injector"
	"github/DaiYuANg/Observa/internal/worker"
	"github/DaiYuANg/Observa/modules/nats"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type AppContainer struct {
}

func NewAppContainer() *AppContainer {
	return &AppContainer{}
}

func (c *AppContainer) Run() error {
	container, _, err := injector.CreateContainer(nats.ServerModule)
	if err != nil {
		return err
	}
	go container.Run()
	time.Sleep(2 * time.Second)
	// 顺序启动
	g, err := gateway.NewGateway()
	if err != nil {
		return err
	}
	go g.GetApp().Run()
	w, err := worker.NewWorker()
	// 等待退出信号
	if err != nil {
		return err
	}
	go w.GetApp().Run()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	return nil
}
