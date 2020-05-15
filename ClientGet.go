package fcp

import (
	//	"fmt"
	"strconv"
)

type clientGet struct { //				Values					Mandatory	default		Criteria 			Description
	ignoreDS   bool   //											no			false							Do we ignore the datastore?
	dSonly     bool   //											no			false							Check only in our local datastore for the file i.e. don't ask other nodes if they have the file. (~= htl 0)
	uRI        string //					freenet URI				yes											The URI of the freenet file you want to download e.g. KSK@sample.txt, CHK@zfwLW...Dvs,AAMC--8/
	identifier string //					Any text					yes											A string to uniquely identify to the client the file you are receiving.
	verbosity  byte   //					bitmask					no			0
	/*
	   0: report when complete
	   1: SimpleProgress messages
	   2: SendingToNetwork messages
	   4: CompatibilityMode messages (since 1255)
	   8: ExpectedHashes messages (since 1255)
	   32: ExpectedMIME messages (since 1307)
	   64: ExpectedDataLength messages (since 1307)
	*/
	maxSize       uint64 //				positive integer			no											Maximum size of returned data in bytes.
	maxTempSize   uint64 //				positive integer			no											Maximum size of intermediary data in bytes.
	maxRetries    int64  //				integer -1~?				no			0								Number of times the node will automatically retry to get the data. -1 means retry forever, and will use ULPRs to maintain the request efficiently.
	priorityClass int64  //										no			4								0 is maximum priority -> 1 V.High -> 2 High -> 3 Medium -> 4 Low -> 5 V.Low -> 6 Paused
	peristence    string //										no 			"connection"						Whether the download stays on the queue across new client "connection"s, Freenet "restart"s, or "forever"
	clientToken   string //				any string				no											Returned in PersistentGet, a hint to the client, so the client doesn't need to maintain its own state.
	global        bool   //										no			false							Whether the download is visible on the global queue or not.
	returnType    string //										no			"direct"
	/*
	   direct: return the data directly to the client via an AllData message, once we have all of it. (For persistent requests, the client will get a DataFound message but must send a GetRequestStatus to ask for the AllData).
	   none: don't return the data at all, just fetch it to the node and tell the client when we have finished.
	   disk: write the data to disk. If you download to disk, you have to do a TestDDARequest.
	   In the future, chunked may also be supported (return it in segments as they are ready), but this is not yet implemented.
	*/
	binaryBlob                bool     //							no			false							If true, return the data blocks required to fetch this site as a binary blob (.fblob) file.
	filterData                bool     //							no			false							Whether to run the fetched content through the content filters.
	allowedMimeTypes          []string //	list of MIME types		no											If set, only allow certain MIME types in the returned data. If the data is of a MIME type which isn't listed, the request will fail with a WRONG_MIME_TYPE error (code 29) as soon as it realizes this.
	filename                  string   //	full path/filename		yes						ReturnType=Disk		Name and path of the file where the download is to be stored.
	tempFilename              string   //	full path/filename		no						ReturnType=Disk		Name and path of a temporary file where the partial download is to be stored.
	realTimeFlag              bool     //							no			false							Whether to fetch the data with the real-time flag set to realtime (true) or bulk (false). (since 1311)
	initialMetadataDataLength uint64   //							no			0								If nonzero, fetch from metadata instead of from a URI. The metadata will be after the EndMessage/Data, and will be the specified length. (since 1380)
}

func (r *FCPClient) NewClientGet() clientGet {
	return clientGet{
		false, false,
		"", "",
		0, 0, 0, 0, 4,
		"connection", "",
		false,
		"direct",
		false, false,
		[]string{},
		"", "",
		false,
		0,
	}
}

func (r *clientGet) SetIgnoreDS(v bool) error {
	//FIXME add sanity checking.
	r.ignoreDS = v
	return nil
}
func (r *clientGet) SetDSonly(v bool) error {
	//FIXME add sanity checking.
	r.dSonly = v
	return nil
}
func (r *clientGet) SetURI(v string) error {
	//FIXME add sanity checking.
	r.uRI = v
	return nil
}
func (r *clientGet) SetIdentifier(v string) error {
	//FIXME add sanity checking.
	r.identifier = v
	return nil
}

