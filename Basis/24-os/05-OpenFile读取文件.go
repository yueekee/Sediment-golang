package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check1(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.OpenFile("my.txt", os.O_RDONLY, os.FileMode(0600))
	check1(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	check1(err)
}
