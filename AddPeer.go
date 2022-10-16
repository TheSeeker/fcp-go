package fcp

type addPeer struct { // Possible values 				Mandatory 	Description
	trust      string // LOW, NORMAL, HIGH		 	yes 			Sets the trust you put in the new peer.
	visibility string // NO, NAME_ONLY, YES 			yes 			Determines whether the new peer will be visible for your peers.
	file       string // filename 					no 			If set, the peer noderef is read from the given filename relative to the node's current directory.
	uRL        string // URL 						no		 	If set, the peer noderef is read from the given URL.
	//<raw_noderef> 			// <key>=<value>			no 			If neither File nor URL are set, the field set is read as a noderef.
	//nodeRef NodeRefSFS	//	                                 TODO: implement allowing a manually configured noderef?
}

func (r *FCPClient) NewAddPeer() addPeer {
	return addPeer{
		"", "", "", "",
	}
}

func (r *addPeer) SetTrust(v string) error {
	//FIXME add sanity checking.
	r.trust = v
	return nil
}

func (r *addPeer) SetVisibility(v string) error {
	//FIXME add sanity checking.
	r.visibility = v
	return nil
}

func (r *addPeer) SetFile(v string) error {
	//FIXME add sanity checking.
	r.file = v
	return nil
}

func (r *addPeer) SetURL(v string) error {
	//FIXME add sanity checking.
	r.uRL = v
	return nil
}

func (r *addPeer) getMessage() message {

	params := []string{}
	params = append(params, "Trust="+r.trust)
	params = append(params, "Visibility="+r.visibility)
	if r.file != "" {
		params = append(params, "File="+r.file)
	} else if r.uRL != "" {
		params = append(params, "URL="+r.uRL)
	} else {
		// This is where we'd parse a raw noderef.
		// unimplemented
	}
	params = append(params, "EndMessage")
	return message{"AddPeer",
		params,
		nil}
}
