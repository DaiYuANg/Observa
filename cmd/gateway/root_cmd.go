package main

import (
	"github/DaiYuANg/Observa/internal/gateway"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{RunE: func(cmd *cobra.Command, args []string) error {
	newGateway, err := gateway.NewGateway()
	if err != nil {
		return err
	}
	newGateway.Run()
	return nil
}}
