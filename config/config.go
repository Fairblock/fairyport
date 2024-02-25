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
	FairyRingNode   Node
	DestinationNode Node
	Mnemonic        string
	DerivePath      string
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

func (c *Config) ExportConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(homeDir)
	fmt.Println(c)

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
		log.Println("creating initial config")

		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer file.Close()
		log.Printf("Created file: %s\n", filePath)
	} else {
		log.Printf("Config file already exists: %s\n", filePath)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(homeDir + "/.fairyport")
	viper.SetConfigType("yml")

	setInitialConfig(*c)

	err = viper.WriteConfigAs(homeDir + "/.fairyport/config.yml")
	log.Println(err)

	return nil
}

func DefaultConfig() Config {
	var cfg Config
	cfg.DestinationNode = defaultNode()
	cfg.FairyRingNode = defaultNode()
	cfg.Mnemonic = "# mnemonic"
	cfg.DerivePath = "m/44'/118'/0'/0/0"
	return cfg
}

func defaultNode() Node {
	var dNode Node
	dNode.IP = "127.0.0.1"
	dNode.GRPCPort = 9090
	dNode.Protocol = "rpc"
	dNode.Port = 26657
	dNode.AccountPrefix = "fairy"
	dNode.ChainId = "fairytest-3"
	return dNode
}

func setInitialConfig(c Config) {
	viper.SetDefault("FairyRingNode.ip", c.FairyRingNode.IP)
	viper.SetDefault("FairyRingNode.port", c.FairyRingNode.Port)
	viper.SetDefault("FairyRingNode.protocol", c.FairyRingNode.Protocol)
	viper.SetDefault("FairyRingNode.grpcport", c.FairyRingNode.GRPCPort)
	viper.SetDefault("FairyRingNode.accountPrefix", c.FairyRingNode.AccountPrefix)
	viper.SetDefault("FairyRingNode.chainId", c.FairyRingNode.ChainId)

	viper.SetDefault("DestinationNode.ip", c.DestinationNode.IP)
	viper.SetDefault("DestinationNode.port", c.DestinationNode.Port)
	viper.SetDefault("DestinationNode.protocol", c.DestinationNode.Protocol)
	viper.SetDefault("DestinationNode.grpcport", c.DestinationNode.GRPCPort)
	viper.SetDefault("DestinationNode.accountPrefix", c.DestinationNode.AccountPrefix)
	viper.SetDefault("DestinationNode.chainId", c.DestinationNode.ChainId)

	viper.SetDefault("Mnemonic", c.Mnemonic)
	viper.SetDefault("derivePath", c.DerivePath)
}
