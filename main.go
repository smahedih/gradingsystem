package main

import (
	"fmt"
)

func main() {

	var number int

	fmt.Printf("Enter your number: \n")

	for true {
		_, err := fmt.Scanf("%d", &number)
		if err == nil {
			break
		}
		fmt.Println("Not a valid number - try again")
		var dump string
		fmt.Scanln(&dump)
	}

	if number >= 80 && number <= 100 {
		fmt.Printf("GPA: 5 and Grade: A+")

	} else if number >= 70 && number <= 79 {
		fmt.Printf("GPA: 4 and Grade: A")

	} else if number >= 60 && number <= 69 {
		fmt.Printf("GPA: 3.5 and Grade: A-")

	} else if number >= 50 && number <= 59 {
		fmt.Printf("GPA: 3 and Grade: B")

	} else if number >= 40 && number <= 49 {
		fmt.Printf("GPA: 2 and Grade: C")

	} else if number >= 33 && number <= 39 {
		fmt.Printf("GPA: 1 and Grade: D")

	} else if number >= 0 && number < 33 {
		fmt.Printf("Failed")

	} else {
		fmt.Printf("Invalid input. Please enter number 0 to 100.")
	}
}
