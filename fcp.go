// fcp project fcp.go
package fcp

import (
	"errors"
	"fmt"
	"net"
)

type FCPClient struct {
	host       *net.TCPAddr
	isSSL      bool
	identifier string
	socket     *net.TCPConn
	msgSender  chan message
}

type message struct {
	message string
	params  []string
	data    *[]byte
}

func NewClient(ipPort string, ssl bool, id string) (FCPClient, error) {
	addr, err := net.ResolveTCPAddr("tcp", ipPort)
	msgSender := make(chan message, 20)
	return FCPClient{addr, ssl, id, nil, msgSender}, err
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

func (r *FCPClient) reciever() {
	buf := make([]byte, 1200)
	for {
		rbytes, _ := r.socket.Read(buf)
		//do stuff with this data, like compose messages out of it and notify
		//the client via channels or something...
		fmt.Print(string(buf[:rbytes]))
	}
}
