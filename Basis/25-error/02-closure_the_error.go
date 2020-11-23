package main

import (
	"encoding/binary"
	"io"
)

func main() {
	
}

type Point struct {

}

func parse(r io.Reader) (*Point, error) {
	var p Point
	var err error
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	//read()
}
