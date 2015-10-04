package main

import (
	"fmt"
	"os"

	"github.com/juju/errors"
)

func main() {
	// START OMIT
	err := errors.Annotate(os.ErrNotExist, "wrapped error")

	// the error
	fmt.Println(err) // HL
	// the cause
	println("")                    //OMIT
	fmt.Println(errors.Cause(err)) // HL
	// the stack
	println("")                         //OMIT
	fmt.Println(errors.ErrorStack(err)) // HL

	//END OMIT
}
