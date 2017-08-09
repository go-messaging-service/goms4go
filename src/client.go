package main

import (
	"fmt"
	"net"
)

func main() {
	client, err := Connect("localhost", "55545")
	if err == nil {
		client.Close()
	} else {
		fmt.Println(err.Error())
	}
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

func (client *GomsClient) Register(topics ...string) {

}

func (client *GomsClient) Close() {
	//TOOD send logout message

	client.connection.Close()
}
