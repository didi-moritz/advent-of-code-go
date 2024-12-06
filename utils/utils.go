package utils

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetRunConfig(defaultPart int, defaultRealData bool) (int, bool) {
	partPtr := flag.Int("p", defaultPart, "Part of the day to run")
	useRealDataPtr := flag.Bool("r", false, "Use real data")
	useTestDataPtr := flag.Bool("t", false, "Use test data")

	flag.Parse()

	var useRealData bool
	if *useRealDataPtr == *useTestDataPtr {
		useRealData = defaultRealData
	} else {
		useRealData = *useRealDataPtr
	}

	fmt.Println("Running Part:", *partPtr, " Real Data:", useRealData)

	return *partPtr, useRealData
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

func ReadFileAsByteArray(filename string) [][]byte {
	file, err := os.Open(filename)
	checkError(err)

	defer func(file *os.File) {
		err := file.Close()
		checkError(err)
	}(file)

	reader := bufio.NewReader(file)

	var data [][]byte

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)
		data = append(data, lineCopy)
	}

	return data
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
		if err != nil {
			break
		}
		data = append(data, string(line))
	}

	return data
}

func ParseInt(s string) int64 {
	value, _ := strconv.ParseInt(s, 10, 64)
	return value
}

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Println("It took", elapsed)
}
