package router

import (
	"net/http"
	"fmt"
	"github.com/netahe/TinyChat-Server/db"
	"github.com/gorilla/mux"
)

var database db.DB = &db.Chat{}

func InitDB() {
	database.InitChat()
	database.CreateChannel("Welcome")
	database.CreateChannel("Another Channel")
}

func ListChannels(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, database.ListChannels())
}

func GetChannels(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, database.GetChannel(mux.Vars(r)["id"]))
}

func CreateChannel(w http.ResponseWriter, r *http.Request) {
	database.CreateChannel(mux.Vars(r)["id"])
}

func AddMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	database.AddMessage(r.FormValue("author"), r.FormValue("content"), vars["id"])
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, database.GetChannel(mux.Vars(r)["id"]).Messages)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, database.GetChannel(mux.Vars(r)["id"]).Users)

}

func MoveUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["user_id"]
	fromChannel := database.FindUser(userName).Name
	toChannel := vars["chan_id"]

	if fromChannel == toChannel {
		// error: user already in channel
	} else {
		database.MoveUser(userName, fromChannel, toChannel)
	}

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["user_id"]
	toChannel := vars["chan_id"]

	if database.FindUser(userName) != nil {
		// error: user already in a channel
	} else {
		database.AddUser(userName, toChannel)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["user_id"]
	fromChannel := vars["chan_id"]

	if database.FindUser(userName).Name != fromChannel {
		// error: user is not in this channel

	} else {
		database.RemoveUser(userName, fromChannel)
	}
}
