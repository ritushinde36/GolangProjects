/*
Copyright © 2025 RITU SHINDE
*/
package list

import (
	"fmt"
	"todo/todo_items"

	"github.com/spf13/cobra"
)

var (
	all bool
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all you tasks",
	Long:  `List all the Complete and Incomplete tasks using --all.`,
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			err := todo_items.List_all_items()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := todo_items.List_incomplete_items()
			if err != nil {
				fmt.Println(err)
			}
		}

	},
}

func init() {
	// rootCmd.AddCommand(listCmd)

	// ListCmd.Flags().StringVarP(&all, "all", "a", "", "List the List all the Complete and Incomplete tasks.")
	ListCmd.Flags().BoolVarP(&all, "all", "a", false, "List the List all the Complete and Incomplete tasks.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
