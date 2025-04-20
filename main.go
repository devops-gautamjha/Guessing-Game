package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	bold   = "\033[1m"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Ensure true randomness

	var guessedNum int
	attempts := 4

	fmt.Println(bold + cyan + "===========================================" + reset)
	fmt.Println(bold + "         🎮  Welcome to Guess It! 🎮" + reset)
	fmt.Println(bold + cyan + "===========================================" + reset)

	secretNum := rand.Intn(100) + 1 // number between 1 and 100
	// fmt.Println("DEBUG: ", secretNum) // Uncomment for testing

	fmt.Println(cyan + "👉 Guess a number between 1 and 100." + reset)

	for i := attempts; i > 0; {
		fmt.Printf(bold + "> " + reset)
		_, err := fmt.Scanf("%d", &guessedNum)

		if err != nil {
			fmt.Println(red + "❌ Invalid input. Please enter a valid number." + reset)
			continue // Don't reduce attempts on invalid input
		}

		if guessedNum == secretNum {
			fmt.Println(green + "🎉 Hurray! You guessed the correct number!" + reset)
			return
		} else if guessedNum < secretNum {
			fmt.Println(yellow + "📉 Too Low!" + reset)
		} else {
			fmt.Println(yellow + "📈 Too High!" + reset)
		}

		i--
		if i > 0 {
			fmt.Printf(cyan+"🔁 You have %d %s left. Try again!\n"+reset, i, pluralize("attempt", i))
		} else {
			fmt.Println(red + "💀 Ooof... You've used all your attempts." + reset)
			fmt.Printf("🔎 The correct number was: %s%d%s\n", bold, secretNum, reset)
		}
	}

	fmt.Println(cyan + "🙏 Thanks for playing! Better luck next time." + reset)
}

func pluralize(word string, count int) string {
	if count == 1 {
		return word
	}
	return word + "s"
}
