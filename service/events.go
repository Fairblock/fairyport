package service

import "strconv"

func processEvents(events map[string][]string) {
	var height uint64
	var data string
	attrs, found := events["keyshare-aggregated.keyshare-aggregated-block-height"]
	if found {
		height, _ = strconv.ParseUint(attrs[0], 10, 64)
	}

	attrs, found = events["keyshare-aggregated.keyshare-aggregated-data"]
	if found {
		data = attrs[0]
	}

	sendTx(height, data)
}
