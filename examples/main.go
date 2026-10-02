package main

import "github.com/datalabsgg/gateway"

func main() {
	proxy := gateway.New()

	proxy.Listen(":19132")
}
