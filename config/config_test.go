package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Fairblock/fairyport/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	cfg := config.DefaultConfig()

	assert.Equal(t, "127.0.0.1", cfg.CosmosRelayConfig.DestinationNode.IP)
	assert.Equal(t, int64(9090), cfg.CosmosRelayConfig.DestinationNode.GRPCPort)
	assert.Equal(t, "tcp", cfg.CosmosRelayConfig.DestinationNode.Protocol)
	assert.Equal(t, int64(26657), cfg.FairyringNodeWS.Port)
	assert.Equal(t, "tcp", cfg.FairyringNodeWS.Protocol)
	assert.Equal(t, "m/44'/118'/0'/0/0", cfg.CosmosRelayConfig.DerivePath)
	assert.Equal(t, "wss://ws.sketchpad-1.forma.art", cfg.EVMRelayTarget.ChainRPC)
}

func TestSetConfig(t *testing.T) {
	cfg := config.DefaultConfig()
	cfg.CosmosRelayConfig.Mnemonic = "test mnemonic"
	viper.Set("CosmosRelayConfig.Mnemonic", "overridden mnemonic")

	cfg.SetConfig()
	assert.Equal(t, "overridden mnemonic", cfg.GetMnemonic())
}

func TestGetDestinationNodeURI(t *testing.T) {
	cfg := config.DefaultConfig()
	expectedURI := "tcp://127.0.0.1:26657"
	assert.Equal(t, expectedURI, cfg.GetDestinationNodeURI())
}

func TestGetFairyNodeURI(t *testing.T) {
	cfg := config.DefaultConfig()
	expectedURI := "tcp://127.0.0.1:26657"
	assert.Equal(t, expectedURI, cfg.GetFairyNodeURI())
}

func TestGetMnemonic(t *testing.T) {
	cfg := config.DefaultConfig()
	assert.Equal(t, "# mnemonic", cfg.GetMnemonic())
}

func TestGetGRPCEndPoint(t *testing.T) {
	cfg := config.DefaultConfig()
	expectedEndpoint := "127.0.0.1:9090"
	assert.Equal(t, expectedEndpoint, cfg.GetGRPCEndPoint())
}

func TestExportConfig(t *testing.T) {
	cfg := config.DefaultConfig()

	// Mock the home directory
	homeDir := t.TempDir()
	os.Setenv("HOME", homeDir)

	err := cfg.ExportConfig()
	assert.NoError(t, err)

	configPath := filepath.Join(homeDir, ".fairyport", "config.yml")
	assert.FileExists(t, configPath)
}

func TestExportConfig_ExistingFile(t *testing.T) {
	cfg := config.DefaultConfig()

	// Mock the home directory and create a config file
	homeDir := t.TempDir()
	os.Setenv("HOME", homeDir)
	configDir := filepath.Join(homeDir, ".fairyport")
	os.MkdirAll(configDir, 0755)
	filePath := filepath.Join(configDir, "config.yml")
	_, err := os.Create(filePath)
	assert.NoError(t, err)

	err = cfg.ExportConfig()
	assert.NoError(t, err)
	assert.FileExists(t, filePath)
}

func TestSetInitialConfig(t *testing.T) {
	cfg := config.DefaultConfig()
	config.SetInitialConfig(cfg)

	assert.Equal(t, "127.0.0.1", viper.GetString("FairyringNodeWS.IP"))
	assert.Equal(t, int64(26657), viper.GetInt64("FairyringNodeWS.Port"))
	assert.Equal(t, "tcp", viper.GetString("FairyringNodeWS.Protocol"))
	assert.Equal(t, "fairyring-testnet-3", viper.GetString("CosmosRelayConfig.DestinationNode.ChainId"))
}
