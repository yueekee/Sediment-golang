package main

import (
	"flag"
	"fmt"
)

var inputName = flag.String("name", "CHENJIAN", "Input Your Name.")
var inputAge = flag.Int("age", 27, "Input Your Age")
var inputGender = flag.String("gender", "female", "Input Your Gender")
var inputFlagvar int

func Init() {
	flag.IntVar(&inputFlagvar, "flagname", 1234, "Help")
}
func main() {
	flag.String("go run 01", "go run ../01*", "zhe shi " )
	Init()
	flag.Parse()
	// func Args() []string
	// Args returns the non-02_flag command-line arguments.
	// func NArg() int
	// NArg is the number of arguments remaining after flags have been processed.
	fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}
	fmt.Println("name=", *inputName)
	fmt.Println("age=", *inputAge)
	fmt.Println("gender=", *inputGender)
	fmt.Println("flagname=", inputFlagvar)

	a :=  []int{49286,49285,49284,46185,46184,46183,46188,46187,46186,46182,
		46181,46177,46176,45505,45504,45503,45502,45501,45492,45491,45490,45495,45494,
		45493,45489,45488,45487,43410,43413,43412,43417,43409,43416,43414,43418,43411,43408,43407}
	fmt.Println("1111", len(a))

}