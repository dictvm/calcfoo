package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {

// also based on: http://piratepad.net/XJnL7LdAUx

    var inputNum, guessedNum, tries, opponentNum int
    var opponentSelected bool = false

    for opponentSelected == false {
        fmt.Println("Choose your opponent: \n 0 = human \n 1 = AI \n")
        fmt.Scanln(&opponentNum)
            if opponentNum == 0 {
                opponentSelected = true
                fmt.Println("Enter a number between 1 and 10 that your opponent shall guess.")
                inputNum = getInput()
            } else if opponentNum == 1 {
                opponentSelected = true
                inputNum = random(1, 10)
                fmt.Println("The AI has chosen a random number for you to guess.")
            }
        }

    fmt.Println("Now try to guess your opponents' input.")

    for tries < 3  {
        fmt.Println("Enter a number between 1 and 10.")
        guessedNum = getInput()
        tries++

        if guessEqualsInput(inputNum, guessedNum) == true {
            fmt.Println("Good job, your guess was correct.")
            fmt.Println("The entered number was", inputNum, ".")
            break // exits the current for-loop
        }

        if guessedNum > inputNum {
            fmt.Println("Your number was bigger than the input of your opponent.")
        } else {
            fmt.Println("Your number was smaller than the input of your opponent.")
        }
    }

    if guessEqualsInput(inputNum, guessedNum) != true {
        fmt.Println("You lost the game. Your opponent entered", inputNum, ".")
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

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}
