package fcp

import "strconv"

type clientPut struct {
	uRI                              string
	contentType                      string
	identifier                       string
	verbosity                        byte
	maxRetries                       int64
	priorityClass                    int64
	getCHKOnly                       bool
	global                           bool
	dontCompress                     bool
	codecs                           []string
	clientToken                      string
	persistence                      string
	targetFilename                   string
	earlyEncode                      bool
	uploadFrom                       string
	dataLength                       uint64
	filename                         string
	targetURI                        string
	fileHash                         string
	binaryBlob                       bool
	forkOnCacheable                  bool
	extraInsertsSingleBlock          uint64
	extraInsertsSplitfileHeaderBlock uint64
	compatibilityMode                string
	localRequestOnly                 bool
	overrideSplitfileCryptoKey       string
	realTimeFlag                     bool
	metadataThreshold                int64
	data                             *[]byte
}

func (r *FCPClient) NewClientPut() clientPut {
	return clientPut{
		"", "", "",
		0, 0, 2,
		false, false, false,
		[]string{},
		"", "", "",
		false,
		"direct",
		0, "", "", "",
		false, true,
		2, 2,
		"COMPAT_CURRENT",
		false,
		"",
		false,
		-1,
		&[]byte{},
	}
}

func (r *clientPut) SetURI(v string) error {
	//FIXME add sanity checking.
	r.uRI = v
	return nil
}
func (r *clientPut) SetContentType(v string) error {
	//FIXME add sanity checking.
	r.contentType = v
	return nil
}
func (r *clientPut) SetIdentifier(v string) error {
	//FIXME add sanity checking.
	r.identifier = v
	return nil
}
func (r *clientPut) SetVerbosity(v byte) error {
	//FIXME add sanity checking.
	r.verbosity = v
	return nil
}
func (r *clientPut) SetMaxRetries(v int64) error {
	//FIXME add sanity checking.
	r.maxRetries = v
	return nil
}
func (r *clientPut) SetPriorityClass(v int64) error {
	//FIXME add sanity checking.
	r.priorityClass = v
	return nil
}
func (r *clientPut) SetGetCHKOnly(v bool) error {
	//FIXME add sanity checking.
	r.getCHKOnly = v
	return nil
}
func (r *clientPut) SetGlobal(v bool) error {
	//FIXME add sanity checking.
	r.global = v
	return nil
}
func (r *clientPut) SetDontCompress(v bool) error {
	//FIXME add sanity checking.
	r.dontCompress = v
	return nil
}
func (r *clientPut) SetCodecs(v []string) error { //FIXME add sanity checking.
	r.codecs = v
	return nil
}
func (r *clientPut) SetClientToken(v string) error { //FIXME add sanity checking.
	r.clientToken = v
	return nil
}
func (r *clientPut) SetPersistence(v string) error { //FIXME add sanity checking.
	r.persistence = v
	return nil
}
func (r *clientPut) SetTargetFilename(v string) error {
	//FIXME add sanity checking.
	r.targetFilename = v
	return nil
}
func (r *clientPut) SetEarlyEncode(v bool) error {
	//FIXME add sanity checking.
	r.earlyEncode = v
	return nil
}
func (r *clientPut) SetUploadFrom(v string) error {
	//FIXME add sanity checking.
	r.uploadFrom = v
	return nil
}
func (r *clientPut) SetDataLength(v uint64) error {
	//FIXME add sanity checking.
	r.dataLength = v
	return nil
}
func (r *clientPut) SetFilename(v string) error {
	//FIXME add sanity checking.
	r.filename = v
	return nil
}
func (r *clientPut) SetTargetURI(v string) error {
	//FIXME add sanity checking.
	r.targetURI = v
	return nil
}
func (r *clientPut) SetFileHash(v string) error {
	//FIXME add sanity checking.
	r.fileHash = v
	return nil
}
func (r *clientPut) SetBinaryBlob(v bool) error {
	//FIXME add sanity checking.
	r.binaryBlob = v
	return nil
}
func (r *clientPut) SetForkOnCacheable(v bool) error {
	//FIXME add sanity checking.
	r.forkOnCacheable = v
	return nil
}
func (r *clientPut) SetExtraInsertsSingleBlock(v uint64) error {
	//FIXME add sanity checking.
	r.extraInsertsSingleBlock = v
	return nil
}
func (r *clientPut) SetExtraInsertsSplitfileHeaderBlock(v uint64) error {
	//FIXME add sanity checking.
	r.extraInsertsSplitfileHeaderBlock = v
	return nil
}
func (r *clientPut) SetCompatibilityMode(v string) error {
	//FIXME add sanity checking.
	r.compatibilityMode = v
	return nil
}
func (r *clientPut) SetLocalRequestOnly(v bool) error {
	//FIXME add sanity checking.
	r.localRequestOnly = v
	return nil
}
func (r *clientPut) SetOverrideSplitfileCryptoKey(v string) error {
	//FIXME add sanity checking.
	r.overrideSplitfileCryptoKey = v
	return nil
}
func (r *clientPut) SetRealTimeFlag(v bool) error {
	//FIXME add sanity checking.
	r.realTimeFlag = v
	return nil
}
func (r *clientPut) SetMetadataThreshold(v int64) error {
	//FIXME add sanity checking.
	r.metadataThreshold = v
	return nil
}
func (r *clientPut) SetData(v *[]byte) error {
	//FIXME add sanity checking.
	r.data = v
	return nil
}

