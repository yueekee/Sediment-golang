package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}


/* 
基准测试
testing.B 可使你访问隐性命名（cryptically named）b.N。
基准测试运行时，代码会运行 b.N 次，并测量需要多长时间。
用 go test -bench=. 来运行基准测试
*/
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

/* 
goos: darwin
goarch: amd64
pkg: github.com/liankui/Sediment-golang/GoWithTests/03-iteration
BenchmarkRepeat-16       9199012               128 ns/op
PASS
*/