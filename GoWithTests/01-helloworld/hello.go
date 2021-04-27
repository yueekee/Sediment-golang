package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const franch = "Franch"
const franchHelloPrefix = "fanch, "

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	} 

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language{
		case spanish: 
			prefix = spanishHelloPrefix
		case franch:
			prefix = franchHelloPrefix
		default:
			prefix = englishHelloPrefix
		}
	return
}

func main() {
	fmt.Println(Hello("eric", ""))
}
