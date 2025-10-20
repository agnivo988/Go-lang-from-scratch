package arrayptrs

import "testing"

func TestSum(t *testing.T) {
       numbers := [5]int{1,2,3,4,5}

	   got := Add(numbers)
	   want := 15
	   if got != want {
		t.Errorf("got %d instead of %d for %v",got,want, numbers)
	   }

}

func Add(numbers [5] int) int {
	sum := 0
	for _,number := range numbers {
          sum += number
	}
	return sum
}