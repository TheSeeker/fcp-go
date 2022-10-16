// Peer
package fcp

import (
	"fmt"
	//	"strconv"
	"time"
)

/*
	Peer
	lastGoodVersion=Fred,0.7,1.0,1134
	opennet=false
	seed=false
	myName=Somebody's Node
	identity=-~AH0...
	location=0.6699208187881001
	testnet=false
	version=Fred,0.7,1.0,1135
	physical.udp=88.x.x.x:xxxxx
	ark.pubURI=SSK@rsj1.../ark
	ark.number=40
	dsaPubKey.y=Y3DQ...
	dsaGroup.p=AIYI...
	dsaGroup.g=UaRa...
	dsaGroup.q=ALFD...
	auth.negTypes=2
	EndMessage

Note

When you specifiied WithMetadata=true� and/or WithVolatile=true� in your ListPeer or ListPeers command the output will contain more information about the peer. WithVolatile=true� will add the following lines to the output:

	volatile.averagePingTime=1172.3187327786154
	volatile.overloadProbability=77.31471256750754
	volatile.percentTimeRoutableConnection=18.556701030927837
	volatile.routingBackoffPercent=10.194888893854625
	volatile.status=CONNECTED
	volatile.totalBytesIn=29870
	volatile.routingBackoffLength=2000
	volatile.lastRoutingBackoffReason=ForwardRejectedOverload
	volatile.routingBackoff=0
	volatile.totalBytesOut=51821

WithMetadata=true� adds the following lines:

	metadata.routableConnectionCheckCount=125
	metadata.hadRoutableConnectionCount=46
	metadata.timeLastConnected=1207738143843
	metadata.timeLastSuccess=1207738140155
	metadata.timeLastRoutable=1207738143843
	metadata.timeLastReceivedPacket=1207738141513
	metadata.detected.udp=x.x.x.x:xxxxx

*/

type volatile struct {
	averagePingTime               float64
	overloadProbability           float64
	percentTimeRoutableConnection float64
	routingBackoffPercent         float64
	status                        string
	totalBytesIn                  int64
	routingBackoffLength          int64
	lastRoutingBackoffReason      string
	routingBackoff                int64
	totalBytesOut                 int64
}

type metadata struct {
	routableConnectionCheckCount int64
	hadRoutableConnectionCount   int64
	timeLastConnected            time.Time
	timeLastSuccess              time.Time
	timeLastRoutable             time.Time
	timeLastReceivedPacket       time.Time
	detected_udp                 string
}

type Peer struct {
	lastGoodVersion []string
	opennet         bool
	seed            bool
	myName          string
	identity        string
	location        float64
	testnet         bool
	version         []string
	physical_udp    string
	ark_pubURI      string
	ark_number      int64
	dsaPubKey_y     string
	dsaGroup_p      string
	dsaGroup_g      string
	dsaGroup_q      string
	auth_negTypes   int64
	vol             volatile
	meta            metadata
}

func (r *Peer) NewMetadata() metadata {
	return metadata{
		0, 0,
		time.Unix(0, 0), time.Unix(0, 0), time.Unix(0, 0), time.Unix(0, 0),
		"",
	}
}

func (r *Peer) parseMessage(rawMsg []string) {
	fmt.Println("%+v", rawMsg)
	for _, v := range rawMsg {
		param, val := SplitParam(v)

		switch param {
		case "Peer":
		//ignore
		case "lastGoodVersion":
		case "opennet":
			r.opennet = (val == "true")
		case "seed":
			r.seed = (val == "true")
		case "myName":
		case "identity":
		case "location":
		case "testnet":
		case "version":
		case "physical.udp":
		case "ark.pubURI":
		case "ark.number":
		case "dsaPubKey.y":
		case "dsaGroup.p":
		case "dsaGroup.g":
		case "dsaGroup.q":
		case "auth.negTypes":
		case "EndMessage":
			break
		default:
			err := fmt.Errorf("unknown parameter in Peer: %+v with value of %+v", param, val)
			fmt.Println(err.Error())
		}
	}
}

func (r *Peer) GetName() string { return "Peer" }
