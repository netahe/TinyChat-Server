package db

type VoltaileMessage struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

type VoltaileChannel struct {
	Name     string            `json:name`
	Messages []VoltaileMessage `json:"messages"`
	Users    []string          `json:"users"`
}

type VoltaileChat struct {
	Channels []VoltaileChannel `json:"channels"`
}

func NewChannel(name string) *VoltaileChannel {
	channel := new(VoltaileChannel)
	channel.Name = name
	return channel
}

func NewChat() *VoltaileChat {
	chat := new(VoltaileChat)

	return chat
}

func (chat *VoltaileChat) AddChannel(name string) {
	chat.Channels = append(chat.Channels, VoltaileChannel{Name: name})
}

func (chat *VoltaileChat) ListChannels() []string {
	var names []string

	for _, channel := range chat.Channels {

		names = append(names, channel.Name)
	}

	return names
}

func (chat *VoltaileChat) GetChannel(name string) *VoltaileChannel {
	for _, channel := range chat.Channels {
		if channel.Name == name {
			return &channel
		}

	}

	return nil
}

func (channel *VoltaileChannel) AddMessage(author string, content string) {
	channel.Messages = append(channel.Messages, VoltaileMessage{author, content})
}

func (channel *VoltaileChannel) AddUser(name string) {
	channel.Users = append(channel.Users, name)
}

func (channel *VoltaileChannel) RemoveUser(name string) {
	for i, v := range channel.Users {
		if v == name {
			channel.Users = append(channel.Users[:i], channel.Users[i+1])
			break
		}
	}
}

func (chat *VoltaileChat) FindUser(name string) string {
	for _, channel := range chat.Channels {
		for _, user := range channel.Users {
			if user == name {
				return channel.Name
			}
		}
	}

	return ""
}
