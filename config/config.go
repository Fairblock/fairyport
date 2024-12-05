package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	FairyringNodeWS   CosmosNode
	CosmosRelayConfig CosmosRelayConfigType
	EVMRelayTarget    EVMRelayTarget
}

type CosmosRelayConfigType struct {
	DestinationNode Node
	Mnemonic        string
	DerivePath      string
	MetricsPort     uint64
}

type EVMRelayTarget struct {
	ChainRPC        string
	ContractAddress string
}

type CosmosNode struct {
	IP       string
	Port     int64
	Protocol string
}

type Node struct {
	IP            string
	Port          int64
	Protocol      string
	GRPCPort      int64
	AccountPrefix string
	ChainId       string
}

func (c *Config) SetConfig() {
	err := viper.Unmarshal(c)
	if err != nil {
		panic(err)
	}
}

func (c *Config) GetDestinationNodeURI() string {
	nodeURI := c.CosmosRelayConfig.DestinationNode.Protocol + "://" + c.CosmosRelayConfig.DestinationNode.IP + ":" + strconv.FormatInt(c.CosmosRelayConfig.DestinationNode.Port, 10)
	return nodeURI
}

func (c *Config) GetFairyNodeURI() string {
	nodeURI := c.FairyringNodeWS.Protocol + "://" + c.FairyringNodeWS.IP + ":" + strconv.FormatInt(c.FairyringNodeWS.Port, 10)
	return nodeURI
}

func (c *Config) GetMnemonic() string {
	return c.CosmosRelayConfig.Mnemonic
}

func (c *Config) GetGRPCEndPoint() string {
	ep := c.CosmosRelayConfig.DestinationNode.IP + ":" + strconv.FormatInt(c.CosmosRelayConfig.DestinationNode.GRPCPort, 10)
	return ep
}

func (c *Config) ExportConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(homeDir + "/.fairyport"); os.IsNotExist(err) {
		// Directory does not exist, create it
		err = os.MkdirAll(homeDir+"/.fairyport", 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	filePath := filepath.Join(homeDir+"/.fairyport", "config.yml")
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		// File does not exist, create it
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer file.Close()
		log.Printf("Created Config File: %s\n", filePath)
	} else {
		log.Printf("Config file already exists: %s\n", filePath)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(homeDir + "/.fairyport")
	viper.SetConfigType("yml")

	SetInitialConfig(*c)

	return viper.WriteConfigAs(homeDir + "/.fairyport/config.yml")
}

func DefaultConfig() Config {
	var cfg Config
	cfg.CosmosRelayConfig.DestinationNode = defaultNode()
	cfg.FairyringNodeWS = CosmosNode{
		IP:       "127.0.0.1",
		Port:     26657,
		Protocol: "tcp",
	}
	cfg.CosmosRelayConfig.Mnemonic = "# mnemonic"
	cfg.CosmosRelayConfig.DerivePath = "m/44'/118'/0'/0/0"
	cfg.CosmosRelayConfig.MetricsPort = 2224

	cfg.EVMRelayTarget = EVMRelayTarget{
		ChainRPC:        "wss://ws.sketchpad-1.forma.art",
		ContractAddress: "0xcA6cC5c1c4Fc025504273FE61fc0E09100B03D98",
	}
	return cfg
}

func defaultNode() Node {
	var dNode Node
	dNode.IP = "127.0.0.1"
	dNode.GRPCPort = 9090
	dNode.Protocol = "tcp"
	dNode.Port = 26657
	dNode.AccountPrefix = "fairy"
	dNode.ChainId = "fairyring-testnet-3"
	return dNode
}

func SetInitialConfig(c Config) {
	viper.SetDefault("FairyringNodeWS.IP", c.FairyringNodeWS.IP)
	viper.SetDefault("FairyringNodeWS.Port", c.FairyringNodeWS.Port)
	viper.SetDefault("FairyringNodeWS.Protocol", c.FairyringNodeWS.Protocol)

	viper.SetDefault("CosmosRelayConfig.DestinationNode.IP", c.CosmosRelayConfig.DestinationNode.IP)
	viper.SetDefault("CosmosRelayConfig.DestinationNode.Port", c.CosmosRelayConfig.DestinationNode.Port)
	viper.SetDefault("CosmosRelayConfig.DestinationNode.Protocol", c.CosmosRelayConfig.DestinationNode.Protocol)
	viper.SetDefault("CosmosRelayConfig.DestinationNode.GrpcPort", c.CosmosRelayConfig.DestinationNode.GRPCPort)
	viper.SetDefault("CosmosRelayConfig.DestinationNode.AccountPrefix", c.CosmosRelayConfig.DestinationNode.AccountPrefix)
	viper.SetDefault("CosmosRelayConfig.DestinationNode.ChainId", c.CosmosRelayConfig.DestinationNode.ChainId)

	viper.SetDefault("CosmosRelayConfig.Mnemonic", c.CosmosRelayConfig.Mnemonic)
	viper.SetDefault("CosmosRelayConfig.DerivePath", c.CosmosRelayConfig.DerivePath)
	viper.SetDefault("CosmosRelayConfig.MetricsPort", c.CosmosRelayConfig.MetricsPort)

	viper.SetDefault("EVMRelayTarget", c.EVMRelayTarget)
}
