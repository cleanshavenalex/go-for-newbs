package main

import (
  _ "os"
  "flag"
  _ "bufio"
  "fmt"
  _ "strings"
  _ "encoding/json"
  _ "io/ioutil"
  _ "net/http"
  "strconv"
  _ "sync"
)

var channelNumbers

type TvChannel struct {
  channelNumber
  channel
  description
}

type tvChannelsMap struct {
  channels map[unit]TvChannel
}
var tvChannels tvChannelsMap

var mode

var requestedChannel



func main() {
  mode = flag.String("mode", "sub", "Mode of operation 'pub' or 'sub'(default)")
  requestedChannel  = flag.Int("channel", 2, "Channel to pub or sub to")

  fmt.Println("Requested Channel is %w", requestedChannel)
}
