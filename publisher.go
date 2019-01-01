package PubSub

import (
	"bufio"
)

type publisherApi interface {
	Publish(topic string, message serverMsg) error
}

type publisher struct {
	id int
	rw *bufio.ReadWriter
}

func (p * publisher) Publish(topice string, message serverMsg){

}
