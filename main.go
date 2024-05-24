package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func myPrint(start, finish int) {
	for i := start; i <= finish; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	time.Sleep(100 * time.Microsecond)
}

func check(a int, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this is a custom error message")
	}

	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Getuid())
	}
		return nil
}

func main() {

	argument := os.Args[1]

	switch argument {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough
	default:
		fmt.Println("Value:", argument)
	}

	value, err := strconv.Atoi(argument)

	if err != nil {
		fmt.Println("Cannot convert to int:", argument)

		return
	}

	switch {
	case value == 0:
		fmt.Println("Zero!")
	case value > 0:
		fmt.Println("Positive integer")
	case value < 0:
		fmt.Println("Negative integer")
	default:
		fmt.Println("This should not happen:", value)
	}

	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")

		i++
	}

	fmt.Println()

	// for loop used as while loop
	i = 0
	for {
		if i == 10 {
			break
		}

		fmt.Print(i*i, " ")
		i++

	}

	fmt.Println()

	aSlice := []int{-1, 2, 1, -1, 2, -2}

	for i, v := range aSlice {
		fmt.Println("index:", i, "value: ", v)
	}

	fmt.Printf("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")

		return
	}

	var min, max float64
	var initialized = 0

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)

		if err != nil {
			continue
		}

		if initialized == 0 {
			min = n
			max = n
			initialized = 1

			continue
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)

	// validating inputs
	var total, nInts, nFloats int

	invalid := make([]string, 0)

	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)

		if err == nil {
			total++
			nInts++

			continue
		}

		_, err = strconv.ParseFloat(k, 64)

		if err == nil {
			total++
			nFloats++

			continue
		}

		invalid = append(invalid, k)
	}

	fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)

	if len(invalid) > total {
		fmt.Println("Too much invalid input:", len(invalid))

		for _, s := range invalid {
			fmt.Println(s)
		}
	}

	for i := 0; i < 4; i++ {
		go myPrint(i, 5)
	}

	time.Sleep(time.Second)
}