package intezers

import "fmt"

func Add(x, y int) int {
	return x + y
}

func sampleAdd() {
	sum := Add(3, 4)
	fmt.Println(sum)
}