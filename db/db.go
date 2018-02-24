package db

/*database agnostic interface for db communication */

type Channel struct {
	Name     string
	Users    []string
	Messages []Message
}

type Message struct {
	Author  string
	Content string
}

type DB interface {
	InitChat()
	CreateChannel(chanName string)
	GetChannel(chanName string) *Channel
	ListChannels() [] string
	FindUser(userName string) *Channel
	AddUser(userName string, chanName string)
	RemoveUser(userName string, chanName string)
	MoveUser(userName string, fromChan string, toChan string)
	AddMessage(author string, content string, chanName string)
}
