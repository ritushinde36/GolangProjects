/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"quiz/cmd/add"
	"quiz/cmd/div"
	"quiz/cmd/mul"
	"quiz/cmd/sub"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "quiz application helps to study addition/subtraction/multiplication/division",
	Long: `quiz application helps to study addition/subtraction/multiplication/division. 
	You have to answer a set of questions within the time limit(if specified) and 
	your score will be displayed once you are done`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addCommand() {
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(sub.SubCmd)
	rootCmd.AddCommand(mul.MulCmd)
	rootCmd.AddCommand(div.DivCmd)

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.quiz.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCommand()
}
