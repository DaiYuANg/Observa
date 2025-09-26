package main

import (
	"github/DaiYuANg/Observa/internal/admin"
	"github/DaiYuANg/Observa/internal/gateway"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin service",
	Run: func(cmd *cobra.Command, args []string) {
		newAdmin, err := admin.NewAdmin()
		if err != nil {
			return
		}
		newGateway, err := gateway.NewGateway()
		if err != nil {
			return
		}
		go newGateway.Run()
		go newAdmin.Run()
		select {}
	},
}
