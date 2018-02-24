package router

import (
	"github.com/gorilla/websocket"
	"net/http"
	"github.com/gorilla/mux"
)

func TrackChannels(writer http.ResponseWriter, request *http.Request) {
	channel := mux.Vars(request)["chan_id"]

}

func TrackMessages(writer http.ResponseWriter, request *http.Request) {
	channel := mux.Vars(request)["chan_id"]
}

func TrackUsers(writer http.ResponseWriter, request *http.Request) {
	channel := mux.Vars(request)["chan_id"]
}
