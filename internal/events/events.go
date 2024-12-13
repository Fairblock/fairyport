package events

import (
	"errors"
	"strconv"
)

func ProcessEvents(events map[string][]string) (uint64, string, string, error) {

	attrs, found := events["keyshare-aggregated.height"]
	if !found {
		oldAttrs, oldAttrfound := events["keyshare-aggregated.keyshare-aggregated-block-height"]
		if !oldAttrfound {
			return 0, "", "", errors.New("aggregated keyshare event not found")
		}
		attrs = oldAttrs
	}

	height, err := strconv.ParseUint(attrs[0], 10, 64)
	if err != nil {
		return 0, "", "", err
	}

	dataAttrs, found := events["keyshare-aggregated.data"]
	if !found {
		oldDataAttrs, oldDataAttrfound := events["keyshare-aggregated.keyshare-aggregated-data"]
		if !oldDataAttrfound {
			return 0, "", "", errors.New("aggregated keyshare event data not found")
		}
		dataAttrs = oldDataAttrs
	}

	pubKeyAttrs, found := events["keyshare-aggregated.pubkey"]
	if !found {
		return 0, "", "", errors.New("aggregated keyshare event pubkey not found")
	}

	return height, dataAttrs[0], pubKeyAttrs[0], nil
}
