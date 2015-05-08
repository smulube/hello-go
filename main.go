package main

import (
	"fmt"
	"strings"
)

func greet(name string) string {
	s := []string{"hello,", name, "\n"}
	return strings.Join(s, " ")
}

func main() {
	fmt.Printf(greet("world"))
}
