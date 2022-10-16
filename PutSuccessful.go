// PutSuccessful
package fcp

import (
	"fmt"
	"strconv"
	"time"
)

type PutSuccessful struct {
	global         bool
	identifier     string
	startupTime    time.Time
	completionTime time.Time
	uRI            string
}

func (r *PutSuccessful) parseMessage(rawMsg []string) {
	fmt.Println("%+v", rawMsg)
	for _, v := range rawMsg {
		param, val := SplitParam(v)

		switch param {
		case "PutSuccessful":
		//ignore
		case "Global":
			r.global = (val == "true")
		case "Identifier":
			r.identifier = val
		case "StartupTime":
			timeInt, _ := strconv.ParseInt(val, 10, 64)
			r.startupTime = time.Unix(timeInt/100, 0)
		case "CompletionTime":
			timeInt, _ := strconv.ParseInt(val, 10, 64)
			r.completionTime = time.Unix(timeInt/100, 0)
		case "URI":
			r.uRI = val
		case "EndMessage":
			break
		default:
			err := fmt.Errorf("unknown parameter in PutSuccessful: %+v with value of %+v", param, val)
			fmt.Println(err.Error())
		}
	}
}

func (r PutSuccessful) GetName() string              { return "PutSuccessful" }
func (r PutSuccessful) GetGlobal() bool              { return r.global }
func (r PutSuccessful) GetIdentifier() string        { return r.identifier }
func (r PutSuccessful) GetStartupTime() time.Time    { return r.startupTime }
func (r PutSuccessful) GetCompletionTime() time.Time { return r.startupTime }
func (r PutSuccessful) GetURI() string               { return r.uRI }
