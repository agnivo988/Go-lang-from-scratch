package arrayptrs

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
    //    numbers := [5]int{1,2,3,4,5}

	//    got := Add(numbers)
	//    want := 15
	//    if got != want {
	// 	t.Errorf("got %d instead of %d for %v",got,want, numbers)
	//    }

	t.Run("with 5 element size",func(t *testing.T) {
		numbers := [] int {1,2,3,4,5}
		got := Add(numbers)
		want := 15
		if got != want {
			t.Errorf("wanted %d but got %d for %v",want,got,numbers)
		}
	})

	t.Run("with collection of any size",func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Add(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d wanted %d for %v",got,want,numbers)
		}

	})

}

func Add(numbers []int) int {
	sum := 0
	for _,number := range numbers {
          sum += number
	}
	return sum
}

func TestSumAll(t *testing.T) {
	got := SumAll ([]int{0,1},[]int{0,9})
	want := []int{1,9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Wanted %v got %v ",want,got)
	}
}

func SumAll (numbersTosum ... []int) []int {
           lengthOfNumbers := len(numbersTosum)
		   sums := make([]int, lengthOfNumbers)
		   for i, numbers := range numbersTosum {
			sums[i] = Add(numbers)
		   }
		   return sums
}