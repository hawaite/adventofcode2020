package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ErrCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func GetLines(path string) []string {
	f, err := os.Open(path)
	ErrCheck(err)

	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, strings.TrimSuffix(scanner.Text(), "\n"))
	}

	return lines
}

func GetLinesForDay(day int, useExample bool) []string {
	filename := "input.txt"
	if useExample {
		filename = "example.txt"
	}

	return GetLines(fmt.Sprintf("./day%d/input/%s", day, filename))
}
