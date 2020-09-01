package main

func main1() {
	defer func() {
		for {
		}
	}()
}

func main2() {
	defer func() { select {} }()
}

func main() {
	defer func() { <-make(chan bool) }()
}
