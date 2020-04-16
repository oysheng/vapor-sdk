package entry

import (
	"github.com/vapor-sdk/bytom"
	"github.com/vapor-sdk/vapor"
)

// DecodeRawTx decode raw transaction for bytom and vapor
func DecodeRawTx(chainName, rawTransaction string) []byte {
	switch chainName {
	case "bytom":
		return bytom.BytomDecodeRawTx(rawTransaction)
	case "vapor":
		return vapor.VaporDecodeRawTx(rawTransaction)
	default:
		return nil
	}
}
