package config

import (
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	FairyRingNode   Node
	DestinationNode Node
	Mnemonic        string
}

type Node struct {
	IP       string
	Port     int64
	Protocol string
	GRPCPort int64
}

func (c *Config) SetConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	//viper.SetDefault("")

	err := viper.Unmarshal(c)
	if err != nil {
		panic(err)
	}
}

func (c *Config) GetDestinationNodeURI() string {
	nodeURI := c.DestinationNode.Protocol + "://" + c.DestinationNode.IP + ":" + strconv.FormatInt(c.DestinationNode.Port, 10)
	return nodeURI
}

func (c *Config) GetFairyNodeURI() string {
	nodeURI := c.FairyRingNode.Protocol + "://" + c.FairyRingNode.IP + ":" + strconv.FormatInt(c.FairyRingNode.Port, 10)
	return nodeURI
}

func (c *Config) GetMnemonic() string {
	return c.Mnemonic
}

func (c *Config) GetGRPCEndPoint() string {
	ep := c.DestinationNode.IP + ":" + strconv.FormatInt(c.DestinationNode.GRPCPort, 10)
	return ep
}
