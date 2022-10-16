// AllData
package fcp

import (
	"fmt"
	"strconv"
	"time"
)

type AllData struct {
	identifier           string
	completionTime       time.Time
	startupTime          time.Time
	dataLength           uint64
	global               bool
	metadata_ContentType string
	data                 []byte
}

func (r *AllData) parseMessage(rawMsg []string) {
	fmt.Println("%+v", rawMsg)
	for _, v := range rawMsg {
		param, val := SplitParam(v)

		switch param {
		case "AllData":
		//ignore
		case "Identifier":
			r.identifier = val
		case "CompletionTime":
			timeInt, _ := strconv.ParseInt(val, 10, 64)
			r.completionTime = time.Unix(timeInt/100, 0)
		case "StartupTime":
			timeInt, _ := strconv.ParseInt(val, 10, 64)
			r.startupTime = time.Unix(timeInt/100, 0)
		case "DataLength":
			r.dataLength, _ = strconv.ParseUint(val, 10, 64)
		case "Global":
			r.global = (val == "true")
		case "Metadata.ContentType":
			r.metadata_ContentType = val
		default:
			err := fmt.Errorf("unknown parameter in AllData: %+v with value of %+v", param, val)
			fmt.Println(err.Error())
		}
	}
}

func (r *AllData) GetName() string                 { return "AllData" }
func (r *AllData) GetIdentifier() string           { return r.identifier }
func (r *AllData) GetCompletionTime() time.Time    { return r.completionTime }
func (r *AllData) GetStartupTime() time.Time       { return r.startupTime }
func (r *AllData) GetDataLength() uint64           { return r.dataLength }
func (r *AllData) GetGlobal() bool                 { return r.global }
func (r *AllData) GetMetadata_ContentType() string { return r.metadata_ContentType }
func (r *AllData) GetData() []byte                 { return r.data }

func (r *AllData) SetData(buf []byte) {
	r.data = buf
}
