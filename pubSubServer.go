package PubSub

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
)

var(
	PORT = ":3000"
	streams []*bufio.ReadWriter = make([]*bufio.ReadWriter,10)
)

type pubSubServerApi interface {
	newRW() (*bufio.ReadWriter, error)
	start() error
}

type pubSubServer struct {
}

type serverMsg struct{
	id int
	class string
	topic string
	message string
}

func (p pubSubServer) newRW()(*bufio.ReadWriter,error){
	conn,err := net.Dial("tcp","localhost"+PORT)
	if err != nil{
		return nil,err
	}
	rw := bufio.NewReadWriter(bufio.NewReader(conn),bufio.NewWriter(conn))
	return rw, nil
}

func (p  pubSubServer) start() error {
	listen,err := net.Listen("tcp",PORT)
	if err != nil {
		return err
	}
	for{
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connections")
			continue
		}
		fmt.Println("Connection accepted")
		go handleConnections(conn)
	}
}

func sendMessage(b *bufio.ReadWriter, message serverMsg)error{
	err := gob.NewEncoder(b).Encode(&message)
	if err != nil {
		return err
	}
	b.Flush()
	return nil
}

func receiveMessage(b *bufio.ReadWriter) (message serverMsg,err error){
	err = gob.NewDecoder(b).Decode(&message)
	if err != nil {
		return message,err
	}
	return message, nil
}

func handleConnections(conn net.Conn){
	defer conn.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(conn),bufio.NewWriter(conn))
	for{
		msg,err := receiveMessage(rw)
		if err != nil{
			fmt.Println("Error receiving message")
			return
		}
		handleReceivedMessages(msg,rw)
	}
}

func handleReceivedMessages(message serverMsg,rw *bufio.ReadWriter){
	switch message.class {
	case addClientMessage:
		streams = append(streams, rw)
	case publisherAddedMessage:
		message.id = publisherId
		publisherId++
		sendMessage(rw,message)
	case publisherPublishedMessage:
		for _,stream := range streams{
			sendMessage(stream,message)
		}
	case subscriberAddedMessage:
		message.id = subscriberId
		subscriberId++
		sendMessage(rw,message)
	case subscriberSubscribedMessage:
		return
	case subscriberUnSubscribedMesssage:
		return
	default:
		fmt.Println("unhandled case has been encountered")
	}
}
