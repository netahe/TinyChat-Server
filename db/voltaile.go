package db

type Chat struct {
	Channels []Channel
}

func (channel *Channel) addMessage(author string, content string) {
	channel.Messages = append(channel.Messages, Message{author, content})
}

func (channel *Channel) addUser(userName string) {
	channel.Users = append(channel.Users, userName)
}

func (channel *Channel) removeUser(userName string) {
	for i, v := range channel.Users {
		if v == userName {
			channel.Users = append(channel.Users[:i], channel.Users[i+1])
			break
		}
	}
}

func (chat *Chat) InitChat() {}

func (chat *Chat) CreateChannel(chanName string) {
	chat.Channels = append(chat.Channels, Channel{Name: chanName})
}

func (chat *Chat) ListChannels() []string {
	var names []string

	for _, channel := range chat.Channels {

		names = append(names, channel.Name)
	}

	return names
}

func (chat *Chat) GetChannel(chanName string) *Channel {
	for _, channel := range chat.Channels {
		if channel.Name == chanName {
			return &channel
		}
	}
	return &Channel{}
}

func (chat *Chat) AddUser(userName string, chanName string) {
	chat.GetChannel(chanName).addUser(userName)
}

func (chat *Chat) RemoveUser(userName string, chanName string) {
	chat.GetChannel(chanName).removeUser(userName)
}

func (chat *Chat) MoveUser(userName string, fromChan string, toChan string) {
	chat.GetChannel(fromChan).removeUser(userName)
	chat.GetChannel(toChan).addUser(userName)
}

func (chat *Chat) AddMessage(author string, content string, chanName string) {
	chat.GetChannel(chanName).addMessage(author, content)
}

func (chat *Chat) FindUser(userName string) *Channel {
	for _, channel := range chat.Channels {
		for _, user := range channel.Users {
			if user == userName {
				return &channel
			}
		}
	}

	return nil
}