package main

import (
	"encoding/json"
	"goms4go/src/material"
	"net"
	"time"
)

func main() {
	client, _ := Connect("localhost", "55545")
	time.Sleep(time.Millisecond * 100)

	client.Register("golang", "news")
	time.Sleep(time.Millisecond * 100)

	client.Send("test data :)", "golang", "news")
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

func (client *GomsClient) Send(data string, topics ...string) error {
	message := material.NewSend(material.TypeSend, topics, data)

	err := client.sendMessage(message)

	return err
}

func (client *GomsClient) Logout(topics ...string) error {
	message := material.NewLogout(material.TypeLogout, topics)

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
