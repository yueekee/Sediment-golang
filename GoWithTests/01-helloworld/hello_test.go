package main

import "testing"

/* 
https://studygolang.gitbook.io/learn-go-with-tests/
t.Helper() 需要告诉测试套件这个方法是辅助函数（helper）。
通过这样做，当测试失败时所报告的行号将在函数调用中而不是在辅助函数内部。
这将帮助其他开发人员更容易地跟踪问题。如果你仍然不理解，请注释掉它，使测试失败并观察测试输出。
 */

func TestHello(t *testing.T) {
	// got := Hello("eric")
	// want := "hello world"

	// if got != want {
	// 	t.Errorf("got '%q', want '%q'", got, want)
	// }

	// t.Run("saying hello to people", func(t *testing.T){
	// 	got := Hello("eric")
	// 	want := "Hello, eric"

	// 	if got != want {
	// 		t.Errorf("got '%q', want '%q'", got, want)
	// 	}
	// })

	// t.Run("saying hello world when an empty string is supplied", func(t *testing.T){
	// 	got := Hello("")
	// 	want := "Hello world"

	// 	if got != want {
	// 		t.Errorf("got '%q', want '%q'", got, want)
	// 	}
	// })

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q', want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("eric", "")
		want := "Hello, eric"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	
	
}
