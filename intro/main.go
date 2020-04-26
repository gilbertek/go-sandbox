package main

import (
	"fmt"
)

// const (
// 	first  = iota
// 	second = iota
// )

// const (
// 	first = iota + 6
// 	second
// )

const (
	first = iota
	second
	third
)

func main() {
	// use var to declare a variable
	var i int = 42
	var f float32 = 3.14

	fmt.Println(i)
	fmt.Println(f)

	// Using the := we can also declare and initialize a variable in go
	firstName := "Arthur"
	fmt.Println(firstName)

	// Boolean types
	b := true
	fmt.Println(b)

	// Complex datatype
	// We can do complex math with buitins tools in go.
	c := complex(3, 4)
	fmt.Println(c)

	// multiple assigment with real value
	r, im := real(c), imag(c)
	fmt.Println(r)
	fmt.Println(im)
	fmt.Println(r, im)

	// Pointer
	/*
		  Dereference Operation
			To deferrence a pointer, tell the compiler to reach into a pointer and grab a value
			You deferrence a variable by preceding the variable with an asterisk (*)

			A pointer is a variable that stores the address of a value, rather than the value itself.
			If you think of a computer's memory (RAM) as a JSON object, a pointer would be like the key, and a
			normal variable would be the value

			1. Creating a pointer
			& point an address in memory
			nString := "Alpha"
			pString := &nString

			2. Describing a pointer

			In a function signature or type definition, the * is used to designate that a value is a pointer
			func passPointer(pointer *string)

			3. Dereferencing a pointer
			It can be slightly confusing, but the * is used to describe a pointer and it's also used as an
			operator to deferrence a pointer.

			func derefPointer(pointer *string) {
				nString := *pointer
			}

			# When to use a Pointer
			- A function that mutates one of its parameters
			- Better performance
				If you have a string that is really expensice to copy, using a pointer may be worthwhile.
			- Need a Nil Value option


	*/
	var lastName *string = new(string)
	*lastName = "King"
	fmt.Println(*lastName)

	// The address of an operator
	nFirstName := "Arthur"
	fmt.Println(nFirstName)

	ptr := &nFirstName
	fmt.Println(ptr, *ptr)

	nFirstName = "Tricia"
	fmt.Println(ptr, *ptr)

	// Constant
	// Constant work the same way as variable do, except constants do not change their value as variable does
	// We have to initialize them at time of declaration
	// The value of the constant has to be able to be determined at compile time.
	const pi = 3.1415
	fmt.Println(pi)

	// iota and constant expression
	// The first time we run with iota, we get the value of 0, the next time the value increment to 1.
	// Everytime an iota is reuse, it actually increment the value by 1.
	// We can also use iota in what is called constant expression.
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)

	// Collections in go
	// Array => Fixed size collection of similar things.
	// Slices => Dynamic size array
	// Maps => key value relationship
	// Struct => Data size of a class and define a concept.

	// Array => Fixed size collection of similar things. (Integers/bool)
	// We are declaring an array variable arr that is 3 element array of integers
	// All values have to be of the same type
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr)
	fmt.Println(arr[1])

	fmt.Println(arr2)
	fmt.Println(arr2[1])

	// Slices => Dynamic size array
	// Slices are build on top of array but are more flexible
	// All values have to be of the same type

	// We are creating a slice datatype from the beginning of arr2 to the end.
	// The slice is pointing to the data that the array is holding
	slice := arr[:]

	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	// When we don't provide the size, the compiler know this is a
	// slice and will manage the underline array for us.
	slic := []int{1, 2, 3}

	fmt.Println(slic)
	slic = append(slic, 34, 42, 27)
	fmt.Println(slic)

	// We can also create slice of slices
	s2 := slic[1:]  // Start at index 1 to the end
	s3 := slic[:2]  // Start at index 0 upto and not including index 2
	s4 := slic[1:2] // Start at index 1 and upto and not including index 2
	fmt.Println(slic, s2, s3, s4)

	// Maps => key value relationship
	// All of keys have to be one type and all values have to be one type
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	// Struct
	// Only collection type that allow us to associate disparate datatype together
	// The fields a fixed at compile time
	type user struct {
		ID        int
		FirstName string
		Lastname  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.Lastname = "Dent"
	fmt.Println(u)

	u2 := user{ID: 2, FirstName: "John", Lastname: "Doe"}
	fmt.Println(u2)

	// Interfaces
	/*
		Interfaces allow us to treat different types as the same typess when it comes to specific behaviors.
		They are central to a Go Programmers toolbelts.

		Interfaces are named collections of method signatures, they are how we achieve a kind of polymorphism in Go.
		type error interface {
			Error() string
		}

		By definition, the error interface encapsulates any type that has an Error() method defined on it.
		It accepts no parameters, and returns a string

		For example
		type networkIssue struct {
			message string
			code int
		}

		func (ni networkIssue) Error() string {
			return fmt.Sprintf("Network error! message: %s, code: %v", np.message, np.code)
		}

		Now we can use an instance of the networkIssue struct whenever an error occurs


		Interfaces are not classes, they are slimmer
		Interfaces don't have constructors or destructors that require that setup/destruction data
		Interfaces are not hierachical by natue
		Interfaces define function signatures, but not underlyong behavior

	*/

	// Looping
	// There's only one loop construct in go the for loop
	/*
		  - Loop till condition
			- Infinite loops
			- Loop till condition with pause clause
			- Loop over collections
	*/

	// Looping till condition
	// var idx int
	// for idx < 5 {
	// 	println(idx)
	// 	idx++
	//
	// 	if idx == 3 {
	// 		continue
	// 	}
	// 	println("Continuing...")
	// }

	// - Loop till condition with pause clause
	// for idx := 0; idx < 5; idx++ {
	// 	println(idx)
	// }

	// - Infinite loops
	// var idx int
	// for {
	// 	if idx == 15 {
	// 		break
	// 	}
	// 	println(idx)
	// 	idx++
	// }
	// - Loop over collections
	// slixe := []int{1, 2, 3, 4}
	// for i, v := range slixe {
	// 	println(i, v)
	// }
	//
	// knownPorts := map[string]int{"http": 80, "https": 443}
	// for k, v := range knownPorts {
	// 	println(k, v)
	// }

	// PANIC
	// panic("Something bad just happened")

}
