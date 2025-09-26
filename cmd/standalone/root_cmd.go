package main

import (
	"github/DaiYuANg/Observa/internal/gateway"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin service",
	RunE: func(cmd *cobra.Command, args []string) error {
		gateway, err := gateway.NewGateway()
		if err != nil {
			return err
		}
		nap := NewAppContainer(gateway.GetApp())
		return nap.Run()
	},
}
