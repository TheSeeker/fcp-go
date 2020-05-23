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
	caller     chan clientMessage
}

type message struct {
	name   string
	params []string
	data   *[]byte
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
			r.socket.Write([]byte("Data\n"))
			r.socket.Write(*msg.data)
		}
	}
}

func (r *FCPClient) handler(caller chan message) {
	for {
		msg := <-r.msgHandler

		switch msg.name {
		case "NodeHello":
			caller <- nodeHello.parseMessage(params)
		case "CloseConnectionDuplicateClientName":
			fmt.Println("Unimplemented")
		case "Peer":
			fmt.Println("Unimplemented")
		case "PeerNote":
			fmt.Println("Unimplemented")
		case "EndListPeers":
			fmt.Println("Unimplemented")
		case "EndListPeerNotes":
			fmt.Println("Unimplemented")
		case "PeerRemoved":
			fmt.Println("Unimplemented")
		case "NodeData":
			fmt.Println("Unimplemented")
		case "ConfigData": // (since 1027)
			fmt.Println("Unimplemented")
		case "TestDDAReply": // (since 1027)
			fmt.Println("Unimplemented")
		case "TestDDAComplete": // (since 1027)
			fmt.Println("Unimplemented")
		case "SSKKeyPair":
			fmt.Println("Unimplemented")
		case "PersistentGet":
			fmt.Println("Unimplemented")
		case "PersistentPut":
			fmt.Println("Unimplemented")
		case "PersistentPutDir":
			fmt.Println("Unimplemented")
		case "URIGenerated":
			fmt.Println("Unimplemented")
		case "PutSuccessful":
			fmt.Println("Unimplemented")
		case "PutFetchable":
			fmt.Println("Unimplemented")
		case "DataFound":
			fmt.Println("Unimplemented")
		case "GetRequestStatus":
			fmt.Println("Unimplemented")
		case "AllData":
			fmt.Println("Unimplemented")
		case "StartedCompression":
			fmt.Println("Unimplemented")
		case "FinishedCompression":
			fmt.Println("Unimplemented")
		case "SimpleProgress":
			fmt.Println("Unimplemented")
		case "ExpectedHashes": // (since 1254)
			fmt.Println("Unimplemented")
		case "ExpectedMIME": // (since 1307)
			fmt.Println("Unimplemented")
		case "ExpectedDataLength": // (since 1307)
			fmt.Println("Unimplemented")
		case "CompatibilityMode": // (since 1254)
			fmt.Println("Unimplemented")
		case "EndListPersistentRequests":
			fmt.Println("Unimplemented")
		case "PersistentRequestRemoved": // (since 1016)
			fmt.Println("Unimplemented")
		case "PersistentRequestModified": // (since 1016)
			fmt.Println("Unimplemented")
		case "SendingToNetwork": // (since 1207)
			fmt.Println("Unimplemented")
		case "EnterFiniteCooldown": // (since 1365)
			fmt.Println("Unimplemented")
		case "GeneratedMetadata": // (since 1380)
			fmt.Println("Unimplemented")

		case "PutFailed":
			fmt.Println("Unimplemented")
		case "GetFailed":
			fmt.Println("Unimplemented")
		case "ProtocolError":
			fmt.Println("Unimplemented")
		case "IdentifierCollision":
			fmt.Println("Unimplemented")
		case "UnknownNodeIdentifier":
			fmt.Println("Unimplemented")
		case "UnknownPeerNoteType":
			fmt.Println("Unimplemented")
		case "SubscribedUSK":
			fmt.Println("Unimplemented")
		case "SubscribedUSKUpdate":
			fmt.Println("Unimplemented")
		case "SubscribedUSKSendingToNetwork": //(since 1365)
			fmt.Println("Unimplemented")
		case "SubscribedUSKRoundFinished": // (since 1365)
			fmt.Println("Unimplemented")

		case "PluginInfo": // (since 1075)
			fmt.Println("Unimplemented")
		case "PluginRemoved": // (since 1227)
			fmt.Println("Unimplemented")
		case "FCPPluginReply": // (since 1075)
			fmt.Println("Unimplemented")

		case "ProbeBandwidth":
			fmt.Println("Unimplemented")
		case "ProbeBuild":
			fmt.Println("Unimplemented")
		case "ProbeError":
			fmt.Println("Unimplemented")
		case "ProbeIdentifier":
			fmt.Println("Unimplemented")
		case "ProbeLinkLengths":
			fmt.Println("Unimplemented")
		case "ProbeLocation":
			fmt.Println("Unimplemented")
		case "ProbeRefused":
			fmt.Println("Unimplemented")
		case "ProbeRejectStats":
			fmt.Println("Unimplemented")
		case "ProbeStoreSize":
			fmt.Println("Unimplemented")
		case "ProbeUptime":
			fmt.Println("Unimplemented")
		default:
			fmt.Println("Unknown Node Message " + msg.name + " | " + msg.params)
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
