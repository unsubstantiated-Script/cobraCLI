/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cobraCLI/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
