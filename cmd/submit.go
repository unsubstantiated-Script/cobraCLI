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

var command string

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a task to the distributed system",
	RunE: func(cmd *cobra.Command, args []string) error {
		httpClient, err := client.NewHTTPClient(httpAddr)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		resp, err := httpClient.SubmitTask(ctx, command)
		if err != nil {
			return err
		}

		fmt.Printf("queued: %s\n", resp.Command)
		fmt.Printf("status: %s\n", resp.Status)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(submitCmd)
	submitCmd.Flags().StringVarP(&command, "command", "c", "", "command to enqueue")
	_ = submitCmd.MarkFlagRequired("command")
}