/*Bitmask...
 0: report when complete
 1: SimpleProgress messages
 2: SendingToNetwork messages
 4: CompatibilityMode messages (since 1255)
 8: ExpectedHashes messages (since 1255)
32: ExpectedMIME messages (since 1307)
64: ExpectedDataLength messages (since 1307)*/
func (r *clientGet) SetVerbosity(v byte) error {
	//FIXME add sanity checking.
	r.verbosity = v
	return nil
}
func (r *clientGet) SetMaxSize(v uint64) error {
	//FIXME add sanity checking.
	r.maxSize = v
	return nil
}
func (r *clientGet) SetMaxTempSize(v uint64) error {
	//FIXME add sanity checking.
	r.maxTempSize = v
	return nil
}
func (r *clientGet) SetMaxRetries(v int64) error {
	//FIXME add sanity checking.
	r.maxRetries = v
	return nil
}
func (r *clientGet) SetPriorityClass(v int64) error {
	//FIXME add sanity checking.
	r.priorityClass = v
	return nil
}
func (r *clientGet) SetPersistence(v string) error {
	//FIXME add sanity checking.
	r.peristence = v
	return nil
}
func (r *clientGet) SetClientToken(v string) error {
	//FIXME add sanity checking.
	r.clientToken = v
	return nil
}
func (r *clientGet) SetGlobal(v bool) error {
	//FIXME add sanity checking.
	r.global = v
	return nil
}
func (r *clientGet) SetReturnType(v string) error {
	//FIXME add sanity checking.
	r.returnType = v
	return nil
}
func (r *clientGet) SetBinaryBlob(v bool) error {
	//FIXME add sanity checking.
	r.binaryBlob = v
	return nil
}
func (r *clientGet) SetFilterData(v bool) error {
	//FIXME add sanity checking.
	r.filterData = v
	return nil
}
func (r *clientGet) SetAllowedMimeTypes(v []string) error { //FIXME add sanity checking.
	r.allowedMimeTypes = v
	return nil
}
func (r *clientGet) SetFilename(v string) error {
	//FIXME add sanity checking.
	r.filename = v
	return nil
}
func (r *clientGet) SetTempFilename(v string) error {
	//FIXME add sanity checking.
	r.tempFilename = v
	return nil
}
func (r *clientGet) SetRealTimeFlag(v bool) error {
	//FIXME add sanity checking.
	r.realTimeFlag = v
	return nil
}
func (r *clientGet) SetInitialMetadataDataLength(v uint64) error {
	//FIXME add sanity checking.
	r.initialMetadataDataLength = v
	return nil
}

func (r *clientGet) getMessage() message {

	params := []string{}
	if r.ignoreDS {
		params = append(params, "IgnoreDS=true")
	}
	if r.dSonly {
		params = append(params, "DSonly=true")
	}
	if r.uRI != "" {
		params = append(params, "URI="+r.uRI)
	}
	if r.identifier != "" {
		params = append(params, "Identifier="+r.identifier)
	}
	if r.verbosity != 0 {
		params = append(params, "Verbosity="+strconv.FormatInt(int64(r.verbosity), 10))
	}
	if r.maxSize != 0 {
		params = append(params, "MaxSize="+strconv.FormatUint(r.maxSize, 10))
	}
	if r.maxTempSize != 0 {
		params = append(params, "MaxTempSize="+strconv.FormatUint(r.maxTempSize, 10))
	}
	if r.maxRetries != 0 {
		params = append(params, "MaxRetries="+strconv.FormatInt(r.maxRetries, 10))
	}
	if r.priorityClass != 4 {
		params = append(params, "PriorityClass="+strconv.FormatInt(r.priorityClass, 10))
	}
	if r.peristence != "connection" {
		params = append(params, "Persistence="+r.peristence)
	}
	if r.clientToken != "" {
		params = append(params, "ClientToken="+r.clientToken)
	}
	if r.global {
		params = append(params, "Global=true")
	}
	if r.returnType != "direct" {
		params = append(params, "ReturnType="+r.returnType)
	}
	if r.binaryBlob {
		params = append(params, "BinaryBlob=true")
	}
	if r.filterData {
		params = append(params, "FilterData=true")
	}
	if len(r.allowedMimeTypes) > 0 {
		mimeList := ""
		for mimeType := range r.allowedMimeTypes {
			mimeList += r.allowedMimeTypes[mimeType] + ","
		}
		params = append(params, "AllowedMimeTypes="+mimeList[:len(mimeList)-1])
	}

	if r.returnType == "disk" {
		params = append(params, "Filename="+r.filename)
		if r.tempFilename != "" {
			params = append(params, "TempFilename="+r.tempFilename)
		}
	}
	if r.realTimeFlag {
		params = append(params, "RealTimeFlag=true")
	}
	if r.initialMetadataDataLength > 0 {
		params = append(params, "InitialMetadata.DataLength="+strconv.FormatUint(r.initialMetadataDataLength, 10))
	}

	return message{"ClientGet",
		params,
		nil}
}

func (r *FCPClient) DoFetch(v clientGet) error {
	r.msgSender <- v.getMessage()
	return nil
}
