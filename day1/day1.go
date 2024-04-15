package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./input/input.txt")
	check(err)

	scanner := bufio.NewScanner(f)
	num_arr := []int{}
	num_arr_size := 0
	for scanner.Scan() {
		token := scanner.Text()
		num, err := strconv.Atoi(token)
		check(err)
		num_arr = append(num_arr, num)
		num_arr_size++
	}

	// Part 1
	for i := 0; i < num_arr_size; i++ {
		for j := (i + 1); j < num_arr_size; j++ {
			if (num_arr[i] + num_arr[j]) == 2020 {
				fmt.Println("Found values: ", num_arr[i], "and", num_arr[j])
				fmt.Println("Result: ", num_arr[i]*num_arr[j])
			}
		}
	}

	// Part 2
	for i := 0; i < num_arr_size; i++ {
		for j := (i + 1); j < num_arr_size; j++ {
			for k := (j + 1); k < num_arr_size; k++ {
				if (num_arr[i] + num_arr[j] + num_arr[k]) == 2020 {
					fmt.Println("Found values: ", num_arr[i], "and", num_arr[j], "and", num_arr[k])
					fmt.Println("Result: ", num_arr[i]*num_arr[j]*num_arr[k])
				}
			}
		}
	}
}
