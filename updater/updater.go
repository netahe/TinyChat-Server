package updater

import (
	"github.com/gorilla/websocket"
	"github.com/netahe/TinyChat-Server/db"
)

const (
	JOINED  = iota
	LEFT    = iota
	CREATED = iota
	DELETED = iota
)

type Update interface {
	Encode() []byte
}

type Subscriber interface {
	Subscribe(conn *websocket.Conn)
	Unsubscribe(conn *websocket.Conn)
}

type Publisher interface {
	Publish(update Update)
}

type PubSub interface {
	Publisher
	Subscriber
}

type Updater struct {
	subscribers [] *websocket.Conn
}

func (updater *Updater) Subscribe(conn *websocket.Conn) {
	updater.subscribers = append(updater.subscribers, conn)
}

func (updater *Updater) Unsubscribe(connToRemove *websocket.Conn) {
	for i, conn := range updater.subscribers {
		if conn == connToRemove {
			updater.subscribers = append(updater.subscribers[:i], updater.subscribers[i+1:]...)
		}
	}
}

func (updater *Updater) Publish(update Update) {
	msg := update.Encode()

	for _, conn := range updater.subscribers {
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}

var channelUpdater Updater = Updater{}
var messageUpdaters map[string]Updater
var userUpdaters map[string]Updater

func SubscribeToChannelsUpdater(conn *websocket.Conn) {
	channelUpdater.subscribers = append(channelUpdater.subscribers, conn)
}

func SubscribeToUsersUpdater(chanName string, conn *websocket.Conn) {
	updater := userUpdaters[chanName]
	updater.subscribers = append(updater.subscribers, conn)
}

func SubscribeToMessagesUpdater(chanName string, conn *websocket.Conn) {
	updater := messageUpdaters[chanName]
	updater.subscribers = append(updater.subscribers, conn)
}

func PublishToChannelUpdater(chanName string, action int) {
	for _, conn := range channelUpdater.subscribers {
		conn.WriteMessage(websocket.TextMessage, []byte(chanName))
	}
}

func PublishToUsersUpdater(chanName string, userName string, action int) {
	updater := userUpdaters[chanName]

	for _, conn := range updater.subscribers {
		conn.WriteMessage(websocket.TextMessage, []byte(chanName))
	}
}

func PublishToMessageUpdater(chanName string, msg db.Message) {
	updater := messageUpdaters[chanName]

	for _, conn := range updater.subscribers {
		conn.WriteMessage(websocket.TextMessage, []byte(chanName))
	}
}
