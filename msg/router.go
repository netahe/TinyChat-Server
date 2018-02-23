package messanger

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func InitServer() {
	r := mux.NewRouter()

	chat := NewChat()
	chat.AddChannel("Welcome")
	chat.AddChannel("Another Channel")

	r.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "<html><head><title>Golang Messager</title></head><body><h1>Hello chat!</h1></body></html>")
		})

	r.HandleFunc("/channels",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, chat.ListChannels())
		})

	r.HandleFunc("/channels/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, chat.GetChannel(mux.Vars(r)["id"]))
		}).Methods("GET")

	r.HandleFunc("/channels/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			chat.AddChannel(mux.Vars(r)["id"])
		}).Methods("POST")

	r.HandleFunc("/channels/{id}/messages",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			chat.GetChannel(vars["id"]).AddMessage(r.FormValue("author"), r.FormValue("content"))
		}).Methods("POST")

	r.HandleFunc("/channels/{id}/messages",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, chat.GetChannel(mux.Vars(r)["id"]).Messages)

		}).Methods("GET")

	r.HandleFunc("/channels/{id}/users",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, chat.GetChannel(mux.Vars(r)["id"]).Users)

		}).Methods("GET")

	r.HandleFunc("/channels/{chan_id}/users/{user_id}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			if chat.FindUser(vars["user_id"]) != vars["chan_id"] {
				// error: user already in channel
			} else {
				chat.GetChannel(vars["chan_id"]).RemoveUser(chat.FindUser(vars["user_id"])) // remove user from current channel before adding them to new channel
				chat.GetChannel(vars["chan_id"]).AddUser(vars["user_id"])
			}

		}).Methods("PUT")

	r.HandleFunc("/channels/{chan_id}/users/{user_id}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)

			if chat.FindUser(vars["user_id"]) == "" {
				// error: user is not in this channel

			} else if chat.FindUser(vars["user_id"]) != vars["chan_id"] {
				// error: user is in another channel

			} else {
				chat.GetChannel(vars["chan_id"]).RemoveUser(vars["user_id"])
			}
		}).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
