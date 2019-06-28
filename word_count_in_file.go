package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var j int

func main() {
	j = 0

	file, err := os.Open("PATH_TO_THE_FILE/wordcount-1.txt")
	_ = err
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		log_file := scanner.Text()
		st1 := strings.Split(log_file, " ")
		j = j + len(st1)

	}

	fmt.Println(j)

}
