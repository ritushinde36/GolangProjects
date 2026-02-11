/*
Copyright © 2025 RITU SHINDE
*/
package update

import (
	"fmt"
	"todo/todo_items"

	"github.com/spf13/cobra"
)

var (
	name     string
	priority string
	status   string
	new_name string
)

// updateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the status or priority of an item in your todo list",
	Long:  `Update the status or priority of an item in your todo list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if status != "" {
			err := todo_items.Update_status(name, status)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("status updated successfully")
			}
		}
		if priority != "" {
			err := todo_items.Update_priority(name, priority)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Priority updated successfully")
			}
		}
		if new_name != "" {
			err := todo_items.Update_name(name, new_name)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Name updated successfully")
			}
		}

	},
}

func init() {
	// rootCmd.AddCommand(updateCmd)

	UpdateCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the item you want to update.")
	UpdateCmd.MarkFlagRequired("name")

	UpdateCmd.Flags().StringVarP(&status, "status", "s", "", "Update the status of your task.The value should be either Complete or Incomplete.")
	UpdateCmd.Flags().StringVarP(&priority, "priority", "p", "", "Update the priority of your task")
	UpdateCmd.Flags().StringVarP(&new_name, "new_name", "e", "", "Update the name of your task")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
