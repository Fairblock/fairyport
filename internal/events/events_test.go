package events_test

import (
	"testing"

	"github.com/Fairblock/fairyport/internal/events"
	"github.com/stretchr/testify/assert"
)

func TestProcessEvents_Success(t *testing.T) {
	eventData := map[string][]string{
		"keyshare-aggregated.height": {"12345"},
		"keyshare-aggregated.data":   {"mockData"},
		"keyshare-aggregated.pubkey": {"mockPubKey"},
	}

	height, data, pubKey, err := events.ProcessEvents(eventData)
	assert.NoError(t, err)
	assert.Equal(t, uint64(12345), height)
	assert.Equal(t, "mockData", data)
	assert.Equal(t, "mockPubKey", pubKey)
}

func TestProcessEvents_OldHeightAttributes(t *testing.T) {
	eventData := map[string][]string{
		"keyshare-aggregated.keyshare-aggregated-block-height": {"54321"},
		"keyshare-aggregated.data":                             {"oldMockData"},
		"keyshare-aggregated.pubkey":                           {"oldMockPubKey"},
	}

	height, data, pubKey, err := events.ProcessEvents(eventData)
	assert.NoError(t, err)
	assert.Equal(t, uint64(54321), height)
	assert.Equal(t, "oldMockData", data)
	assert.Equal(t, "oldMockPubKey", pubKey)
}

func TestProcessEvents_HeightNotFound(t *testing.T) {
	eventData := map[string][]string{
		"keyshare-aggregated.data":   {"mockData"},
		"keyshare-aggregated.pubkey": {"mockPubKey"},
	}

	height, data, pubKey, err := events.ProcessEvents(eventData)
	assert.Error(t, err)
	assert.Equal(t, uint64(0), height)
	assert.Equal(t, "", data)
	assert.Equal(t, "", pubKey)
	assert.EqualError(t, err, "aggregated keyshare event not found")
}

func TestProcessEvents_InvalidHeight(t *testing.T) {
	eventData := map[string][]string{
		"keyshare-aggregated.height": {"invalidHeight"},
		"keyshare-aggregated.data":   {"mockData"},
		"keyshare-aggregated.pubkey": {"mockPubKey"},
	}

	height, data, pubKey, err := events.ProcessEvents(eventData)
	assert.Error(t, err)
	assert.Equal(t, uint64(0), height)
	assert.Equal(t, "", data)
	assert.Equal(t, "", pubKey)
}

func TestProcessEvents_DataNotFound(t *testing.T) {
	eventData := map[string][]string{
		"keyshare-aggregated.height": {"12345"},
		"keyshare-aggregated.pubkey": {"mockPubKey"},
	}

	height, data, pubKey, err := events.ProcessEvents(eventData)
	assert.Error(t, err)
	assert.Equal(t, uint64(0), height)
	assert.Equal(t, "", data)
	assert.Equal(t, "", pubKey)
	assert.EqualError(t, err, "aggregated keyshare event data not found")
}

func TestProcessEvents_PubKeyNotFound(t *testing.T) {
	eventData := map[string][]string{
		"keyshare-aggregated.height": {"12345"},
		"keyshare-aggregated.data":   {"mockData"},
	}

	height, data, pubKey, err := events.ProcessEvents(eventData)
	assert.Error(t, err)
	assert.Equal(t, uint64(0), height)
	assert.Equal(t, "", data)
	assert.Equal(t, "", pubKey)
	assert.EqualError(t, err, "aggregated keyshare event pubkey not found")
}
