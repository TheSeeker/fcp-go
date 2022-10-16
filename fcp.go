// fcp project fcp.go
package fcp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type FCPClient struct {
	host       *net.TCPAddr
	isSSL      bool
	identifier string
	socket     *net.TCPConn
	msgSender  chan message
	msgHandler chan message
	caller     chan nodeMessager
}

type message struct {
	name   string   // message name
	params []string // don't include EndMessage in your params list.
	data   []byte   // if data is nil, EndMessage is automatically appended to the end of the params list.
}

type nodeMessager interface {
	parseMessage([]string)
	GetName() string
}

func NewClient(ipPort string, ssl bool, id string) (FCPClient, error) {
	addr, err := net.ResolveTCPAddr("tcp", ipPort)
	msgSender := make(chan message, 20)
	msgHandler := make(chan message, 20)
	caller := make(chan nodeMessager, 20)
	return FCPClient{addr, ssl, id, nil, msgSender, msgHandler, caller}, err
}

func (r *FCPClient) sender() {
	for {
		msg := <-r.msgSender
		fmt.Println(len(msg.data), "bytes of data after message sender receipt.")
		if r.socket == nil {
			err := errors.New("Trying to send a message while socket is closed.")
			fmt.Println(err.Error())
			return
		}
		r.socket.Write([]byte(msg.name))
		r.socket.Write([]byte("\n"))
		fmt.Println(msg.name) //crappy debug output is crappy
		for i := range msg.params {
			r.socket.Write([]byte(msg.params[i]))
			r.socket.Write([]byte("\n"))
			fmt.Println(msg.params[i]) //crappy debug output is crappy
		}

		if msg.data == nil {
			r.socket.Write([]byte("EndMessage\n"))
			fmt.Println("EndMessage") //crappy debug output is crappy
		} else {
			fmt.Println("Data: ", len(msg.data), "bytes")
			r.socket.Write([]byte("Data\n"))
			t1 := time.Now()
			written, err := r.socket.Write(msg.data)
			if err == nil {
				fmt.Println(written, "bytes written in", time.Since(t1).Seconds(), "seconds")
			} else {
				fmt.Println(err.Error())
			}
		}
	}
}

// reciever() monitors the FCP connection for incoming messages, and hands the raw messages off
// to the message handler.
func (r *FCPClient) reciever() {
	scanner := bufio.NewScanner(r.socket)
	started := false
	msg := message{}
	params := []string{}
	var datalen int64 = 0
	var data []byte = nil

	for {
		if ok := scanner.Scan(); !ok {
			break
		}
		if !started {
			msg.name = scanner.Text()
			started = true
			if ok := scanner.Scan(); !ok {
				break
			}
		}
		param, val := SplitParam(scanner.Text())
		switch param {
		case "EndMessage":
			params = append(params, scanner.Text())
			msg.params = make([]string, len(params))
			copy(msg.params, params)
			r.msgHandler <- msg
			started = false
			params = []string{""}
			datalen = 0
			data = nil
		case "DataLength":
			params = append(params, scanner.Text())
			datalen, _ = strconv.ParseInt(val, 10, 64)
		case "Data":
			params = append(params, scanner.Text())
			data = make([]byte, datalen)
			rb, err := io.ReadFull(r.socket, data)
			if int64(rb) < datalen {
				fmt.Println("Was expecting", datalen, "bytes, got", rb, "Bytes instead. ", err)
			}
			msg.data = make([]byte, datalen)
			copy(msg.data, data)

			msg.params = make([]string, len(params))
			copy(msg.params, params)
			r.msgHandler <- msg
			started = false
			params = []string{""}
			datalen = 0
			data = nil

		default:
			params = append(params, scanner.Text())

		}

		fmt.Println("RAW DUMP:", scanner.Text())
	}
	fmt.Println("Connection Closed")
	os.Exit(1)
}

