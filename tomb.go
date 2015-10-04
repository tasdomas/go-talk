package main

import (
	"errors"
	"fmt"

	"gopkg.in/tomb.v2"
)

// START OMIT
func main() {
	t := &tomb.Tomb{}
	for i := 0; i < 5; i++ {
		t.Go(func() error {
			for {
				select {
				case <-t.Dying(): // HL
					fmt.Printf("dying because: %v\n", t.Err()) // HL
					return nil                                 // HL
				}
			}
		})
	}

	// shutdown
	t.Kill(errors.New("shutdown")) // HL
	err := t.Wait()
	fmt.Println(err)
}

// END OMIT
