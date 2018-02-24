package db

/* basic data structures for the chat*/

type Channel struct {
	Name     string
	Users    []string
	Messages []Message
}

type Message struct {
	Author  string
	Content string
}
