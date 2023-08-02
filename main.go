package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var greeting = "Welcome to Go Choose Your Own Adventure. Press any key to begin the story:"
var storyIntro = "Once upon a time there was a girl. She wanted to play a game."
var choiceOne = "checkers"
var choiceTwo = "chess"

func main() {
	fmt.Println(greeting)
	_, _, _ = keyboard.GetSingleKey()

	fmt.Println(storyIntro)
	fmt.Printf("She could play %s or she could play %s.\n", choiceOne, choiceTwo)
	fmt.Println("Which game should she play?\n Select option 1 or 2:")
	fmt.Printf("1. %s\n2. %s\n", choiceOne, choiceTwo)

	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	playStory(string(char))
}

func playStory(choice string) {
	switch choice {
	case "1":
		fmt.Printf("The girl plays %s. She has fun.\n", choiceOne)
	case "2":
		fmt.Printf("The girl plays %s. She has fun.\n", choiceTwo)
	default:
		fmt.Println("That wasn't one of her options.")
	}
}
