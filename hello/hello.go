package main

import "fmt"

// const englishhelloPrefix = "Hello, "

const (
	spanish = "Spanish"
	french = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)


func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return grettingPrefix(language) + name
}

func grettingPrefix(language string)(prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish: 
		prefix = spanishHelloPrefix
	default: 
		prefix = englishHelloPrefix
	}
	return 
}

func main() {
	fmt.Println(Hello("mintu",""))
	fmt.Println(Hello("","Spanish"))
}
