package fcp

import (
	"fmt"
	"strconv"
	"time"
)

/*
SimpleProgress
Total=12288 // 12,288 blocks we can fetch
Required=8192 // we only need 8,192 of them (because of splitfile redundancy)
Failed=452 // 452 of them have failed due to running out of retries
FatallyFailed=0 // none of them have encountered fatal errors
Succeeded=1027 // we have successfully fetched 1,027 blocks
FinalizedTotal=true // the Total will not increase any further (if this is false, it may increase; it will never decrease)
Identifier=Request Number One
Global=true // true if the request is on the global queue, false otherwise
EndMessage
*/
type SimpleProgress struct {
	succeeded             uint64
	identifier            string
	required              uint64
	finalizedTotal        bool
	minSuccessFetchBlocks uint64
	failed                uint64
	total                 uint64
	lastProgress          time.Time
	fatallyFailed         uint64
	global                bool
}

func (r *SimpleProgress) parseMessage(rawMsg []string) {
	fmt.Println("%+v", rawMsg)
	for _, v := range rawMsg {
		param, val := SplitParam(v)

		switch param {
		case "SimpleProgress":
		//ignore
		case "Succeeded":
			r.succeeded, _ = strconv.ParseUint(val, 10, 64)
		case "Identifier":
			r.identifier = val
		case "Required":
			r.required, _ = strconv.ParseUint(val, 10, 64)
		case "FinalizedTotal":
			r.finalizedTotal = (val == "true")
		case "MinSuccessFetchBlocks":
			r.minSuccessFetchBlocks, _ = strconv.ParseUint(val, 10, 64)
		case "Failed":
			r.failed, _ = strconv.ParseUint(val, 10, 64)
		case "Total":
			r.total, _ = strconv.ParseUint(val, 10, 64)
		case "LastProgress":
			timeInt, _ := strconv.ParseInt(val, 10, 64)
			r.lastProgress = time.Unix(timeInt/100, 0)
		case "FatallyFailed":
			r.fatallyFailed, _ = strconv.ParseUint(val, 10, 64)
		case "Global":
			r.global = (val == "true")
		case "EndMessage":
			break
		default:
			err := fmt.Errorf("unknown parameter in SimpleProgress: %+v with value of %+v", param, val)
			fmt.Println(err.Error())
		}
	}
}

func (r SimpleProgress) GetName() string                  { return "SimpleProgress" }
func (r SimpleProgress) GetSucceeded() uint64             { return r.succeeded }
func (r SimpleProgress) GetIdentifier() string            { return r.identifier }
func (r SimpleProgress) GetRequired() uint64              { return r.required }
func (r SimpleProgress) GetFinalizedTotal() bool          { return r.finalizedTotal }
func (r SimpleProgress) GetMinSuccessFetchBlocks() uint64 { return r.minSuccessFetchBlocks }
func (r SimpleProgress) GetFailed() uint64                { return r.failed }
func (r SimpleProgress) GetTtotal() uint64                { return r.total }
func (r SimpleProgress) GetLastProgress() time.Time       { return r.lastProgress }
func (r SimpleProgress) GetFatallyFailed() uint64         { return r.fatallyFailed }
func (r SimpleProgress) GetGlobal() bool                  { return r.global }
