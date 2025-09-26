package main

import (
	"github/DaiYuANg/Observa/internal/gateway/router"
	"github/DaiYuANg/Observa/modules/http"
	"github/DaiYuANg/Observa/modules/logger"
	"github/DaiYuANg/Observa/modules/nats"
	"github/DaiYuANg/Observa/modules/scheduler"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var rootCmd = &cobra.Command{Run: func(cmd *cobra.Command, args []string) {
	fx.New(
		logger.Module,
		scheduler.Module,
		router.Module,
		http.Module,
		nats.Module,
	).Run()
}}
