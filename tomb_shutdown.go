package main

import (
	"errors"
	"fmt"

	"gopkg.in/tomb.v2"
)

func main() {
	t := &tomb.Tomb{}
	// start workers
	for i := 0; i < 5; i++ {
		t.Go(func() error {
			for {
				select {
				case <-t.Dying():
					fmt.Printf("dying because: %v\n", t.Err())
					return nil
				}
			}
		})
	}
	// START OMIT
	// start bad worker
	t.Go(func() error { // HL
		return errors.New("bad worker died")
	})

	err := t.Wait()
	fmt.Println(err)
	// END OMIT

}
