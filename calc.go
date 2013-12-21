package main

import (
    "fmt"
)

func main() {

// also based on: http://piratepad.net/XJnL7LdAUx

    var inputNum int
    var guessedNum int
    var tries int = 0

    fmt.Println("Enter a number between 1 and 10 that your opponent shall guess.")
    inputNum = getInput()

    fmt.Println("Now try to guess your opponents' input.")

    for tries < 3  {
        fmt.Println("Enter a number between 1 and 10.")
        guessedNum = getInput()
        tries++

        if guessEqualsInput(inputNum, guessedNum) == true {
            fmt.Println("Good job, your guess was correct.")
            break // exits the current for-loop
        }

        if guessedNum > inputNum {
            fmt.Println("Your number was bigger than the input of your opponent.")
        } else {
            fmt.Println("Your number was smaller than the input of your opponent.")
        }
    }

    if guessEqualsInput(inputNum, guessedNum) != true {
        fmt.Println("You lost the game. Your opponent entered ", inputNum, ".")
    }

}

func getInput() int {
    var inputNum int
    var isInputCorrect bool = false

    for isInputCorrect == false {
        fmt.Scanln(&inputNum)
        if inputNum <= 10 && inputNum >= 1 {
            isInputCorrect = true
        } else {
            fmt.Println("Nope. Way out of scope. Stay between 1 and 10.")
        }
    }
    return inputNum
}

func guessEqualsInput(inputNum int, guessedNum int) bool {
    if guessedNum == inputNum {
        return true
    } else {
    return false
    }
}
