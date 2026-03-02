/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package div

import (
	"log"
	"quiz/quiz"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	file  string
	limit string
)

// divCmd represents the div command
var DivCmd = &cobra.Command{
	Use:   "div",
	Short: "command to generate a quiz with division problems",
	Long:  `command to generate a quiz with division problems`,
	Run: func(cmd *cobra.Command, args []string) {
		time_limit, err := strconv.Atoi(limit)
		if err != nil {
			log.Fatal(err)
		}
		if file == "" {
			err := quiz.Pocess_questions("./problem_files/division_problems.csv", time_limit)

			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := quiz.Pocess_questions(file, time_limit)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	DivCmd.Flags().StringVarP(&file, "file", "f", "", "add the name of a csv file that contains a set of division questions. The format has to be 'question,answer' ex = 10/2,3")
	DivCmd.Flags().StringVarP(&limit, "limit", "l", "30", "the time limit for the quiz in seconds")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
