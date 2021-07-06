package main

import "fmt"

type BodyPart int

const (
	Head BodyPart = iota // Head = 0
	Shoulder // Shoulder = 1
	Knee // Knee = 2
	Toe // Toe = 3
)

func (bp BodyPart) String(i BodyPart) string {
	s := []string{"Head", "Shoulder", "Knee", "Toe"}
	return s[i]
}

func main() {
	var bp BodyPart
	s := bp.String(Head)
	fmt.Println(s)
}
