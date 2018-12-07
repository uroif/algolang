package lib

import "fmt"

func DoPrint(c string, number int) {
	for i := 0; i < number; i++ {
		fmt.Printf("%s", c)
	}
}
