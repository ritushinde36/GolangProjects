package todo_items

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	filename string = "todo_item_list.json"
	// todo_item_list []Todo_item
	flag_status   int = 0
	flag_priority int = 0
	flag_delete   int = 0
	flag_name     int = 0
)

type Todo_item struct {
	Item_name     string `json:"item_name"`
	Item_status   string `json:"item_status"`
	Item_priority string `json:"item_priority"`
}

func New(name, priority string) (Todo_item, error) {
	caser := cases.Title(language.English)
	priority = caser.String(priority)

	if name == "" {
		return Todo_item{}, errors.New("the name of the item cannot be empty")
	}

	if priority == "Today" || priority == "Tomorrow" || priority == "This Week" || priority == "This Month" {
		item := Todo_item{
			Item_name:     name,
			Item_status:   "Incomplete",
			Item_priority: priority,
		}
		return item, nil
	} else {
		return Todo_item{}, errors.New("invalid value for priority. Acceptable values are - Today, Tomorrow, This week, This Month ")
	}

}

func Process_item(item Todo_item) error {
	//reading the data
	var temp_todo_item_list []Todo_item
	value, _ := os.ReadFile(filename)
	err := json.Unmarshal(value, &temp_todo_item_list)
	if err != nil {
		return errors.New("unable to read file")
	}

	//append the new data in the array
	temp_todo_item_list = append(temp_todo_item_list, item)

	//write the array to the file
	json_data, _ := json.Marshal(temp_todo_item_list)
	err = os.WriteFile(filename, json_data, 0644)
	if err != nil {
		return errors.New("unable to write to file")
	}
	return nil

}

func Check_duplicate_items(item Todo_item) (bool, error) {
	//reading the data
	var temp_todo_item_list []Todo_item
	value, _ := os.ReadFile(filename)
	err := json.Unmarshal(value, &temp_todo_item_list)
	if err != nil {
		return true, errors.New("unable to read file")
	}

	for _, data_value := range temp_todo_item_list {
		if data_value.Item_name == item.Item_name {
			return true, errors.New("entry already exists")
		}
	}

	return false, nil

}

func List_incomplete_items() error {
	//reading the data
	var temp_todo_item_list []Todo_item
	var sorted_todo_item_list []Todo_item
	value, _ := os.ReadFile(filename)
	err := json.Unmarshal(value, &temp_todo_item_list)
	if err != nil {
		return errors.New("unable to read file")
	}

	var label = [4]string{"Today", "Tomorrow", "This Week", "This Month"}

	for i := 0; i < 4; i++ {
		val := label[i]
		for _, data := range temp_todo_item_list {
			if data.Item_priority == val {
				sorted_todo_item_list = append(sorted_todo_item_list, data)
			}

		}

	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Item Name", "Item Status", "Item Priority"})

	for _, data_value := range sorted_todo_item_list {
		if data_value.Item_status == "Incomplete" {
			if data_value.Item_priority == "Today" {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgMagentaColor}, {tablewriter.FgRedColor}, {tablewriter.FgMagentaColor}})

			} else if data_value.Item_priority == "Tomorrow" {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgBlueColor}, {tablewriter.FgRedColor}, {tablewriter.FgBlueColor}})

			} else if data_value.Item_priority == "This Week" {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgCyanColor}, {tablewriter.FgRedColor}, {tablewriter.FgCyanColor}})

			} else {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgYellowColor}, {tablewriter.FgRedColor}, {tablewriter.FgYellowColor}})

			}
		}
	}
	table.Render()

	return nil
}

func List_all_items() error {
	//reading the data
	var temp_todo_item_list []Todo_item
	var sorted_todo_item_list []Todo_item
	value, _ := os.ReadFile(filename)
	err := json.Unmarshal(value, &temp_todo_item_list)
	if err != nil {
		return errors.New("unable to read file")
	}

	var label = [4]string{"Today", "Tomorrow", "This Week", "This Month"}

	for i := 0; i < 4; i++ {
		val := label[i]
		for _, data := range temp_todo_item_list {
			if data.Item_priority == val {
				sorted_todo_item_list = append(sorted_todo_item_list, data)
			}

		}

	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Item Name", "Item Status", "Item Priority"})

	for _, data_value := range sorted_todo_item_list {
		if data_value.Item_status == "Incomplete" {
			if data_value.Item_priority == "Today" {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgMagentaColor}, {tablewriter.FgRedColor}, {tablewriter.FgMagentaColor}})

			} else if data_value.Item_priority == "Tomorrow" {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgBlueColor}, {tablewriter.FgRedColor}, {tablewriter.FgBlueColor}})

			} else if data_value.Item_priority == "This Week" {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgCyanColor}, {tablewriter.FgRedColor}, {tablewriter.FgCyanColor}})

			} else {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgYellowColor}, {tablewriter.FgRedColor}, {tablewriter.FgYellowColor}})

			}
		}
	}

	table.Append([]string{"", "", ""})

	for _, data_value := range sorted_todo_item_list {
		if data_value.Item_status == "Complete" {
			if data_value.Item_priority == "Today" {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgMagentaColor}, {tablewriter.FgGreenColor}, {tablewriter.FgMagentaColor}})

			} else if data_value.Item_priority == "Tomorrow" {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgBlueColor}, {tablewriter.FgGreenColor}, {tablewriter.FgBlueColor}})

			} else if data_value.Item_priority == "This Week" {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgCyanColor}, {tablewriter.FgGreenColor}, {tablewriter.FgCyanColor}})

			} else {
				table.Rich([]string{data_value.Item_name, data_value.Item_status, data_value.Item_priority}, []tablewriter.Colors{{tablewriter.FgYellowColor}, {tablewriter.FgGreenColor}, {tablewriter.FgYellowColor}})

			}
		}
	}

	table.Render()
	return nil

}

