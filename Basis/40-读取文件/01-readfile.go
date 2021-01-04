package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//GetFile()
	fileName := "data.txt"
	num, err := GetFloats(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
}

func GetFile() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func GetFloats(fileName string) ([3]float64, error) {
	var num [3]float64
	file, err := os.Open("data.txt")
	if err != nil {
		return num, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return num, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return num, err
	}
	if scanner.Err() != nil {
		return num, scanner.Err()
	}
	return num, nil
}