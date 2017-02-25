package main

import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

func TestFizBuzz(t *testing.T) {

    Convey("Given numbers that are multipliers of three", t, func() {
        numberWhichIsMultiplierOfThree := 3

        Convey(fmt.Sprintf("when calculates fizzbuzz for: %d", numberWhichIsMultiplierOfThree), func() {
            result := FizzBuzz(numberWhichIsMultiplierOfThree)

            Convey("returns Fizz", func() {
                So(result, ShouldEqual, "Fizz")
            })
        })
    })

    Convey("Given numbers that are multipliers of five", t, func() {
        numberWhichIsMultiplierOfFive := 5

        Convey(fmt.Sprintf("when calculates fizzbuzz for: %d", numberWhichIsMultiplierOfFive), func() {
            result := FizzBuzz(numberWhichIsMultiplierOfFive)

            Convey("returns Buzz", func() {
                So(result, ShouldEqual, "Buzz")
            })
        })
    })

    Convey("Given numbers that are multipliers of five and three", t, func() {
        for numberWhichIsMultiplierOfFiveAndThree := 15; numberWhichIsMultiplierOfFiveAndThree < 151; numberWhichIsMultiplierOfFiveAndThree += 15 {

            Convey(fmt.Sprintf("when calculates fizzbuzz for: %d", numberWhichIsMultiplierOfFiveAndThree), func() {
                result := FizzBuzz(numberWhichIsMultiplierOfFiveAndThree)

                Convey("returns FizzBuzz", func() {
                    So(result, ShouldEqual, "FizzBuzz")
                })
            })
        }
    })

    Convey("Given number isn't multipler of three and five", t, func() {
        numberWhichIsntMultiplierOfThreeAndFive := 4

        Convey(fmt.Sprintf("when calculates fizzbuzz for: %d", numberWhichIsntMultiplierOfThreeAndFive), func() {
            result := FizzBuzz(numberWhichIsntMultiplierOfThreeAndFive)

            Convey("returns input number", func() {
                So(result, ShouldEqual, fmt.Sprintf("%d", numberWhichIsntMultiplierOfThreeAndFive))
            })
        })
    })
}
