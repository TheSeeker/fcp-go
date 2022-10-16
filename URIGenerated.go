// URIGenerated
package fcp

import (
	"fmt"
)

type URIGenerated struct {
	identifier string
	uRI        string
}

func (r *URIGenerated) parseMessage(rawMsg []string) {
	fmt.Println("%+v", rawMsg)
	for _, v := range rawMsg {
		param, val := SplitParam(v)

		switch param {
		case "URIGenerated":
			//ignore
		case "Identifier":
			r.identifier = val
		case "URI":
			r.uRI = val
		case "EndMessage":
			break
		default:
			err := fmt.Errorf("unknown parameter in URIGenerated: %+v with value of %+v", param, val)
			fmt.Println(err.Error())
		}
	}
}

func (r URIGenerated) GetName() string       { return "URIGenerated" }
func (r URIGenerated) GetIdentifier() string { return r.identifier }
func (r URIGenerated) GetURI() string        { return r.uRI }
