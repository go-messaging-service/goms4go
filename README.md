# gom4go
A go library for goMS.
# Installation
Just use the `go get` command:
```bash
go get -u github.com/go-messaging-service/goms4go
```
# Usage
Here's an simple example receiving incoming messages for the topics `golang` and `news`
```go
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

```
