/*
Copyright © 2026 RITU SHINDE
*/
package cmd

import (
	"context"
	"os"

	"github.com/ritushinde36/GolangProjects/golang-aws-cost-display-cli/aws_operations"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "awscost",
	Short: "AWS cost explorer CLI",
	Long:  `awscost is a CLI tool to explore and analyze your AWS cloud costs`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		err := aws_operations.CheckAwsCredentials(ctx)
		if err != nil {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommand() {
	rootCmd.AddCommand(ServiceCmd)
	rootCmd.AddCommand(ServicesCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommand()
}
