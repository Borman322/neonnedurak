package main

import (
	"fmt"

	"./habr"
	telega "./send"
)

func main() {

	f := habr.HabrGo()
	telega.SendMessage(f)
	fmt.Println(f)
}
