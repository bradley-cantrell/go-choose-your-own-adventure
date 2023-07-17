package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(
		"Welcome to Go Choose Your Own Adventure. Press Enter to begin the story:",
	)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	
	fmt.Println("Once upon a time there was a girl. She wanted to play a game. She could play Tic-Tac-Toe or Checkers. Which game should she play?.")
	fmt.Println("Type 1 or 2 and then press Enter:\n 1. Tic-Tac-Toe\n 2. Checkers") 

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	switch choice {
	case "1\n":
		fmt.Println("The girl plays Tic-Tac-Toe. She has fun.")
	case "2\n":
		fmt.Println("The girl plays Checkers. She has fun.")
	default:
		fmt.Println("That wasn't one of her options.")
	}

}
