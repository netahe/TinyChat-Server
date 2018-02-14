package messanger

type Message struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

type Channel struct {
	Name     string    `json:name`
	Messages []Message `json:"messages"`
	Users    []string  `json:"users"`
}

type Chat struct {
	Channels []Channel `json:"channels"`
}

func NewChannel(name string) *Channel {
	channel := new(Channel)
	channel.Name = name
	return channel
}

func NewChat() *Chat {
	chat := new(Chat)

	return chat
}

func (chat *Chat) AddChannel(name string) {
	chat.Channels = append(chat.Channels, Channel{Name: name})
}

func (chat *Chat) ListChannels() []string {
	var names []string

	for _, channel := range chat.Channels {

		names = append(names, channel.Name)
	}

	return names
}

func (chat *Chat) GetChannel(name string) *Channel {
	for _, channel := range chat.Channels {
		if channel.Name == name {
			return &channel
		}

	}

	return nil
}

func (channel *Channel) AddMessage(author string, content string) {
	channel.Messages = append(channel.Messages, Message{author, content})
}

func (channel *Channel) AddUser(name string) {
	channel.Users = append(channel.Users, name)
}

func (channel *Channel) RemoveUser(name string) {
	for i, v := range channel.Users {
		if v == name {
			channel.Users = append(channel.Users[:i], channel.Users[i+1])
			break
		}
	}
}

func(chat *Chat) FindUser(name string) channel string {
	return nil
}
