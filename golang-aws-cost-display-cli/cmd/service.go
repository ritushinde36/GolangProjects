/*
Copyright © 2026 RITU SHINDE
*/
package awscost

import (
	"fmt"

	"github.com/ritushinde36/GolangProjects/golang-aws-cost-display-cli/subcommand_operations"
	"github.com/spf13/cobra"
)

var (
	name    string
	month   string
	between string
)

// serviceCmd represents the service command
var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Display cost for a specific AWS service",
	Long: `Display cost breakdown for a specific AWS service.

Examples:
  awscost service --name ec2
  awscost service --name ec2 --month 2024-01
  awscost service --name ec2 --between 2024-01-01 2024-01-31`,
	Run: func(cmd *cobra.Command, args []string) {
		err := subcommand_operations.CheckServiceCmdFlags(name, month, between)
		if err != nil {
			fmt.Printf("Error : %v", err)
		}
		fmt.Println("service called")
	},
}

func init() {
	ServiceCmd.Flags().StringVarP(&name, "name", "n", "", "AWS service to fetch costs for (e.g. ec2, s3, rds)")
	ServiceCmd.MarkFlagRequired("name")

	ServiceCmd.Flags().StringVarP(&month, "month", "m", "", "Month to fetch costs for (e.g. 2024-05)")
	ServiceCmd.Flags().StringVarP(&between, "between", "b", "", "Fetch costs between two dates (e.g. 2024-05-01,2024-05-31)")

}
