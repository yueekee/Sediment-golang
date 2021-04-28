package main

import (
	"fmt"
	"testing"
)

/*
如果Wallet{}对应的函数为值引用函数，这里无法通过测试。

*/

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		fmt.Println("address of balance in test is", &wallet.balance)

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)

		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(100)

		assertBalance(t, wallet, Bitcoin(20))

		assertError(t, err, "cannot withdraw, insufficient funds")
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s but want %s", got, want)
	}
}

/*
对err的情况进行分类处理
*/
func assertError(t *testing.T, err error, want string) {
	if err == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if err.Error() != want {
		t.Errorf("err '%s', want '%s'", err, want)
	}
}
