package main

import "fmt"

func nprime(number int) {
	isPrime := true
	if number == 0 || number == 1 {
		fmt.Printf(" %d is not a  prime number\n", number)
	} else {
		for i := 2; i <= number/2; i++ {
			if number%i == 0 {
				fmt.Printf(" %d is not a  prime number\n", number)
				isPrime = false
				break
			}
		}
		if isPrime == true {
			fmt.Printf(" %d is a prime number\n", number)
		}
	}
}

func main() {
	var num int
	fmt.Println("Enter your number")
	fmt.Scan(&num)
	nprime(num)
}
