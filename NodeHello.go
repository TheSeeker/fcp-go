// NodeHello
package fcp

import (
	"fmt"
	"strconv"
	"strings"
)

type NodeHello struct {
	node                 string
	fCPVersion           string
	version              []string
	build                string // integer?
	revision             string
	extBuild             string // integer?
	extRevision          string
	testnet              bool
	numcodecs            uint64
	compressionCodecs    []string
	connectionIdentifier string
	nodeLanguage         string
}

func (r *NodeHello) parseMessage(rawMsg []string) {
	for _, v := range rawMsg {
		param, val := SplitParam(v)

		switch param {
		case "NodeHello":
		//ignore
		case "FCPVersion":
			r.fCPVersion = val
		case "Node":
			r.node = val
		case "Version":
			r.version = strings.Split(val, ",")
		case "Build":
			r.build = val
		case "Revision":
			r.revision = val
		case "ExtBuild":
			r.extBuild = val
		case "ExtRevision":
			r.extRevision = val
		case "Testnet":
			r.testnet = (val == "true")
		case "CompressionCodecs":
			{
				incodecs := strings.Split(val, ",")                                                              // codecs have their index appended, which implies they might not be sent in numerical order?
				r.numcodecs, _ = strconv.ParseUint(strings.TrimSpace(strings.Split(incodecs[0], "-")[0]), 10, 8) // parse the codec count
				r.compressionCodecs = make([]string, r.numcodecs)
				incodecs[0] = strings.TrimSpace(strings.SplitAfterN(incodecs[0], "-", 2)[1]) //strip off the count
				for _, v2 := range incodecs {
					curcodec := strings.Split(v2, "(")[0]                                                           //grab just the codec part
					curcodecindex, _ := strconv.ParseInt(strings.TrimSuffix(strings.Split(v2, "(")[1], ")"), 10, 8) //trim and parse the codec index.  Why did this have to be formatted so badly for machine parsing?
					r.compressionCodecs[curcodecindex] = curcodec                                                   //set the codec to the given index, in case they are sent out of order?
				}
			}
		case "ConnectionIdentifier":
			r.connectionIdentifier = val
		case "NodeLanguage":
			r.nodeLanguage = val
		case "EndMessage":
			break
		default:
			err := fmt.Errorf("unknown parameter in NodeHello: %+v with value of %+v", param, val)
			fmt.Println(err.Error())
		}
	}
}

// Getters
func (r *NodeHello) GetName() string                 { return "NodeHello" }
func (r *NodeHello) GetNode() string                 { return r.node }
func (r *NodeHello) GetFCPVersion() string           { return r.fCPVersion }
func (r *NodeHello) GetVersion() []string            { return r.version }
func (r *NodeHello) GetBuild() string                { return r.build } // integer?
func (r *NodeHello) GetRevision() string             { return r.revision }
func (r *NodeHello) GetExtBuild() string             { return r.extBuild } // integer?
func (r *NodeHello) GetExtRevision() string          { return r.extRevision }
func (r *NodeHello) GetTestnet() bool                { return r.testnet }
func (r *NodeHello) GetNumcodecs() uint64            { return r.numcodecs }
func (r *NodeHello) GetCompressionCodecs() []string  { return r.compressionCodecs }
func (r *NodeHello) GetConnectionIdentifier() string { return r.connectionIdentifier }
func (r *NodeHello) GetNodeLanguage() string         { return r.nodeLanguage }
