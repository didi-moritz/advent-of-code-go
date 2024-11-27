package utils

import (
	"bufio"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetFileName(year int, day int, real bool) string {
	var fileName string

	if real {
		fileName = "data"
	} else {
		fileName = "test.data"
	}

	var dayString string
	if day < 10 {
		dayString = "0" + strconv.Itoa(day)
	} else {
		dayString = strconv.Itoa(day)
	}

	return strconv.Itoa(year) + "/day-" + dayString + "/" + fileName
}

func ReadFileAsStringArray(filename string) []string {
	file, err := os.Open(filename)
	checkError(err)

	defer func(file *os.File) {
		err := file.Close()
		checkError(err)
	}(file)

	reader := bufio.NewReader(file)

	var data []string

	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			data = append(data, string(line))
		}
		if err != nil {
			break
		}
	}

	return data
}

func ParseInt(s string) int64 {
	value, _ := strconv.ParseInt(s, 10, 64)
	return value
}
