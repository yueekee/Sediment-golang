package main

import (
	"github.com/google/uuid"
	"fmt"
)

func main() {
	data := []byte("hello")
	namedUUID := NamedUUID(data)
	fmt.Println(namedUUID)
}

// Generate a UUID from a byte array.
//
// Generate a Version 3 UUID compatible with Java's
// UUID.namedUUIDFromBytes() method.
//
func NamedUUID(data []byte) *uuid.UUID {
	return uuid.NewMD5(nil, data)
}
