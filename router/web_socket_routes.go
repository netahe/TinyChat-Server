package router

import (
	"github.com/gorilla/websocket"
	"net/http"
	"log"
	"github.com/netahe/TinyChat-Server/updater"
	"github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func TrackChannels(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		log.Println(err)
		return
	}

	updater.SubscribeToChannelsUpdater(conn)

}

func TrackMessages(writer http.ResponseWriter, request *http.Request) {
	channel := mux.Vars(request)["chan_id"]
	conn, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		log.Println(err)
		return
	}

	updater.SubscribeToMessagesUpdater(channel, conn)

}

func TrackUsers(writer http.ResponseWriter, request *http.Request) {
	channel := mux.Vars(request)["chan_id"]
	conn, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		log.Println(err)
		return
	}

	updater.SubscribeToUsersUpdater(channel, conn)
}
