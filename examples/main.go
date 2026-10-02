package main

import "github.com/datalabsgg/gateway"

func main() {
	proxy := gateway.New()

	proxy.beforeJoin =
	proxy.onJoin =

	proxy.Listen(":19132")
}
