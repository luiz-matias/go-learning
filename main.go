package main

import (
	"errors"
	"fmt"

	"github.com/luiz-matias/go-learning/structs"
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

	/*
		Arrays
		- Fixed length
		- Same type
		- Indexable
		- Contiguous in Memory
	*/
	var intArr [3]int32 = [3]int32{1, 2, 3} //arrays has 3 positions of int32 types
	fmt.Println(intArr)

	/*
		Loops
	*/

	//Traditional for
	sum := 0
	for i := 0; i < 4; i++ {
		sum += i
	}
	fmt.Printf("Sum of values is %v\n", sum)

	//While equivalent in go: For without init and post statements
	counter := 0
	//While sum is less than 100
	sum = 1
	for sum < 100 {
		//sum receives it's own value
		sum += sum
		counter++
	}
	fmt.Printf("Sum of values for the continued loop is %v (%v interactions)\n", sum, counter)

	//infinite loop
	sum = 1
	for {
		sum += sum
		break
	}

	/*
		Conditional statements
	*/
	if 10 > 5 {
		fmt.Println("10 is greater than 5")
	}

	//If with short statement (assigning variables before conditional)
	if v := 1; v < 5 {
		fmt.Println("1 is lower than 5")
	}

	//Switch (that also supports short statement)
	//Switchs evaluation is from top to bottom, so if z is equal to 1, other cases
	//won't be executed
	switch z := 2; z {
	case 1:
		fmt.Println("z is 1")
	case 2, 3, 4:
		fmt.Println("z is 2, 3 or 4")
	default:
		fmt.Println("z is something else")
	}

	//Switchs without condition can be used to write long if chains
	switch {
	case 1 < 2:
		fmt.Println("1 is lower than 2")
	default:
		fmt.Println("default behavior")
	}

	//Defer statements
	//Can be used to execute a function only when the sorrounding function returns
	//Useful to close database connections, for example
	deferFunction()

	//Stacked defer statements works as a stack
	//Last IN First OUT
	//0123456789 -> 9876543210
	stackedDeferFunction()

	/*
		Pointers
	*/
	pointers()

	person := structs.CreatePerson("Luiz", 10)
	fmt.Println(person)
}

func pointers() {
	var p *int
	i := 42 //i holds the value of 42
	p = &i  //p points to the memory address of i

	fmt.Printf("Memory address: %v\n", p)
	fmt.Printf("Value: %v\n", *p)
}

func deferFunction() {
	defer fmt.Println("Calling defer close statement")
	fmt.Println("Starting function defer")
}

func stackedDeferFunction() {
	fmt.Println("Starting stack function defer")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done!")
}

func print(value string) {
	fmt.Println(value)
}

func intDivision(numerator int, denominator int) int {
	return numerator / denominator
}

/*
Functions can have multiple returns
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
