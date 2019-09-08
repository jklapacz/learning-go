package main

import (
	"fmt"
	"strings"
)
func main() {
	greetings := "\t Hello, world "
	fmt.Printf("%d %s\n", len(greetings), greetings)
	trimmed := strings.TrimSpace(greetings)
	fmt.Printf("%d %s\n", len(trimmed), trimmed)
	helloWorld := greetings[6:]
	fmt.Println(helloWorld)
	// replace part of string with another value
	replaceWorld := "Hello, World. How are you World?"
	helloMars := strings.Replace(replaceWorld, "World", "Mars", -1)
	fmt.Println(helloMars)
	// escaping characters
	escapeWorld := "Hello world, this is \"Kuba\" \\t\\n Hello"
	fmt.Println(escapeWorld)
	// capitalize values
	capitalizeWorld := "hello world, capitalize me. i want another capitalization"
	titledWorld := strings.Title(capitalizeWorld)
	fmt.Println(titledWorld)
	upperWorld := strings.ToUpper(capitalizeWorld)
	fmt.Println(upperWorld)
}
