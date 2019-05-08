package main

import (
	"regexp"
	"fmt"
)

func main() {
	n := "Chapter 01: kdjasfdlk"
	exp := regexp.MustCompile(`(\d+)`)
	match := exp.FindStringSubmatch(n)
	fmt.Printf("%q\n", match)
	fmt.Printf("%s\n", match[1])
	fmt.Println(len(match))
}