/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cobraCLI/internal/client"
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient, err := client.NewHTTPClient(httpAddr)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpClient.HealthCheck(ctx); err != nil {
			return err
		}

		fmt.Println("master reachable")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
