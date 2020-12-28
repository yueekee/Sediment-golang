// 这个实例展示了关于flag包的更复杂的使用
package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

// 实例1：一个单独的字符串flag，叫“species”，其默认值为“goher”
var species = flag.String("species", "gopher", "the species we are studying")

// 实例2： 两个flag分享一个变量，所以我们可以一起写
// 初始化顺序没有定义，所以可以同时使用两个默认值。这必须在初始化函数中定义。
var gopherType string

func init() {
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")
}

// 实例3：用户定义flag类型，一个时间段的切片
type interval []time.Duration

// String是一个用来格式化flag值(02_flag.Value接口的一部分)的方法
// String方法的输出将被用于调试
func (i *interval) String() string {
	return fmt.Sprint(*i)
}

// Set是一个用来设置flag值(02_flag.Value接口的一部分)的方法
// Set的参数是String类型，用于设置为flag
// 这是一个以逗号为分隔符的数组，我们需要分离它
func (i *interval) Set(value string) error {
	// 如果flag能被设置为多时间，加速度值，如果有如此声明，我们将会删除这些
	// 这些将会允许很多组合，例如"-deltaT 10s -deltaT 15s"
	if len(*i) > 0 {
		return errors.New("interval 02_flag already set")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*i = append(*i, duration)
	}
	return nil
}

// 将一个flag定义为堆积期间。因为它还有个特殊类型，我们需要使用Var函数，从而在初始化中创建flag
var intervalFlag interval

func init() {
	// 将命令行flag与intervalFlag绑定，并设置使用信息
	flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {
	// 所有有趣的信息都在上面了，但是如果想要使用flag包
	// 最好的方法就是去执行，特别是在main函数（而不是init函数）前执行 02_flag.Parse()
	// 我们这里并不运行，因为它不是个main函数，而且测试单元会详细设计flag内容
}
