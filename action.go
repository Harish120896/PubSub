package PubSub

const (
	addClientMessage = "Add Client"
	publisherAddedMessage = "Publisher Added"
	publisherPublishedMessage = "Publisher Published"
	subscriberAddedMessage = "Subscriber Added"
	subscriberSubscribedMessage = "Subscriber Subscribed"
	subscriberUnSubscribedMesssage = "Subscriber Unsubscribed"
	chaLen = 1
)

var (
	channelMap map[string]chan serverMsg = make(map[string]chan serverMsg)
)

func InitChannels(){
	channelMap[publisherAddedMessage]chan = make(chan serverMsg,chaLen)
	channelMap[publisherPublishedMessage]chan = make(chan serverMsg, chaLen)
	channelMap[subscriberAddedMessage]chan = make(chan serverMsg, chaLen)
	channelMap[subscriberSubscribedMessage]chan = make(chan serverMsg, chaLen)
	channelMap[subscriberUnSubscribedMesssage]chan = make(chan serverMsg,chaLen)
}
