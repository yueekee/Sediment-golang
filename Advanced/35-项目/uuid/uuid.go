package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	// Creating UUID Version 4
	// panic on error
	u1 := uuid.NewV1()
	fmt.Printf("UUIDv1: %s\n", u1)

	b := byte('l')
	u2 := uuid.NewV2(b)
	fmt.Printf("UUIDv2: %s\n", u2)

	u3 := uuid.NewV3(u2, "123")
	fmt.Printf("UUIDv3: %s\n", u3)

	u4 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u4)

	// Parsing UUID from string input
	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}

//func GetUUID() (string, error) {
//	u2, err := uuid.NewV4()
//	if err != nil {
//		fmt.Printf("Something went wrong: %s", err)
//		return "", err
//	}
//	return u2.String(), nil
//}
