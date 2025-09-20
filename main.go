// Davids Smits 241RDB187
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Davids Smits 241RDB187")
	scanner := bufio.NewScanner(os.Stdin)
	var in string
	for {

		fmt.Println("Enter: 1 for leastPrimeFactors, 2 to examine two lists, 3 for exit ")
		scanner.Scan()
		in = scanner.Text()

		switch in {
		case "3":
			os.Exit(3)
		case "1":
			getLeastPrimeFactors()
		case "2":
			getCheckEqual()

		}

	}

}

func getCheckEqual() {
	var number int
	var firstList []int
	var secondList []int

	//Get input for first list
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	number, _ = strconv.Atoi(scanner.Text())
	for number != 0 {
		// Switch statement to check if input is a number and if its valid
		switch reflect.TypeOf(number).Kind() {
		case reflect.Int:
			if number > 0 {
				firstList = append(firstList, number)
			} else {
				fmt.Printf("Input error: %d is negative\n", number)
			}

		default:
			fmt.Println("Input error: Unknown error")
		}
		scanner.Scan()
		number, _ = strconv.Atoi(scanner.Text())
	}
	//Get input for second list
	scanner.Scan()
	number, _ = strconv.Atoi(scanner.Text())
	for number != 0 {
		// Switch statement to check if input is a number and if its valid
		switch reflect.TypeOf(number).Kind() {
		case reflect.Int:
			if number > 0 {
				secondList = append(secondList, number)
			} else {
				fmt.Printf("Input error: %d is negative\n", number)
			}
		default:
			fmt.Println("Input error: Unknown error")
		}
		scanner.Scan()
		number, _ = strconv.Atoi(scanner.Text())
	}
	//Output called list checker
	fmt.Println(checkNonEqual(firstList, secondList))

}

func getLeastPrimeFactors() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inpt := scanner.Text()

	newInpt := strings.Split(inpt, " ")
	//Check for invalid input
	for i := 0; i < len(newInpt); i++ {
		val, err := strconv.Atoi(newInpt[i])
		if val <= 2 {
			fmt.Printf("Input error: %d is negative", val)
			fmt.Println()
			continue
		}
		if err != nil {
			fmt.Printf("Input error: %d ", err)
			fmt.Println()
			continue
		}
		//Call data getter
		divList := leastPrimeFactors(nil, val)
		// Data output
		n := strconv.Itoa(val)
		fmt.Printf("%s: ", n)
		for id := range divList {
			n2 := strconv.Itoa(divList[id])
			fmt.Printf("%s ", n2)
		}
		fmt.Println()
	}

}

// 12
// 6 2
// 3 2 2
func leastPrimeFactors(numbers []int, number int) []int {
	// Try to divide by the index till get a non-floating point number
	// If divider is acquired, update the list with the index and call another checker with the new number to divide
	for i := 2; i <= number; i++ {
		if number%i == 0 {
			numbers = append(numbers, i)
			numbers = leastPrimeFactors(numbers, number/i)
			break
		}
	}
	return numbers
}

func checkNonEqual(firstList []int, secondList []int) []int {
	var nonEqualList []int
	isEqual := false
	for i := 0; i < len(firstList); i++ {
		for j := 0; j < len(secondList); j++ {
			if firstList[i] == secondList[j] {
				isEqual = true
				break
			}
		}
		if !isEqual {

			nonEqualList = append(nonEqualList, firstList[i])
		}
		isEqual = false
	}
	return nonEqualList
}
