package goms4go

import (
	"bufio"
	"encoding/json"
	"fmt"
	"goms4go/material"
	"net"
	"time"
)

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

func (client *GomsClient) runHandler(handler func(string), topics ...string) {
	go func(handler func(string), conn net.Conn, topics []string) {
		reader := bufio.NewReader(conn)
		for {
			line, err := reader.ReadString('\n')

			if err == nil {
				client.handleLine(handler, line, topics)
			} else {
				fmt.Printf("ERROR OCCURED: %s\n", err.Error())
				time.Sleep(time.Second)
			}
		}
	}(handler, *client.connection, topics)
}

func (client *GomsClient) handleLine(handler func(string), line string, topics []string) {
	rawMessage := &material.AbstractMessage{}
	json.Unmarshal([]byte(line), rawMessage)

	switch rawMessage.Messagetype {
	case material.TypeMessage:
		message := &material.Send{}
		json.Unmarshal([]byte(line), message)

		for _, topic := range topics {
			if contains(message.Topics, topic) {
				handler(message.Data)
			}
		}
	}
}

func (client *GomsClient) Register(handler func(string), topics ...string) error {
	message := material.NewRegister(material.TypeRegister, topics)

	err := client.sendMessage(message)

	if err == nil {
		client.runHandler(handler, topics...)
	}

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
