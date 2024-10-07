/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-stress-test",
	Short: "This is a CLI tool to stress test a Go application",
	Long:  `This is a CLI tool to stress test a Go application. It will send a number of requests to a given URL and return the response time and status code of each request.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
