package fcp

import "errors"

type disconnect struct {
}

func (r *FCPClient) newDisconnect() disconnect {
	ret := disconnect{}
	return ret
}

func (r *disconnect) getMessage() message {
	return message{"Disconnect", nil, nil}
}

func (r *FCPClient) Disconnect() error {
	if r.socket == nil {
		err := errors.New("Could not Disconnect, not connected.")
		return err
	}

	disconnect := r.newDisconnect()
	r.msgSender <- disconnect.getMessage()

	return nil
}

func (r *disconnect) GetName() string { return "Disconnect" }
