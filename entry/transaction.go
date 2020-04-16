package entry

import (
	bytomsdk "github.com/vapor-sdk/bytom"
	vaporsdk "github.com/vapor-sdk/vapor"
)

// DecodeRawTx decode raw transaction for bytom and vapor
func DecodeRawTx(chainName, rawTransaction string) []byte {
	switch chainName {
	case "bytom":
		return bytomsdk.BytomDecodeRawTx(rawTransaction)
	case "vapor":
		return vaporsdk.VaporDecodeRawTx(rawTransaction)
	default:
		return nil
	}
}
