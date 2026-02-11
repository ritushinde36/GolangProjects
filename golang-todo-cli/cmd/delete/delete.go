/*
Copyright © 2025 RITU SHINDE
*/
package delete

import (
	"fmt"
	"strings"
	"todo/todo_items"

	"github.com/spf13/cobra"
)

var (
	name string
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an item from your todo list",
	Long:  `delete an item from your todo list`,
	Run: func(cmd *cobra.Command, args []string) {
		var answer string
		fmt.Printf("Are you sure you want to delete item - %v ?\n", name)
		fmt.Printf("Please enter Y or N : ")
		fmt.Scan(&answer)
		answer = strings.ToLower(answer)
		if answer == "y" {
			err := todo_items.Delete_item(name)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Item is deleted")
			}
		} else if answer == "n" {
			fmt.Printf("Item will not be deleted")
		} else {
			fmt.Printf("Please enter either Y or N")
		}

	},
}

func init() {
	// rootCmd.AddCommand(deleteCmd)
	DeleteCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the item you want to delete")
	DeleteCmd.MarkFlagRequired("name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
