package main

type tvChannel struct {
	ChannelNumber uint
	Channel       chan tvChannelMessage
	Description   string
}

type tvChannelMessage struct {
	ChannelNumber uint
	Message       string
}

type tvChannelsMap struct {
	channels map[uint]tvChannel
}
