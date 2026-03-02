package quiz

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func Pocess_questions(filename string, time_limit int) error {
	score := 0
	var correct_questions []int
	var incorrect_questions []int
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	timer := time.NewTimer(time.Duration(time_limit) * time.Second)
	// <-timer.C

	for index, value := range records {
		var data string
		fmt.Printf("Problem #%v -> %v\n", index+1, value[0])
		fmt.Printf("Answer : ")
		datachan := make(chan string)
		go func() {
			fmt.Scan(&data)
			datachan <- data
		}()

		select {
		case <-timer.C:
			// cyan.Printf("\n\nTime is up!\n")
			color.Cyan("\n\nTime is up!\n")
			display_score(score, correct_questions, incorrect_questions)
			return nil
		case data := <-datachan:
			answer, err := strconv.Atoi(data)
			if err != nil {
				return errors.New("answer should be an integer,unable to convert to int")
			}
			correct_answer, err := strconv.Atoi(value[1])
			if err != nil {
				return errors.New("unable to convert to int")
			}

			if answer == correct_answer {
				score++
				correct_questions = append(correct_questions, index+1)
			} else {
				incorrect_questions = append(incorrect_questions, index+1)

			}
			fmt.Println()

		}

	}

	display_score(score, correct_questions, incorrect_questions)

	// fmt.Print(records)
	return nil

}

func display_score(score int, correct_questions []int, incorrect_questions []int) {

	color.Magenta("Final Score : %v", score)

	color.Set(color.FgGreen)
	fmt.Print("Correct Questions : ")
	for index, question := range correct_questions {
		if len(correct_questions)-1 == index {
			fmt.Printf("%v ", question)
		} else {
			fmt.Printf("%v , ", question)
		}

	}
	color.Unset()
	color.Set(color.FgRed)
	fmt.Print("\nIncorrect Questions : ")
	for index, question := range incorrect_questions {
		if len(incorrect_questions)-1 == index {
			fmt.Printf("%v ", question)

		} else {
			fmt.Printf("%v , ", question)
		}
	}
	color.Unset()

}
