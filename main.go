package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		Hello world in Go
		Note: use always fmt for production code.
		Just println can be helpful for development but far from
		being as good as the fmt library
	*/
	fmt.Println("Hello world!")

	/*
		Data types
		Booleans: defined as bool -> false / true
		Ints -> can be int, int8, int16, int32, int64
		Unsigned Ints -> can be uint, uint8, uint16, uint32, uint64
		Floats -> can be float32, float64
		Strings -> defined as string
	*/

	var boolean bool = true
	fmt.Println(boolean)
	var num int = 123
	fmt.Println(num)
	var float float32 = 123.123
	fmt.Println(float)
	var bigFloat float64 = 123456.123456
	fmt.Println(bigFloat)
	var text string = "I'm a string!"
	fmt.Println(text)

	/*
		Variables can have inferred type when you start then
		with the value right away.

		Also, you can simplify the attribution by using :=
		x := y means var x (type of y) = y
	*/

	var var1 string = "Text" //Traditional way of declaration
	var var2 = "Text"        //Inferred type string
	var3 := "Text"           //Inferred type string and simplified syntax
	fmt.Println(var1, var2, var3)

	/*
		Go also supports initializing multiple variables
		at once.
	*/
	var var5, var6 int = 1, 2 //Traditional way of declaration
	var var7, var8 = 1, 2     //Inferred type string
	var9, var10 := 1, 2       //Inferred type string and simplified syntax
	fmt.Println(var5, var6, var7, var8, var9, var10)

	print("Hello functions!")
	fmt.Println(intDivision(5, 2))
	fmt.Println(intDivisionWithRemainder(5, 2))

	var result, remainder, err = intDivisionWithRemainderAndErrorHandling(5, 0)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	var11 := 5

	switch var11 {
	case 1:
		fmt.Println("var11 is equal to 1")
	case 2:
		fmt.Println("var11 is equal to 2")
	case 3, 4:
		fmt.Println("var11 is equal to 3 or 4")
	default:
		fmt.Printf("var11 is equal to %v\n", var11)
	}

	/*
		Arrays
		- Fixed length
		- Same type
		- Indexable
		- Contiguous in Memory
	*/
	var intArr [3]int32 = [3]int32{1, 2, 3} //arrays has 3 positions of int32 types
	fmt.Println(intArr)

}

func print(value string) {
	fmt.Println(value)
}

func intDivision(numerator int, denominator int) int {
	return numerator / denominator
}

/*
Functions can have multiple returns! Thats freaking awesome
*/
func intDivisionWithRemainder(numerator int, denominator int) (int, int) {
	return numerator / denominator, numerator % denominator
}

/*
Error handling in Go is provided by the error type
*/
func intDivisionWithRemainderAndErrorHandling(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	return numerator / denominator, numerator % denominator, err
}
