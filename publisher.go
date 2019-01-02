package PubSub

import (
	"bufio"
)

type publisherApi interface {
	Publish(topic string, message string) error
}

type publisher struct {
	id int
	rw *bufio.ReadWriter
}

var (
	publisherId = 1
)

func newPublisher(rw *bufio.ReadWriter) (publisherApi, error) {
	publisherObj := publisher{
		rw:rw,
	}
	err := sendMessage(publisherObj.rw,serverMsg{
		class:publisherAddedMessage,
	})
	if err != nil{
		return nil,err
}
	message := <- channelMap[publisherAddedMessage]
	publisherObj.id = message.id
	return publisherObj,err
}


func (p  publisher) Publish(topic string, message string) error {
	return sendMessage(p.rw,serverMsg{
		id:p.id,
		class:publisherPublishedMessage,
		topic:topic,
		message:message,
	})
}
