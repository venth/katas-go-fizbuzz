package main

import "fmt"

func main() {
    for number := 0; number < 101; number++ {
        println(FizzBuzz(number))
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

type Bang func(number int) (string, bool)

func FizzBuzzBang(number int, bangs ... Bang) string {
    result := fmt.Sprintf("%d", number)

    for _, bang := range bangs {
        if bangResult, bangMatches := bang(number); bangMatches {
            result = bangResult
        }
    }

    return result
}
