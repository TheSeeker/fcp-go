//ExpectedHashes
package fcp

import (
	"fmt"
)

type ExpectedHashes struct {
	identifier    string
	global        bool
	hashes_SHA512 string
	hashes_SHA256 string
	hashes_MD5    string
	hashes_SHA1   string
	hashes_TTH    string
	hashes_ED2K   string
}

func (r *ExpectedHashes) parseMessage(rawMsg []string) {
	fmt.Println("%+v", rawMsg)
	for _, v := range rawMsg {
		param, val := SplitParam(v)

		switch param {
		case "ExpectedHashes":
		//ignore
		case "Identifier":
			r.identifier = val
		case "Global":
			r.global = (val == "true")
		case "Hashes.SHA512":
			r.hashes_SHA512 = val
		case "Hashes.SHA256":
			r.hashes_SHA256 = val
		case "Hashes.MD5":
			r.hashes_MD5 = val
		case "Hashes.SHA1":
			r.hashes_SHA1 = val
		case "Hashes.TTH":
			r.hashes_TTH = val
		case "Hashes.ED2K":
			r.hashes_ED2K = val
		case "EndMessage":
			break
		default:
			err := fmt.Errorf("unknown parameter in ExpectedHashes: %+v with value of %+v", param, val)
			fmt.Println(err.Error())
		}
	}
}

func (r *ExpectedHashes) GetIdentifier() string    { return r.identifier }
func (r *ExpectedHashes) GetGlobal() bool          { return r.global }
func (r *ExpectedHashes) GetHashes_SHA512() string { return r.hashes_SHA512 }
func (r *ExpectedHashes) GetHashes_SHA256() string { return r.hashes_SHA256 }
func (r *ExpectedHashes) GetHashes_SHAMD5() string { return r.hashes_MD5 }
func (r *ExpectedHashes) GetHashes_SHA1() string   { return r.hashes_SHA1 }
func (r *ExpectedHashes) GetHashes_TTH() string    { return r.hashes_TTH }
func (r *ExpectedHashes) GetHashes_ED2K() string   { return r.hashes_ED2K }
