package main

import (
    "fmt"
)

func main() {

// also based on: http://piratepad.net/XJnL7LdAUx

    var inputNum int
    var guesses []int
    var guessTries int

    fmt.Println("Enter the number your opponent shall guess.")
    fmt.Scanln(&inputNum)

    for guessTries < 2 {
        guessTries++
        fmt.Println("Guess the number your opponent entered.")
        fmt.Scanln(&guesses)
            if guessCorrect(inputNum, guesses) == true {
                fmt.Println("You guessed correctly, good job.")
            } else {
                fmt.Println("Try again.")
            }
    }
}

func guessCorrect (inputNum int, guesses []int) bool {
    for guessContainsInput, inputNum := range guesses {
        if guessContainsInput == inputNum {
            return true
        }
    }
    return false
}
