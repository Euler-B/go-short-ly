package main

import (
	"github.com/euler-b/go-short-ly/model"
	"github.com/euler-b/go-short-ly/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
