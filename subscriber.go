package PubSub

import "bufio"

type subscriberApi interface {
	Subscribe(topic string) error
	UnSubscribe(topic string) error
}

type subscriber struct {
	id int
	rw *bufio.ReadWriter
}

func (s * subscriber) Subscribe(topic string) error{
	return nil
}

func (s * subscriber) UnSunscribe(topic string) error{
	return nil
}
