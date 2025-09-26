package main

import (
	"github/DaiYuANg/Observa/modules/config"
	"github/DaiYuANg/Observa/modules/db"
	"github/DaiYuANg/Observa/modules/logger"
	"github/DaiYuANg/Observa/modules/nats"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var rootCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Gateway is a service that provides a REST API for the system",
	Run: func(cmd *cobra.Command, args []string) {
		fx.New(
			logger.Module,
			config.Module,
			db.Module,
			nats.ClientModule,
		)
	},
}
