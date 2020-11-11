package typhoon

import (
	"encoding/json"
	"io/ioutil"
)

type BufferConfig struct {
	HandshakeAddress int `json:"handshake_address"`
	PlayerName       int `json:"player_name"`
	ChatMessage      int `json:"chat_message"`
}

type Config struct {
	ListenAddress string       `json:"listen_address"`
	MaxPlayers    int          `json:"max_players"`
	Motd          string       `json:"motd"`
	Restricted    bool         `json:"restricted"`
	Logs          bool         `json:"logs"`
	Compression   bool         `json:"enable_compression"`
	Threshold     int          `json:"compression_threshold"`
	BufferConfig  BufferConfig `json:"buffer_config"`
}

var (
	config  Config
	favicon string
)

func initConfig() (err error) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}
	return
}

func (c *Core) GetConfig(config interface{}) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, config); err != nil {
		panic(err)
	}
}
