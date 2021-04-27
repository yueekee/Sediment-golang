package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

/*
在 Go 中不能对切片使用等号运算符。你可以写一个函数迭代每个元素来检查它们的值。
但是一种比较简单的办法是使用 reflect.DeepEqual，它在判断两个变量是否相等时十分有用。
需要注意的是 reflect.DeepEqual 不是「类型安全」的，所以有时候会发生比较怪异的行为。
比如比较slice和string，可以通过编译。
*/
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// 防止空切片的错误
func TestSumAllTails(t *testing.T) {
    checkSums := func(t *testing.T, got, want []int) {
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }

    t.Run("make the sums of some slices", func(t *testing.T) {
        got := SumAllTails([]int{1,2}, []int{0,9})
        want := []int{2, 9}

        checkSums(t, got, want)
    })

    t.Run("safely sum empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{3, 4, 5})
        want :=[]int{0, 9}

        checkSums(t, got, want)
    })
}
