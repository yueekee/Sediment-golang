package main

import "fmt"

type Error interface {
	Error() string
}

type RPCError struct {
	Code 		int
	Message 	string
}

func (e RPCError) Error() string {
	return fmt.Sprintf("mesg: %s, code: %d", e.Message, e.Code)
}

func main() {
	e := RPCError{
		Code:400,
		Message:"unknown err",
	}
	s := e.Error()
	fmt.Println(s)
}
