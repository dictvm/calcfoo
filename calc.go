package main

import (
    "fmt"
    "math/rand"
    "os"
    "io"
    "time"
    "strconv"
    "encoding/csv"
    )

func main() {

// based on part #2 of: https://pads.ccc.de/g7cegAmhkN
    var user string
    var inputNum, guessedNum, tries, opponentNum, score int
    var opponentSelected bool = false
    var userNameSet bool = false

// loop through username-input
    for userNameSet == false {
        fmt.Println("Enter your name.")
        fmt.Scanln(&user)
        if user != "" {
            userNameSet = true
        } else {
            fmt.Println("You did not enter a valid name.")
        }
    }

// loop through opponent selection
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

// loop through core-gameplay (number guessing)
    for tries < 3  {
        fmt.Println("Enter a number between 1 and 10.")
        guessedNum = getInput()
        tries++

        if guessEqualsInput(inputNum, guessedNum) == true {
            fmt.Println("Good job, your guess was correct.")
            fmt.Println("The entered number was", inputNum, ".")
            fmt.Println("You've earned 1 point.")
            score++
            writeToFile(user, score)
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
        fmt.Println("There will be no point for you. Good luck next time.")
    }
//    printUserScores(getUserScores(user))
      fmt.Println("Userscore should appear here.")
}

func getInput() int {
    var inputNum int
    var isInputCorrect bool = false

    for isInputCorrect == false {
        fmt.Scanln(&inputNum)
        if inputNum <= 10 && inputNum >= 1 {
            isInputCorrect = true
        } else {
            fmt.Println("Nope. You're way out of scope. Stay between 1 and 10.")
        }
    }
    return inputNum
}

/* check if the guesses number is the same as the number that was entered
   by the AI or the human opponent. */
func guessEqualsInput(inputNum int, guessedNum int) bool {
    if guessedNum == inputNum {
        return true
    } else {
    return false
    }
}

// create random data for the pseudo-artificial opponent
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

// scoring should end up in user-specific files as csv
func writeToFile(user string, score int) {
    fmt.Println("writing your score to file...")
    f, err := os.OpenFile("/Users/dictvm/go/tmp/"+user, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("Error: something went wrong when writing your score.")
        fmt.Println(err)
    }
    n, err := f.WriteString(user + "," + strconv.Itoa(score) + "\n")
    if err != nil {
        fmt.Println("Error 2: Oh noez. ")
    fmt.Println(n, err)
    }
    f.Close()
}

/* let's grab the raw score data from the written files and analyse them for
   highscore-handling
func getUserScores(user string) []string {
    reader := csv.NewReader("/Users/dictvm/go/tmp/"+user)
    content, err := ioutil.ReadFile("/Users/dictvm/go/tmp/"+user)
    if err != nil {
        fmt.Println("Error: ")
        fmt.Println(err)
    }
}
*/

func getUserScores(user string) {
    file, err := os.Open("/Users/dictvm/go/tmp/"+user)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()
    reader := csv.NewReader(file)
    for true {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Println(record) // record has the type []string
    }
}

/* this should not longer be needed because we're using native encoding/csv
    csvIndex := strings.Split(string(content), ",")
    return csvIndex
*/

func printUserScores(highScore []string) {
    fmt.Println("#####################")
    fmt.Println("Username: " +highScore[0])
    fmt.Println("Highscore: " +highScore[1])
    fmt.Println("#####################")
}
