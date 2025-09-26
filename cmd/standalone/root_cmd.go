package main

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin service",
	Run: func(cmd *cobra.Command, args []string) {
		fx.New()
	},
}
