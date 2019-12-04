package Utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Reads a file and parses it to int lines
func ReadFileByLinesInt(filename string) []int {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(dir + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var l []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			log.Fatal(err)
		} else {
			l = append(l, int(i))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return l
}

// Reads a file and parses it by separation at the given separator
func ReadFileBySeparatorInt(filename string, separator string) []int {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dat, err := ioutil.ReadFile(dir + filename)
	if err != nil {
		log.Fatal(err)
	}

	text := string(dat)
	textArr := strings.Split(text, separator)

	var intArr []int

	for i := 0; i < len(textArr); i++ {
		x, err := strconv.ParseInt(textArr[i], 10, 32)
		if err != nil {
			log.Fatal(err)
		} else {
			intArr = append(intArr, int(x))
		}
	}

	return intArr
}
