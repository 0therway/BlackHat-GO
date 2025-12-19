package main

import (
	"fmt"
)

func main() {
	var f string = "file1"
	i := 255

	fmt.Printf("%T\n", f)
	fmt.Printf("%v\n", f)
	fmt.Printf("%#v\n", f)
	fmt.Printf("%v%%\n", f)

	fmt.Print("\n")

	fmt.Printf("%T\n", i)
	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)

}
