package main

import "github.com/mclargo/go-framework/cmd/bootstrap"

// TODO: main starts a server, but could be good to add another sample of cli
func main() {
	if err := bootstrap.RunServer(); err != nil {
		panic(err)
	}
}
