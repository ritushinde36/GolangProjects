/*
Copyright © 2025 RITU SHINDE
*/
package add

import (
	"fmt"
	"todo/todo_items"

	"github.com/spf13/cobra"
)

var (
	name     string
	priority string
)

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an Item to you todo list",
	Long:  `Add an Item to you todo list. The priority will help you identify which item to work on`,
	Run: func(cmd *cobra.Command, args []string) {
		item, err := todo_items.New(name, priority)
		if err != nil {
			fmt.Println(err)
		} else {
			isduplicate, err := todo_items.Check_duplicate_items(item)
			if isduplicate {
				fmt.Println(err)
			} else {
				err = todo_items.Process_item(item)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("Item - %v added with priority - %v\n", name, priority)
				}

			}

		}
	},
}

func init() {
	// rootCmd.AddCommand(addCmd)

	AddCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the item you want to add")
	AddCmd.MarkFlagRequired("name")

	AddCmd.Flags().StringVarP(&priority, "priority", "p", "Today", "Assign the item a priority. Acceptable values are - Today, Tomorrow, This week, This Month")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
