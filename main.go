package main

import "fmt"

func main() {
    for number := 0; number < 101; number++ {
        println(FizzBuzz(number))
    }
    for number := 0; number < 101; number++ {
        println(FizzBuzzBang(number, Fizz, Buzz, Bang))
    }
}

func FizzBuzz(number int) string {
    result := fmt.Sprintf("%d", number)

    if number % 3 == 0 && number % 5 == 0 {
        result = "FizzBuzz"
    } else if number % 3 == 0 {
        result = "Fizz"
    } else if number % 5 == 0 {
        result = "Buzz"
    }
    
    return result
}

type BangFunc func(number int) (string, bool)


func Fizz(number int) (string, bool) {
    if number % 3 == 0 {
        return "Fizz", true
    }

    return fmt.Sprintf("%d", number), false
}


func Buzz(number int) (string, bool) {
    if number % 5 == 0 {
        return "Buzz", true
    }

    return fmt.Sprintf("%d", number), false
}


func Bang(number int) (string, bool) {
    if number % 7 == 0 {
        return "Bang", true
    }

    return fmt.Sprintf("%d", number), false
}


func FizzBuzzBang(number int, bangs ... BangFunc) string {
    result := ""
    matches := false

    for _, bang := range bangs {
        if bangResult, bangMatches := bang(number); bangMatches {
            result += bangResult

            matches = matches || bangMatches
        }
    }

    if !matches {
        result = fmt.Sprintf("%d", number)
    }

    return result
}
