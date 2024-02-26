package events

import (
	"errors"
	"strconv"
)

func ProcessEvents(events map[string][]string) (uint64, string, string, error) {

	attrs, found := events["keyshare-aggregated.height"]
	if !found {
		return 0, "", "", errors.New("aggregated keyshare event not found")
	}

	height, err := strconv.ParseUint(attrs[0], 10, 64)
	if err != nil {
		return 0, "", "", err
	}

	dataAttrs, found := events["keyshare-aggregated.data"]
	if !found {
		return 0, "", "", errors.New("aggregated keyshare event data not found")
	}

	pubKeyAttr, found := events["keyshare-aggregated.pubkey"]
	if !found {
		return 0, "", "", errors.New("aggregated keyshare event pubkey not found")
	}

	return height, dataAttrs[0], pubKeyAttr[0], nil
}
