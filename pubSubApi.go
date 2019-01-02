package PubSub

import (
	"bufio"
)

type pubSubApi interface {
	NewPublisher() publisherApi
	NewSubscriber() subscriberApi
}

type pubSubFactory struct {
	server pubSubServerApi
	rw * bufio.ReadWriter
}

func pubSubServerStart() (pubSubApi,error){
	pubsubobj := pubSubFactory{server:pubSubServer{}}
	err := pubsubobj.server.start()
	if err != nil{
		return nil,err
	}
	return pubsubobj, nil
}

func pubsubApi() (pubSubApi,error){
	InitChannels()
	pubsubObj := pubSubFactory{server:pubSubServer{}}
	rw,err := pubsubObj.server.newRW()
	if err != nil {
		return nil,err
	}
	sendMessage(rw,serverMsg{
		class:addClientMessage,
	})
	go listen(rw)
	go broadcast()
	return pubsubObj,nil
}

func listen(rw *bufio.ReadWriter){
	message,err := receiveMessage(rw)
	if err != nil{
		return
	}
	channelMap[message.class]  <- message
}

func broadcast(){
	for {
		message := <-channelMap[publisherPublishedMessage]
		broadcastMsg(message)
	}
}

func broadcastMsg(msg serverMsg){
	subs,ok := subscriptionMap[msg.topic]
	if ok {
		for sub,flag := range subs{
			if flag {
				sub.callbackMap[msg.topic](msg.message)
			}
		}
	}
}

func (p  pubSubFactory) NewPublisher() publisherApi{
	publisherObj,_:= newPublisher(p.rw)
	return publisherObj
}

func (p  pubSubFactory) NewSubscriber() subscriberApi{
	subscriberObj, _ := newSubscriber(p.rw)
	return subscriberObj
}
