package PubSub

import (
	"bufio"
	"errors"
	"fmt"
)

type subscriberApi interface {
	Subscribe(topic string, callback Func) error
	UnSubscribe(topic string) error
}

type subscriber struct {
	id int
	rw *bufio.ReadWriter
	callbackMap map[string]Func
}

type Func func(string);

var(
	subscriberId = 1
	subscriptionMap  = make(map[string]map[*subscriber]bool)
)

func newSubscriber(rw * bufio.ReadWriter) (subscriberApi,error) {
	subscriberObj := &subscriber{rw:rw}
	err := sendMessage(rw,serverMsg{
		class:subscriberAddedMessage,
	})
	if err != nil{
		return nil, err
	}
	message := <-channelMap[subscriberAddedMessage]
	subscriberObj.id = message.id
	return subscriberObj,nil
}

func (s  * subscriber) Subscribe(topic string,callback Func) error{
	s.callbackMap[topic] = callback
	_,ok := subscriptionMap[topic]
	if !ok{
		subscriptionMap[topic] = make(map[*subscriber]bool)
	}
	if _,ok := subscriptionMap[topic]; ok {
		fmt.Println("Subscription already exists!!")
		return errors.New("subscription already exist")
	}
	subscriptionMap[topic][s] = true
	return nil
}

func (s  * subscriber) UnSubscribe(topic string) error{
	subscriberMap,ok := subscriptionMap[topic]
	if ok {
			_,ok := subscriberMap[s]
			if ok {
				delete(subscriberMap,s)
			}
	}
	delete(s.callbackMap,topic)
	return nil
}
