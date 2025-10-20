package main

import (
	"fmt"
)

// const bengali = "Bengali"
// const french = "French"
// const spanish = "Spanish"
// const spanishhello = "Hola, "
// const englishHello = "Hello, "
// const frenchhello = "Bonjour, "
// const bengalihello = "Nomoskar, "

//refactoring the code further for better readability
const (
	bengali = "Bengali"
	spanish = "Spanish"
	french = "French"
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix = "Bonjour, "
	spanishHelloPrefix = "Hola, "
	bengaliHelloPrefix = "Nomoskar, "
)

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetPrefix(language)+name
}

func greetPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case bengali:
	    prefix = bengaliHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main(){
	fmt.Println(hello("World",""))
}