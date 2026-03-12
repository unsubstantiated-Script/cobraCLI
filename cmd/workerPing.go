/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cobraCLI/internal/client"
	"fmt"

	"github.com/spf13/cobra"
)

var workerName string

var workerPingCmd = &cobra.Command{
	Use:   "workerPing",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		grpcClient, err := client.NewGRPCClient(grpcAddr)
		if err != nil {
			return err
		}
		defer grpcClient.Close()

		resp, err := grpcClient.ReportStatus(cmd.Context(), workerName)
		if err != nil {
			return err
		}

		fmt.Printf("grpc response: %s\n", resp)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(workerPingCmd)
	workerPingCmd.Flags().StringVar(&workerName, "name", "cobra-cli", "worker/client name")
}
