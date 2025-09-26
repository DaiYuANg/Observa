package main

import (
	"github/DaiYuANg/Observa/internal/admin"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "worker",
	Short: "worker",
	RunE: func(cmd *cobra.Command, args []string) error {
		newAdmin, err := admin.NewAdmin()
		if err != nil {
			return err
		}
		newAdmin.Run()
		return nil
	},
}
