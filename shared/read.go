package shared

import (
	"fmt"
	"os"
	"strings"
)

func GetStringForDay(day int, use_example bool) string {
	var the_path string
	file_name := fmt.Sprintf("day%d.txt", day)

	if use_example {
		the_path = "examples/" + file_name
	} else {
		the_path = "inputs/" + file_name
	}

	data, err := os.ReadFile(the_path)
	if err != nil {
		panic(err)
	}

	str_data := string(data)
	str_data = strings.ReplaceAll(str_data, "\r", "")

	return str_data
}
