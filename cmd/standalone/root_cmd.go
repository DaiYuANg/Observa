package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin service",
	RunE: func(cmd *cobra.Command, args []string) error {
		standalone := NewAppContainer()
		return standalone.Run()
	},
}
