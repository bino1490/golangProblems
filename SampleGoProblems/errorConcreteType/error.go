// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

//In Go, error is a built-in interface used for error handling. It’s simple by design.
//type error interface {
//    Error() string
//}

type userError struct {
	name string
}

func (u userError) Error() string {
	return fmt.Sprintf("error %s", u.name)
}

func getError(name string) error {
	ue := userError{name: name}
	return ue
}

func main() {
	// getError("Ramu") returns an error interface that contains a userError value.
	// When .Error() is called (explicitly or by fmt.Println), Go dynamically dispatches to userError.Error() because it’s the concrete type inside the interface.

	// fmt.Println has special handling for error
	//if value implements error {
	//	call Error()
	//}
	err := getError("Vijay").Error() // explicit
	err1 := getError("Ramu")         // implicit

	fmt.Println(err)
	fmt.Printf("%T\n", err)

	fmt.Println(err1)
	fmt.Printf("%T\n", err1)
}
