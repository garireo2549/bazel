package main

import (
	"fmt"
	"strconv"
)

const (
	fizz = "FIZZ!"
	buzz = "BUZZ!"
	fbzz = "FIZZ BUZZ!"
)

func main() {
	fmt.Println(fizzBuzz(15))
}

func fizzBuzz(i int) string {
	switch {
	case i%15 == 0:
		return fbzz
	case i%3 == 0:
		return fizz
	case i%5 == 0:
		return buzz
	}

	return strconv.Itoa(i)
}
