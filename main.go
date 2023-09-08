package main

import (
	"fmt"
	"learn_go/hello"
	"os"
)

func main() {
	fmt.Println(hello.Say(os.Args[1:]))
}
