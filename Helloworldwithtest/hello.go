package main 
import "fmt"

const englishHello = "Hello, "

func hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHello + name
}

func main(){
	fmt.Println(hello("World"))
}