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