package fcp

import (
	"strconv"
)

type watchGlobal struct {
	enabled bool
	/* bitmask ...
	SPLITFILE_PROGRESS = 1
	SENT_TO_NETWORK = 2
	COMPATIBILITY_MODE = 4
	EXPECTED_HASHES = 8
	// Nothing defined for the 16 bit
	EXPECTED_TYPE = 32
	EXPECTED_SIZE = 64
	ENTER_FINITE_COOLDOWN = 128
	PUT_FETCHABLE = 256
	COMPRESSION_START_END = 512 */
	verbosityMask uint32
}

func (r *FCPClient) newWatchGlobal() watchGlobal {
	ret := watchGlobal{false, 0}
	return ret
}

func (r *watchGlobal) setEnabled() {
	r.enabled = true
}

func (r *watchGlobal) setDsabled() {
	r.enabled = false
}

func (r *watchGlobal) setVerbosity(mask uint32) {
	r.verbosityMask = mask
}

func (r *watchGlobal) getMessage() message {
	params := []string{}
	if r.enabled {
		params = append(params, "Enabled=true")
	}
	if r.verbosityMask > 0 {
		params = append(params, "VerbosityMask"+strconv.FormatInt(int64(r.verbosityMask), 10))
	}

	return message{"WatchGlobal",
		params,
		nil} // the message type has a data field, but we aren't using it here.
}
