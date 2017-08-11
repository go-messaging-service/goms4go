package main

import (
	"encoding/json"
	"fmt"
	"goms4go/src/material"
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
	time.Sleep(time.Millisecond * 100)

	err = client.Register("golang", "news")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	time.Sleep(time.Millisecond * 100)

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
	message := material.NewRegister(material.TypeRegister, topics)

	err := client.sendMessage(message)

	return err
}

func (client *GomsClient) Close() error {
	message := material.NewClose(material.TypeClose)

	err := client.sendMessage(message)

	(*client.connection).Close()

	return err
}

func (client *GomsClient) sendMessage(message interface{}) error {
	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	data = append(data, '\n')
	_, err = (*client.connection).Write(data)

	return err
}
