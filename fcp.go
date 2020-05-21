// fcp project fcp.go
package fcp

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
)

type FCPClient struct {
	host       *net.TCPAddr
	isSSL      bool
	identifier string
	socket     *net.TCPConn
	msgSender  chan message
	msgHandler chan message
}

type message struct {
	message string
	params  []string
	data    *[]byte
}

func NewClient(ipPort string, ssl bool, id string) (FCPClient, error) {
	addr, err := net.ResolveTCPAddr("tcp", ipPort)
	msgSender := make(chan message, 20)
	msgHandler := make(chan message, 20)
	return FCPClient{addr, ssl, id, nil, msgSender, msgHandler}, err
}

func (r *FCPClient) sender() {
	for {
		msg := <-r.msgSender
		if r.socket == nil {
			err := errors.New("Trying to send a message while socket is closed.")
			fmt.Println(err.Error())
			return
		}
		r.socket.Write([]byte(msg.message))
		r.socket.Write([]byte("\n"))
		fmt.Println(msg.message) //crappy debug output is crappy
		for i := range msg.params {
			r.socket.Write([]byte(msg.params[i]))
			r.socket.Write([]byte("\n"))
			fmt.Println(msg.params[i]) //crappy debug output is crappy
		}

		if msg.data == nil {
			r.socket.Write([]byte("EndMessage\n"))
			fmt.Println("EndMessage") //crappy debug output is crappy
		} else {
			r.socket.Write([]byte("Data\n"))
			r.socket.Write(*msg.data)
		}
	}
}

func (r *FCPClient) handler() {
	for {
		msg := <-r.msgHandler

		switch msg.message {
		case "NodeHello":
		case "CloseConnectionDuplicateClientName":
		case "Peer":
		case "PeerNote":
		case "EndListPeers":
		case "EndListPeerNotes":
		case "PeerRemoved":
		case "NodeData":
		case "ConfigData": // (since 1027)
		case "TestDDAReply": // (since 1027)
		case "TestDDAComplete": // (since 1027)
		case "SSKKeyPair":
		case "PersistentGet":
		case "PersistentPut":
		case "PersistentPutDir":
		case "URIGenerated":
		case "PutSuccessful":
		case "PutFetchable":
		case "DataFound":
		case "GetRequestStatus":
		case "AllData":
		case "StartedCompression":
		case "FinishedCompression":
		case "SimpleProgress":
		case "ExpectedHashes": // (since 1254)
		case "ExpectedMIME": // (since 1307)
		case "ExpectedDataLength": // (since 1307)
		case "CompatibilityMode": // (since 1254)
		case "EndListPersistentRequests":
		case "PersistentRequestRemoved": // (since 1016)
		case "PersistentRequestModified": // (since 1016)
		case "SendingToNetwork": // (since 1207)
		case "EnterFiniteCooldown": // (since 1365)
		case "GeneratedMetadata": // (since 1380)

		case "PutFailed":
		case "GetFailed":
		case "ProtocolError":
		case "IdentifierCollision":
		case "UnknownNodeIdentifier":
		case "UnknownPeerNoteType":
		case "SubscribedUSK":
		case "SubscribedUSKUpdate":
		case "SubscribedUSKSendingToNetwork": //(since 1365)
		case "SubscribedUSKRoundFinished": // (since 1365)

		case "PluginInfo": // (since 1075)
		case "PluginRemoved": // (since 1227)
		case "FCPPluginReply": // (since 1075)

		case "ProbeBandwidth":
		case "ProbeBuild":
		case "ProbeError":
		case "ProbeIdentifier":
		case "ProbeLinkLengths":
		case "ProbeLocation":
		case "ProbeRefused":
		case "ProbeRejectStats":
		case "ProbeStoreSize":
		case "ProbeUptime":
		default:
			//Unknown Message
		}
	}
}

func (r *FCPClient) reciever() {
	scanner := bufio.NewScanner(r.socket)
	for {
		if ok := scanner.Scan(); !ok {
			break
		}

		//do stuff with this data, like compose messages out of it and notify
		//the client via channels or something...
		fmt.Println(scanner.Text())
	}
	fmt.Println("Connection Closed")
	os.Exit(1)
}
