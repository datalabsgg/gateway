package main

import (
	"fmt"

	"github.com/datalabsgg/gateway"
)

func main() {
	proxy := gateway.New()

	proxy.BeforeJoin = func(string) {
		fmt.Println("test")
	}

	proxy.OnJoin = func(string) {
		fmt.Println("test")
	}

	//proxy.Listen(":19132")
}
