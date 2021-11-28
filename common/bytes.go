package common

import "encoding/hex"

func FromHex(hexValue string) (bytes []byte, err error) {
	return hex.DecodeString(hexValue)
}
