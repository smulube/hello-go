package main

import (
	"fmt"
	"strings"
)

func greet(name string) string {
	s := []string{"hello", name}
	return strings.Join(s, ", ")
}

func main() {
	fmt.Printf(strings.Join([]string{greet("world"), "\n"}, ""))
}
