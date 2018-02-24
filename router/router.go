package router

import (
	"fmt"
	mux "github.com/gorilla/mux"
	"net/http"
	"github.com/netahe/TinyChat-Server/db"

)

func InitServer() {
	r := mux.NewRouter()
	var db db.DB = &db.Chat{}
	db.InitChat()

	db.CreateChannel("Welcome")
	db.CreateChannel("Another Channel")

	r.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "<html><head><title>Golang Messager</title></head><body><h1>Hello chat!</h1></body></html>")
		})

	r.HandleFunc("/channels",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, db.ListChannels())
		})

	r.HandleFunc("/channels/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, db.GetChannel(mux.Vars(r)["id"]))
		}).Methods("GET")

	r.HandleFunc("/channels/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			db.CreateChannel(mux.Vars(r)["id"])
		}).Methods("POST")

	r.HandleFunc("/channels/{id}/messages",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			db.AddMessage(r.FormValue("author"), r.FormValue("content"), vars["id"])
		}).Methods("POST")

	r.HandleFunc("/channels/{id}/messages",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, db.GetChannel(mux.Vars(r)["id"]).Messages)

		}).Methods("GET")

	r.HandleFunc("/channels/{id}/users",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, db.GetChannel(mux.Vars(r)["id"]).Users)

		}).Methods("GET")

	r.HandleFunc("/channels/{chan_id}/users/{user_id}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			userName := vars["user_id"]
			fromChannel := db.FindUser(userName).Name
			toChannel := vars["chan_id"]

			if fromChannel == toChannel {
				// error: user already in channel
			} else {
				db.MoveUser(userName, fromChannel, toChannel)
			}

		}).Methods("PUT")

	r.HandleFunc("/channels/{chan_id}/users/{user_id}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			userName := vars["user_id"]
			toChannel := vars["chan_id"]

			if db.FindUser(userName) != nil {
				// error: user already in a channel
			} else {
				db.AddUser(userName, toChannel)
			}

		}).Methods("POST")

	r.HandleFunc("/channels/{chan_id}/users/{user_id}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			userName := vars["user_id"]
			fromChannel := vars["chan_id"]

			if db.FindUser(userName).Name != fromChannel {
				// error: user is not in this channel

			} else {
				db.RemoveUser(userName, fromChannel)
			}
		}).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