// handler(nodeMessager) listens on a message channel, creates and populates the appropriate
// type, then hands off the message to the provided nodeMessager channel.
func (r *FCPClient) handler(caller chan nodeMessager) {
	for {
		msg := <-r.msgHandler

		// This is the list of all known FCP Server->Client messages.
		// If I can figure out how to make reflection work, I might be able to generalize these
		// since they all have the exact same pattern
		switch msg.name {
		case "NodeHello":
			nh := &NodeHello{}
			nh.parseMessage(msg.params)
			caller <- nh
		case "CloseConnectionDuplicateClientName":
			fmt.Println("Unimplemented message:", msg.name)
		case "Peer":
			fmt.Println("Unimplemented message:", msg.name)
		case "PeerNote":
			fmt.Println("Unimplemented message:", msg.name)
		case "EndListPeers":
			fmt.Println("Unimplemented message:", msg.name)
		case "EndListPeerNotes":
			fmt.Println("Unimplemented message:", msg.name)
		case "PeerRemoved":
			fmt.Println("Unimplemented message:", msg.name)
		case "NodeData":
			fmt.Println("Unimplemented message:", msg.name)
		case "ConfigData": // (since 1027)
			fmt.Println("Unimplemented message:", msg.name)
		case "TestDDAReply": // (since 1027)
			fmt.Println("Unimplemented message:", msg.name)
		case "TestDDAComplete": // (since 1027)
			fmt.Println("Unimplemented message:", msg.name)
		case "SSKKeyPair":
			fmt.Println("Unimplemented message:", msg.name)
		case "PersistentGet":
			fmt.Println("Unimplemented message:", msg.name)
		case "PersistentPut":
			fmt.Println("Unimplemented message:", msg.name)
		case "PersistentPutDir":
			fmt.Println("Unimplemented message:", msg.name)
		case "URIGenerated":
			ug := &URIGenerated{}
			ug.parseMessage(msg.params)
			caller <- ug
		case "PutSuccessful":
			ps := &PutSuccessful{}
			ps.parseMessage(msg.params)
			caller <- ps
		case "PutFetchable":
			fmt.Println("Unimplemented message:", msg.name)
		case "DataFound":
			fmt.Println("Unimplemented message:", msg.name)
		case "GetRequestStatus":
			fmt.Println("Unimplemented message:", msg.name)
		case "AllData":
			ad := &AllData{}
			ad.parseMessage(msg.params)
			ad.SetData(msg.data)
			caller <- ad
		case "StartedCompression":
			fmt.Println("Unimplemented message:", msg.name)
		case "FinishedCompression":
			fmt.Println("Unimplemented message:", msg.name)
		case "SimpleProgress":
			sp := &SimpleProgress{}
			sp.parseMessage(msg.params)
			caller <- sp
		case "ExpectedHashes": // (since 1254)
			eh := &ExpectedHashes{}
			eh.parseMessage(msg.params)
			caller <- eh
		case "ExpectedMIME": // (since 1307)
			fmt.Println("Unimplemented message:", msg.name)
		case "ExpectedDataLength": // (since 1307)
			fmt.Println("Unimplemented message:", msg.name)
		case "CompatibilityMode": // (since 1254)
			fmt.Println("Unimplemented message:", msg.name)
		case "EndListPersistentRequests":
			fmt.Println("Unimplemented message:", msg.name)
		case "PersistentRequestRemoved": // (since 1016)
			fmt.Println("Unimplemented message:", msg.name)
		case "PersistentRequestModified": // (since 1016)
			fmt.Println("Unimplemented message:", msg.name)
		case "SendingToNetwork": // (since 1207)
			fmt.Println("Unimplemented message:", msg.name)
		case "EnterFiniteCooldown": // (since 1365)
			fmt.Println("Unimplemented message:", msg.name)
		case "GeneratedMetadata": // (since 1380)
			fmt.Println("Unimplemented message:", msg.name)
		case "PutFailed":
			fmt.Println("Unimplemented message:", msg.name)
		case "GetFailed":
			fmt.Println("Unimplemented message:", msg.name)
		case "ProtocolError":
			fmt.Println("Unimplemented message:", msg.name)
		case "IdentifierCollision":
			fmt.Println("Unimplemented message:", msg.name)
		case "UnknownNodeIdentifier":
			fmt.Println("Unimplemented message:", msg.name)
		case "UnknownPeerNoteType":
			fmt.Println("Unimplemented message:", msg.name)
		case "SubscribedUSK":
			fmt.Println("Unimplemented message:", msg.name)
		case "SubscribedUSKUpdate":
			fmt.Println("Unimplemented message:", msg.name)
		case "SubscribedUSKSendingToNetwork": //(since 1365)
			fmt.Println("Unimplemented message:", msg.name)
		case "SubscribedUSKRoundFinished": // (since 1365)
			fmt.Println("Unimplemented message:", msg.name)
		case "PluginInfo": // (since 1075)
			fmt.Println("Unimplemented message:", msg.name)
		case "PluginRemoved": // (since 1227)
			fmt.Println("Unimplemented message:", msg.name)
		case "FCPPluginReply": // (since 1075)
			fmt.Println("Unimplemented message:", msg.name)
		case "ProbeBandwidth":
			fmt.Println("Unimplemented message:", msg.name)
		case "ProbeBuild":
			fmt.Println("Unimplemented message:", msg.name)
		case "ProbeError":
			fmt.Println("Unimplemented message:", msg.name)
		case "ProbeIdentifier":
			fmt.Println("Unimplemented message:", msg.name)
		case "ProbeLinkLengths":
			fmt.Println("Unimplemented message:", msg.name)
		case "ProbeLocation":
			fmt.Println("Unimplemented message:", msg.name)
		case "ProbeRefused":
			fmt.Println("Unimplemented message:", msg.name)
		case "ProbeRejectStats":
			fmt.Println("Unimplemented message:", msg.name)
		case "ProbeStoreSize":
			fmt.Println("Unimplemented message:", msg.name)
		case "ProbeUptime":
			fmt.Println("Unimplemented message:", msg.name)
		default:
			fmt.Println("Unknown Node Message " + msg.name + " | " + fmt.Sprint(msg.params))
			//Unknown Message
		}
	}
}

// Caller() returns the callback channel for this instance of fcp-go
func (r *FCPClient) Caller() chan nodeMessager {
	return r.caller
}

// SplitParam() takes a raw string from an FCP message and split it into a parameter and a value.
// If there is no value, just return the parameter and an empty string.
func SplitParam(v string) (param string, val string) {
	val = ""
	param = ""
	split := strings.SplitAfterN(v, "=", 2)
	param = split[0]
	if strings.HasSuffix(param, "=") {
		param = strings.TrimSuffix(param, "=")
		val = split[1]
	}
	return
}
