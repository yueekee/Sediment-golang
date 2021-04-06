package main

import (
	"flag"
	"fmt"
	"log"
)

var name string

func main() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "go命令", "help")

	javaCmd := flag.NewFlagSet("java", flag.ExitOnError)
	javaCmd.StringVar(&name, "name2", "java命令", "help")

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	fmt.Printf("%#v\n", args)	// []string{"go", "-name=go111"}
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "java":
		_ = javaCmd.Parse(args[1:])
	}
	log.Printf("name: %s", name)
}

/*
go run 02-demo2.go go -name=go111
*/
