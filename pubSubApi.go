package PubSub

import "bufio"

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

func pubsubApi(){

}

func (p  pubSubFactory) NewPublisher() publisherApi{
	return nil
}

func (p  pubSubFactory) NewSubscriber() subscriberApi{
	return nil
}
