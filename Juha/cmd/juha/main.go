package main

import (
	"flag"
	"log"
)

func main() {
	configPath := flag.String("c", "../config/config.json", "configuration file")

	flag.Parse()
	log.Println("configPath:", *configPath)
	//router := initRouter.SetupRouter()
	//router.Run()
}
