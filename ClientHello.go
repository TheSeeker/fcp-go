package fcp

import "net"

type clientHello struct {
	name,
	expectedVersion string
}

func (r *FCPClient) newClientHello() clientHello {
	ret := clientHello{"", "2.0"}
	return ret
}
func (r *clientHello) setName(newname string) error {
	r.name = newname
	return nil // maybe do checking on this later?
}

func (r *clientHello) getMessage() message {
	return message{"ClientHello",
		[]string{
			"Name=" + r.name,
			"ExpectedVersion=" + r.expectedVersion,
		}, nil}
}

func (r *FCPClient) Connect() error {
	socket, err := net.DialTCP("tcp", nil, r.host)
	if err != nil {
		return err
	}
	r.socket = socket

	go r.sender()
	go r.reciever()
	go r.handler()

	hello := r.newClientHello()
	hello.setName(r.identifier)
	r.msgSender <- hello.getMessage()

	return nil
}
