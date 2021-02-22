package prose

import "testing"

// 表驱动测试是处理输入和预期输出的"表"的测试。
// 它们将每一组输入传递给正在测试的代码，并检查代码的输出是否与预期值匹配。
// 表驱动测试最好的一点是，但需要新的测试时，很容易添加它们。

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{list: []string{}, want: ""},	// 传入一个空切片会导致panic，这时就需要更新被测试的代码
		{list: []string{"apple"}, want: "apple"},
		{list: []string{"apple", "orange"}, want: "apple and orange"},
		{list: []string{"apple", "orange", "pear"}, want: "apple, orange, and pear"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			t.Errorf("JoinWithCommas（%#v）= \"%s\", want \"%s\"", test.list, got, test.want)
		}
	}
}