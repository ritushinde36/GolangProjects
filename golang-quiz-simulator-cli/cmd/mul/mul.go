/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package mul

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

// mulCmd represents the mul command
var MulCmd = &cobra.Command{
	Use:   "mul",
	Short: "command to generate a quiz with multiplication problems",
	Long:  `command to generate a quiz with multiplication problems`,
	Run: func(cmd *cobra.Command, args []string) {

		time_limit, err := strconv.Atoi(limit)
		if err != nil {
			log.Fatal(err)
		}

		if file == "" {
			err := quiz.Pocess_questions("./problem_files/multiplication_problems.csv", time_limit)
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
	MulCmd.Flags().StringVarP(&file, "file", "f", "", "add the name of a csv file that contains a set of multiplication questions.The format has to be 'question,answer' ex = 5*2,10")
	MulCmd.Flags().StringVarP(&limit, "limit", "l", "30", "the time limit for the quiz in seconds")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mulCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mulCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
