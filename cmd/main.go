package main

import (
	"fmt"

	"github.com/icrxz/grpc-go/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
