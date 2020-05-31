// NodeHello
package fcp

import (
	"fmt"
	"strconv"
	"strings"
)

type nodeHello struct {
	node                 string
	fCPVersion           string
	version              []string
	build                string // integer?
	revision             string
	extBuild             string // integer?
	extRevision          string
	testnet              bool
	numcodecs            uint
	compressionCodecs    []string
	connectionIdentifier string
	nodeLanguage         string
}

func (r *nodeHello) parseMessage(rawMsg []string) nodeHello {
	for _, v := range rawMsg {
		split := strings.SplitAfterN(v, "=", 2)
		param := split[0]
		val := ""
		if strings.HasSuffix(param, "=") {
			strings.TrimSuffix(param, "=")
			val = split[1]
		}

		switch param {
		case "FCPVersion":
			r.fCPVersion = val
		case "Node":
			r.node = val
		case "Version":
			r.version = val
		case "Build":
			r.build = val
		case "Revision":
			r.revision = val
		case "ExtBuild":
			r.extBuild = valof
		case "ExtRevision":
			r.extRevision = val
		case "Testnet":
			r.testnet = (val == "true")
		case "CompressionCodecs":
			{
				incodecs := strings.Split(val, ",")                                                           // codecs have their index appended, which implies they might not be sent in numerical order?
				r.numcodecs, _ = strconv.ParseUint(strings.TrimSpace(strings.Split(incodecs[0], "-")[0]), 10) // parse the codec count
				incodecs[0] = strings.TrimSpace(strings.Splitafter(incodecs[0], "-", 2)[1])                   //strip off the count

				for _, v2 := range incodecs {
					curcodec := strings.Split(v2, "(")[0]                                   //grab just the codec part
					curcodecindex := strconv.ParseInt(strings.TrimSuffix(curcodec[1], ")")) //trim and parse the codec index.  Why did this have to be formatted so badly for machine parsing?
					r.compressionCodecs[curcodecindex] = curcodec                           //set the codec to the given index, in case they are sent out of order?
				}
			}
		case "ConnectionIdentifier":
			r.connectionIdentifier = val
		case "NodeLanguage":
			r.nodeLanguage = val
		case "EndMessage":
			break
		default:
			err := fmt.Errorf("unknown parameter in NodeHello: %s with value of %s", param, val)
		}
	}

	return r
}
