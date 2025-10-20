//Go has inbuilt feature of testing no exteral libraries needed
package main

import (
	"testing"
)

// func TestHello(t *testing.T) {
// 	got := hello("Agniva")
// 	want := "Hello, Agniva"

// 	if got!=want {
// 		t.Errorf("got %q want %q",got,want)
// 	}
// }

/*Here, we are introducing another tool in our testing arsenal: subtests. 
Sometimes, it is useful to group tests around a "thing" and then
 have subtests describing different scenarios.*/

func TestHello(t *testing.T) {
	t.Run("say Hello to people",func(t *testing.T){
		got := hello("Agniva","")
		want := "Hello, Agniva"
		if got != want {
			t.Errorf("got %q want %q",got,want)
		}
		assertCorrectMessage(t,got,want)
	})
	t.Run("say 'Hello, world' when empty string is supplied",func(t *testing.T) {
		got := hello("","")
		want := "Hello, World"
		if got!= want {
			t.Errorf("got %q want %q",got, want)
		}
		assertCorrectMessage(t,got,want)
	})

	t.Run("say in spanish",func(t *testing.T) {
		got := hello("Elodie","Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t,got, want)
	})

	t.Run("say in french", func(t *testing.T) {
		got := hello("Macron","French")
		want := "Bonjour, Macron"
		assertCorrectMessage(t,got,want)
	})

	t.Run("say in Bengali", func(t *testing.T) {
		got := hello("Riju","Bengali")
		want := "Nomoskar, Riju"
		assertCorrectMessage(t, got , want)
	})
}

func assertCorrectMessage(t testing.TB, got,want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q",got,want)
	}
}