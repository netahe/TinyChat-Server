package router

import (
	"fmt"
	mux "github.com/gorilla/mux"
	"net/http"
)

func InitServer() {
	r := mux.NewRouter()
	InitDB()

	r.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "<html><head><title>Golang Messager</title></head><body><h1>Hello chat!</h1></body></html>")
		})

	r.HandleFunc("/channels", ListChannels).Methods("GET")

	r.HandleFunc("/channels/{id}", GetChannels).Methods("GET")

	r.HandleFunc("/channels/{id}", CreateChannel).Methods("POST")

	r.HandleFunc("/channels/{id}/messages", AddMessage).Methods("POST")

	r.HandleFunc("/channels/{id}/messages", GetMessages).Methods("GET")

	r.HandleFunc("/channels/{id}/users", GetUsers).Methods("GET")

	r.HandleFunc("/channels/{chan_id}/users/{user_id}", MoveUser).Methods("PUT")

	r.HandleFunc("/channels/{chan_id}/users/{user_id}", AddUser).Methods("POST")

	r.HandleFunc("/channels/{chan_id}/users/{user_id}", DeleteUser).Methods("DELETE")

	r.HandleFunc("/updater/channels", TrackChannels)

	r.HandleFunc("/updater/channels/{chan_id}/messages", TrackMessages)

	r.HandleFunc("/updater/channels/{chan_id}/users", TrackUsers)

	http.ListenAndServe(":8000", r)
}