func (r *clientPut) getMessage() message {

	params := []string{}
	if r.uRI != "" {
		params = append(params, "URI="+r.uRI)
	}
	if r.contentType != "" {
		params = append(params, "ContentType="+r.contentType)
	}
	if r.identifier != "" {
		params = append(params, "Identifier="+r.identifier)
	}
	if r.verbosity != 0 {
		params = append(params, "Verbosity="+strconv.FormatInt(int64(r.verbosity), 10))
	}
	if r.maxRetries != 0 {
		params = append(params, "MaxRetries="+strconv.FormatInt(r.maxRetries, 10))
	}
	if r.priorityClass != 2 {
		params = append(params, "PriorityClass="+strconv.FormatInt(r.priorityClass, 10))
	}
	if r.getCHKOnly {
		params = append(params, "GetCHKOnly=true")
	}
	if r.global {
		params = append(params, "Global=true")
	}
	if r.dontCompress {
		params = append(params, "DontCompress=true")
	}
	if len(r.codecs) > 0 {
		codecsList := ""
		for codec := range r.codecs {
			codecsList += r.codecs[codec] + ","
		}
		params = append(params, "Codecs="+codecsList[:len(codecsList)-1])
	}
	if r.clientToken != "" {
		params = append(params, "ClientToken="+r.clientToken)
	}
	if r.persistence != "" {
		params = append(params, "Persistence="+r.persistence)
	}
	if r.targetFilename != "" {
		params = append(params, "TargetFilename="+r.targetFilename)
	}
	if r.earlyEncode {
		params = append(params, "EarlyEncode=true")
	}
	if r.uploadFrom != "direct" {
		params = append(params, "UploadFrom="+r.uploadFrom)
	}
	if r.uploadFrom == "direct" {
		params = append(params, "DataLength="+string(r.dataLength))
	}
	if r.uploadFrom == "disk" {
		params = append(params, "Filename="+r.filename)
	}
	if r.uploadFrom == "redirect" {
		params = append(params, "TargetURI="+r.targetURI)
	}
	if r.fileHash != "" {
		params = append(params, "FileHash="+r.fileHash)
	}
	if r.binaryBlob {
		params = append(params, "BinaryBlob=true")
	}
	if !r.forkOnCacheable {
		params = append(params, "ForkOnCacheable=false")
	}
	if r.extraInsertsSingleBlock != 2 {
		params = append(params, "ExtraInsertsSingleBlock="+strconv.FormatUint(r.extraInsertsSingleBlock, 10))
	}
	if r.extraInsertsSplitfileHeaderBlock != 2 {
		params = append(params, "ExtraInsertsSplitfileHeaderBlock="+string(r.extraInsertsSplitfileHeaderBlock))
	}
	if r.compatibilityMode != "COMPAT_CURRENT" {
		params = append(params, "CompatibilityMode="+r.compatibilityMode)
	}
	if r.localRequestOnly {
		params = append(params, "LocalRequestOnly=true")
	}
	if r.overrideSplitfileCryptoKey != "" {
		params = append(params, "OverrideSplitfileCryptoKey="+r.overrideSplitfileCryptoKey)
	}
	if r.realTimeFlag {
		params = append(params, "RealTimeFlag=true")
	}
	if r.metadataThreshold != -1 {
		params = append(params, "MetadataThreshold="+strconv.FormatInt(r.metadataThreshold, 10))
	}

	return message{"ClientPut",
		params,
		r.data}
}

func (r *FCPClient) DoInsert(v clientPut) error {
	r.msgSender <- v.getMessage()
	return nil
}
