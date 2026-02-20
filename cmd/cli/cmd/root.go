package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const version = "1.0.0"

var rootCmd = &cobra.Command{
	Use:     "go-backend",
	Short:   "Go Backend - Scaffold production-ready Go backend projects",
	Long:    `Go Backend is a CLI tool that scaffolds a production-ready Go backend project with clean architecture, authentication, database setup, and more.`,
	Version: version,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
