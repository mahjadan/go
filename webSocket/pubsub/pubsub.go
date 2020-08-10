package pubsub

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type PubSub struct {
	Subscriptions []*Subscription
}

type Client struct {
	Id string
	Conn *websocket.Conn
 }

 const DEFAULT_TOPIC="echo"
 var (
 	PUBLISH = "publish"
 	SUBSCRIBE = "subscribe"
 	UNSUBSCRIBE= "unsubscribe"
 )
type Message struct {
	Action string `json:"action"`
	Topic string `json:"topic"`
	Message json.RawMessage `json:"message"`
}

type Subscription struct {
	Topic string
	Client *Client
}
 func (ps *PubSub)AddClient( c *Client) {
 	fmt.Println("adding client ",c.Id)
 	sub:=&Subscription{
 		Topic: "echo",
 		Client: c,
	}
	 ps.Subscriptions = append(ps.Subscriptions, sub)
 }

func (ps *PubSub) HandleMessage(c Client, messageType int,  msg []byte) {
	m:= Message{}
	fmt.Printf("%s",msg)
	err:= json.Unmarshal(msg,&m)
	if err!=nil {
		fmt.Println("error Incorrect message:",err)
		return
	}
	fmt.Println("got message: ",m.Action,m.Topic,string(m.Message))

	switch m.Action {
	case PUBLISH:
		ps.Publish(m.Message,messageType, m.Topic ,c)
	case SUBSCRIBE:
		ps.Subscribe(c,m.Topic)
	case UNSUBSCRIBE:

		ps.UnSubscribe(&c,m.Topic)
	default:
		fmt.Println("unknown action...")
	}
}

func (ps *PubSub) Subscribe(c Client, topic string){
	ps.UnSubscribe(&c,DEFAULT_TOPIC)
	if !ps.AlreadySubscribed(c,topic) {
		fmt.Printf("subscribing %s to %s.\n",c.Id, topic	)
		s := &Subscription{
			Topic:  topic,
			Client: &c,
		}
		ps.Subscriptions = append(ps.Subscriptions, s)
	}
}


func (ps *PubSub) Publish(message json.RawMessage, messageType int, topic string, c Client) {
	for _, subscription := range ps.Subscriptions {
		//send to all with same topic except the sender.
		if subscription.Topic == topic && subscription.Client.Id != c.Id{
			fmt.Println("publishing to :", subscription.Client.Id)
			subscription.Client.Conn.WriteMessage(messageType,message)
		}
	}
	fmt.Println("subscriptions are :",ps.Subscriptions)
}

func (ps *PubSub) UnSubscribe(c *Client, topic string){
	for i,sub:= range ps.Subscriptions{
		if sub.Topic == topic && sub.Client.Id == c.Id {
			fmt.Printf("Unsubscribing %s from : %s\n", sub.Client.Id,sub.Topic)
			ps.Subscriptions = append(ps.Subscriptions[:i], ps.Subscriptions[i+1:]...)
			return
		}
	}
}

func (s Subscription)String() string{
	return "topid:" + s.Topic +"," + "Id:" +s.Client.Id
}

func (ps *PubSub) AlreadySubscribed(c Client,topic string) bool{
	for _,sub := range ps.Subscriptions{
		if sub.Client.Id == c.Id && sub.Topic == topic{
			return true
		}
	}
	return false
}

func (ps *PubSub) RemoveClient(id string) {
	for i,sub:= range ps.Subscriptions{
		if sub.Client.Id == id {
			fmt.Printf("Removing client id: %s\n",id)
			ps.Subscriptions = append(ps.Subscriptions[:i], ps.Subscriptions[i+1:]...)
			return
		}
	}
}
