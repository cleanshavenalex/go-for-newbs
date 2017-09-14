package main

import (
	_ "bufio"
	"flag"
	"fmt"
	_ "net/http"
	"os"
	_ "sync"
)

var tvMode string // 'set' Tv Set channel 'watcher', 'station' Tv Station publishes to channels
var tvChannels tvChannelsMap
var tvActiveChannel int
var tvChannelNumbers = [12]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

func main() {
	tvMode := flag.String("mode", "set", "Set - watching a channel or Station - publishing content to channels")
	tvActiveChannel := flag.Int("active-channel", 2, "Active Channel to watch or publish content to")
	tvStationApiKey := flag.String("tv-station-key", "noaccess", "Tv Station API Key")
	channelMessage := flag.String("tv-station-message", "Hello Universe!", "Tv Channel Message")
	requiredTvStationKey := "admin"
	flag.Parse()

	tvChannels = tvChannelsMap{
		channels: make(map[uint]tvChannel),
	}

	for _, r := range tvChannelNumbers {
		newChan := make(chan tvChannelMessage, 1)
		channelDescription := fmt.Sprintf("Channel %v", r)
		// fmt.Printf("channelDescription for channel %v, %v\n", r, channelDescription)
		tvChannels.channels[uint(r)] = tvChannel{Channel: newChan, ChannelNumber: uint(r), Description: channelDescription}

	}

	fmt.Println("tvMode:", *tvMode)
	fmt.Println("tvActiveChannel:", *tvActiveChannel)
	// this will never actually work because two instances of the program would have their own channels
	// would need to implement a network service that has the channels and sends/receives channel messags over the network
	// not practical just a way to work in GO syntax
	if *tvMode == "station" {
		if *tvStationApiKey == requiredTvStationKey {
			fmt.Println("tv-station-key is correct, running in station mode...")
			channelToControl := tvChannels.channels[uint(*tvActiveChannel)].Channel
			channelToControl <- tvChannelMessage{ChannelNumber: uint(*tvActiveChannel), Message: *channelMessage}
		} else {
			fmt.Println("You must pass in the correct tv-station-key to run this program in station mode")
			os.Exit(1)
		}
	} else {
		fmt.Println("Running in Tv Set mode...")
		channelToWatch := tvChannels.channels[uint(*tvActiveChannel)].Channel
		for elem := range channelToWatch {
			fmt.Println(elem)
		}
	}

}
