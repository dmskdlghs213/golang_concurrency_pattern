package main

import (
	"fmt"

	"github.com/dmskdlghs213/golang_concurrency_pattern/conccurrency"
	"github.com/labstack/gommon/log"
)

func main() {
	fmt.Println("main")

	msg, err := conccurrency.ConccurrencyCall()
	if err != nil {
		log.Error(err)
	}

	fmt.Println(msg)
}
