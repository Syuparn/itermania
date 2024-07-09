package itermania

import (
	"fmt"
	"strconv"
)

func ExamplePrimeNumbers() {
	prime := Bind(Inc(2), func(n int) Gen[int] {
		return Where(Const(n), All(Not(Eq(Mod(Const(n), Range(2, n, 1)), Const(0)))))
	})

	for i := range Head(prime, 10)() {
		fmt.Println(i)
	}
	// Output:
	// 2
	// 3
	// 5
	// 7
	// 11
	// 13
	// 17
	// 19
	// 23
	// 29
}

func ExampleFizzBuzz() {
	fizzbuzz := Bind(Inc(1), func(n int) Gen[string] {
		return If(Eq(Mod(Const(n), Const(15)), Const(0)), Const("FizzBuzz"),
			If(Eq(Mod(Const(n), Const(3)), Const(0)), Const("Fizz"),
				If(Eq(Mod(Const(n), Const(5)), Const(0)), Const("Buzz"),
					Const(strconv.Itoa(n)))))
	})

	for i := range Head(fizzbuzz, 15)() {
		fmt.Println(i)
	}
	// Output:
	// 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}
