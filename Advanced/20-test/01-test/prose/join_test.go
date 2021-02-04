package prose

import (
	"fmt"
	"testing"
)

// 使用go test的一些约定条件：
// 可以不需要将测试代码与正在测试的代码放在一个包里，但是如果想从包中访问未导出的类型或者函数，则需要这样做；
// 测试文件的名称以_test结尾；
// 测试函数应该以Test开头；
// 需要导入"testing"的包；
// 测试函数应该接收单个参数：一个指向testing.T值的指针；
// 可以通过对testing.T值调用方法(例如Error)来报告测试失败。(例如：t.Error("no test written yet"))

// 使用 go test github.com/liankui/Sediment-golang/Advanced/20-test/01-test/prose
// 注意使用带有github.com的路径，不然会报错：package prose is not in GOROOT (/usr/local/go/src/prose)

// 这里测试的结果是希望有2个参数时，表示为"apple and orange"
// 有3个时，表示为"apple, orange, and pear"
func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	got := JoinWithCommas(list)
	want := "apple"
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	got := JoinWithCommas(list)
	want := "apple and orange"
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	got := JoinWithCommas(list)
	want := "apple, orange, and pear"
	if got != want {
		t.Error(errorString(list, got, want))
	}
}

// helper函数：用来减少重复代码
func errorString(list []string, got, want string) string {
	return fmt.Sprintf("JoinWithCommas（%#v）= \"%s\", want \"%s\"", list, got, want)
}
