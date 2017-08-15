package main

import (
	"fmt"

	"github.com/go-messaging-service/goms4go"
)

func main() {
	client, _ := goms4go.Connect("localhost", "55545")

	client.Register(testHandler, "golang", "news")

	select {}
}

func testHandler(data string) {
	fmt.Println("incoming: " + data)
}
