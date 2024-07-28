package main

import "go-gin-sarulabs/src/router"

func main() {
	err := router.SetRouter()
	if err != nil {
		panic(err)
	}
}