func Update_status(name, status string) error {

	if name == "" {
		return errors.New("the name of the item cannot be empty")
	}

	//check if the value of status is either complete or incomplete
	status = strings.ToLower(status)
	if status == "complete" || status == "incomplete" {
		caser := cases.Title(language.English)
		status = caser.String(status)

		//read data and store it into an array
		var temp_todo_item_list []Todo_item
		value, _ := os.ReadFile(filename)
		err := json.Unmarshal(value, &temp_todo_item_list)
		if err != nil {
			return errors.New("unable to read file")
		}

		for index, data_value := range temp_todo_item_list {
			if data_value.Item_name == name {
				flag_status = 1
				if status == data_value.Item_status {
					text := fmt.Sprintf("the status is already in %v state", status)
					return errors.New(text)
				} else {
					temp_todo_item_list[index].Item_status = status
					break
				}
			}
		}
		if flag_status == 0 {
			return errors.New("item name not found")
		}

		json_data, _ := json.Marshal(temp_todo_item_list)
		err = os.WriteFile(filename, json_data, 0644)
		if err != nil {
			return errors.New("unable to write to file")
		}
		return nil

	} else {
		return errors.New("the value of status should be either complete or incomplete")
	}

}

func Update_priority(name, priority string) error {
	if name == "" {
		return errors.New("the name of the item cannot be empty")
	}

	priority = strings.ToLower(priority)
	if priority == "today" || priority == "tomorrow" || priority == "this week" || priority == "this month" {
		caser := cases.Title(language.English)
		priority = caser.String(priority)

		//read data and store it into an array
		var temp_todo_item_list []Todo_item
		value, _ := os.ReadFile(filename)
		err := json.Unmarshal(value, &temp_todo_item_list)
		if err != nil {
			return errors.New("unable to read file")
		}

		for index, data_value := range temp_todo_item_list {
			if data_value.Item_name == name {
				flag_priority = 1
				if priority == data_value.Item_priority {
					text := fmt.Sprintf("the priority is already set to %v", priority)
					return errors.New(text)
				} else {
					temp_todo_item_list[index].Item_priority = priority
					break
				}
			}
		}
		if flag_priority == 0 {
			return errors.New("item name not found")
		}

		json_data, _ := json.Marshal(temp_todo_item_list)
		err = os.WriteFile(filename, json_data, 0644)
		if err != nil {
			return errors.New("unable to write to file")
		}
		return nil

	} else {
		return errors.New("invalid value for priority. Acceptable values are - Today, Tomorrow, This week, This Month ")
	}

}

func Update_name(name, new_name string) error {
	if name == "" || new_name == "" {
		return errors.New("the name of the item cannot be empty")
	}

	//read data and store it into an array
	var temp_todo_item_list []Todo_item
	value, _ := os.ReadFile(filename)
	err := json.Unmarshal(value, &temp_todo_item_list)
	if err != nil {
		return errors.New("unable to read file")
	}

	for index, data_value := range temp_todo_item_list {
		if data_value.Item_name == name {
			flag_name = 1
			temp_todo_item_list[index].Item_name = new_name
			break
		}
	}
	if flag_name == 0 {
		return errors.New("item name not found")
	}
	json_data, _ := json.Marshal(temp_todo_item_list)
	err = os.WriteFile(filename, json_data, 0644)
	if err != nil {
		return errors.New("unable to write to file")
	}
	return nil

}

func Delete_item(name string) error {
	if name == "" {
		return errors.New("the name of the item cannot be empty")
	}
	//read data and store it into an array
	var temp_todo_item_list []Todo_item
	value, _ := os.ReadFile(filename)
	err := json.Unmarshal(value, &temp_todo_item_list)
	if err != nil {
		return errors.New("unable to read file")
	}

	for index, data_value := range temp_todo_item_list {
		if name == data_value.Item_name {
			flag_delete = 1
			temp_todo_item_list = append(temp_todo_item_list[:index], temp_todo_item_list[index+1:]...)
			break
		}

	}
	if flag_delete == 0 {
		return errors.New("item name not found")
	}
	json_data, _ := json.Marshal(temp_todo_item_list)
	err = os.WriteFile(filename, json_data, 0644)
	if err != nil {
		return errors.New("unable to write to file")
	}
	return nil

}
