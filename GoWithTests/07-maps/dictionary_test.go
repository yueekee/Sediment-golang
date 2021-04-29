package main

import "testing"

func TestSearch(t *testing.T) {
	d := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T){
		got, _ := d.Search("test")
		want := "this is just a test"
	
		assertStrings(t, got, want)
	} )

	t.Run("unknown word", func(t *testing.T){
		_, err := d.Search("unknown")
	
		// assertStrings(t, got, want)
		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	d := Dictionary{}
	d.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := d.Search("test")
	assertStrings(t, got, want)
	assertError(t, err, ErrNotFound)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
        t.Errorf("got '%s' want '%s' given, '%s'", got, want, "test")
    }
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
        t.Errorf("got err '%s' want '%s'", got, want)
    }
}