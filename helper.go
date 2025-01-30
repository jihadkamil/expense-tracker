package main

import (
	"fmt"
	"strconv"
	"strings"
)

// // get latest number used as id
// func getLastId(listTask []Task) int {
// 	if len(listTask) == 0 {
// 		return 0 // or handle it differently if no tasks exist
// 	}
// 	return listTask[len(listTask)-1].Id
// }

// validating input from user has 2 args & the last argument is suit with the task name rule
func isValidAddInput(inputs []string) bool {
	if len(inputs) < 3 {
		return false
	}
	// Check if input starts with 'add "'
	// parameter description is required
	if inputs[1] != "--description" {
		return false
	}
	// description
	if strings.HasPrefix(inputs[2], `"`) && strings.HasSuffix(inputs[2], `"`) {
		// Remove the "add " part and check if the value inside quotes is non-empty
		content := strings.TrimSpace(inputs[2][1:])
		fmt.Println("content", content)
		if len(content) >= 2 && strings.HasPrefix(content, `"`) && strings.HasSuffix(content, `"`) {
			return true
		}
	}
	return false
}

// validating input from user has 2 args
// this func can be used by other command with same input validation
func isValidDeleteInput(inputs []string) bool {
	if len(inputs) > 1 {
		// check id
		if _, err := strconv.Atoi(inputs[1]); err == nil {
			return true
		}
	}
	return false
}
