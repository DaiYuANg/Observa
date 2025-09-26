package main

import (
	"github/DaiYuANg/Observa/modules/http"
	"github/DaiYuANg/Observa/modules/logger"
	"github/DaiYuANg/Observa/modules/nats"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var rootCmd = &cobra.Command{Run: func(cmd *cobra.Command, args []string) {
	fx.New(
		logger.Module,
		http.Module,
		nats.Module,
	).Run()
}}
