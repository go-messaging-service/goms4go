package main

import (
	"net"
)

func main() {
	client, err := Connect("localhost", "55545")

}

type GomsClient struct {
	connection net.Conn
}

func Connect(address string, port string) (*GomsClient, error) {
	connection, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		return nil, err
	}

	result := GomsClient{
		connection: connection,
	}

	return &result, nil
}
