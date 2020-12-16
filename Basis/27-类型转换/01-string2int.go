package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ID uint64 = 103860

	id := strconv.FormatUint(ID, 10)
	fmt.Println("id:", id)
}
