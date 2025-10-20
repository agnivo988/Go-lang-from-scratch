package main 
import "fmt"

const bengali = "Bengali"
const french = "French"
const spanish = "Spanish"
const spanishhello = "Hola, "
const englishHello = "Hello, "
const frenchhello = "Bonjour, "
const bengalihello = "Nomoskar, "

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishhello + name
	}
    if language == french {
		return frenchhello + name 
	}
	if language == bengali {
		return bengalihello +name
	}

	return englishHello + name
}

func main(){
	fmt.Println(hello("World",""))
}