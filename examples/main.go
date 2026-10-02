package main

import (
	"fmt"

	"github.com/datalabsgg/gateway"
)

func main() {
	proxy := gateway.New( /* gateway.Config{} */ )

	proxy.BeforeJoin = func(string) {
		fmt.Println("test")
	}

	proxy.OnJoin = func(string) {
		fmt.Println("test")
	}

	//proxy.Listen(":19132")
}
