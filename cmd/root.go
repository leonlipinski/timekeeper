package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// rootCmd is the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "timekeeper",
    Short: "timekeeper is a CLI to track time spent on tasks",
    Long: `A simple command-line app to track how much time you
spend on tasks for different customers, stored in a CSV file.`,
}

// Execute adds all child commands to the root command and sets flags
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
