package main

import "github.com/mclargo/go-framework/cmd/bootstrap"

func main() {
	if err := bootstrap.RunServer(); err != nil {
		panic(err)
	}
}
