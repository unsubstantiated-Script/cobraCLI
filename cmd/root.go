/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	httpAddr string
	grpcAddr string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobraCLI",
	Short: "CLI for the Go distributed system",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&httpAddr,
		"http-addr",
		":8080",
		"master HTTP server address",
	)
	rootCmd.PersistentFlags().StringVar(
		&grpcAddr,
		"grpc-addr",
		":50051",
		"master gRPC server address",
	)
}
