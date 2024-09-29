package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing game")
	fmt.Println("A random number will be generated. Try to guess it.")
	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("What is your guess for the number?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}
		switch {
		case chuteInt < x:
			fmt.Println("Too low")

		case chuteInt > x:
			fmt.Println("Too high")

		case chuteInt == x:
			fmt.Println("You guessed it!")
			fmt.Println("You guessed in", i+1, "attempts")
			fmt.Printf("Guesses: %v\n", chutes[:i])
			return

		}
		fmt.Printf("You have %d guesses left\n", 9-i)
		chutes[i] = chuteInt

	}
	fmt.Printf("You Lose, Guesses: %v\n", chutes)

}
