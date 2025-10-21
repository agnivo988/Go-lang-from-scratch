package intezers

import (
	// "fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(19, 26)
	expected := 45
	if sum != expected {
		t.Errorf("sum %d expected %d", sum, expected)
	}
}

// func Add(x, y int) int {
// 	return x + y
// }

// func sampleAdd() {
// 	sum := Add(3, 4)
// 	fmt.Println(sum)

// }
