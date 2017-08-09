package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	client, err := Connect("localhost", "55545")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	time.Sleep(time.Second)

	err = client.Register("golang", "news")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	time.Sleep(time.Second)

	client.Close()
}

type GomsClient struct {
	connection *net.Conn
}

func Connect(address string, port string) (*GomsClient, error) {
	connection, err := net.Dial("tcp", address+":"+port)
	if err != nil {
		return nil, err
	}

	result := GomsClient{
		connection: &connection,
	}

	return &result, nil
}

func (client *GomsClient) Register(topics ...string) error {
	message := getRegisterMessage(topics)
	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	data = append(data, '\n')
	_, err = (*client.connection).Write(data)

	return err
}

func (client *GomsClient) Close() {
	//TOOD send logout message

	(*client.connection).Close()
}
