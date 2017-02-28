package main

import (
    "fmt"
    "testing"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/stretchr/testify/assert"
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

func TestFizzBuzzBang(t *testing.T) {
    Convey("returns numbers when there are no bangs", t, func() {
        Convey("given input number", func() {
            inputNumber := 3
            Convey("when calculates fizzbuzzbang without any bangs", func() {
                gotResult := FizzBuzzBang(inputNumber)
                Convey("then returns input number", func() {
                    So(gotResult, ShouldEqual, fmt.Sprintf("%d", inputNumber))
                })
            })
        })
    })

    Convey("returns bang when bang matches the input number", t, func() {
        Convey("given input number matching by bang", func() {
            inputNumberMatchingByBang := 3
            Convey("and bang matches the input number", func() {
                bangValue := "bang"
                bang := func (number int) (string, bool) {
                    return bangValue, true
                }
                Convey("when calculates fizzbuzzbang with bang that matches input number", func() {
                    gotResult := FizzBuzzBang(inputNumberMatchingByBang, bang)
                    Convey("then returns bang", func() {
                        So(gotResult, ShouldEqual, bangValue)
                    })
                })
            })
        })
    })
}

func TestReturnsFizzBangWhenBothMatchSameNumber(t *testing.T) {
    // given
    inputNumberMatchingByFizzAndBang := 3

    // and
    bangValue := "bang"
    bang := func (number int) (string, bool) {
        return bangValue, true
    }

    // and
    fizzValue := "fizz"
    fizz := func (number int) (string, bool) {
        return fizzValue, true
    }

    // when
    gotResult := FizzBuzzBang(inputNumberMatchingByFizzAndBang, fizz, bang)

    // then
    assert.Equal(t, fizzValue + bangValue, gotResult)
}
