package transaction

import (
	"encoding/hex"
	"encoding/json"

	"github.com/bytom/bytom/blockchain/txbuilder"
	"github.com/bytom/bytom/common"
	"github.com/bytom/bytom/consensus"
	"github.com/bytom/bytom/consensus/segwit"
	"github.com/bytom/bytom/protocol/bc"
	"github.com/bytom/bytom/protocol/bc/types"
	"github.com/bytom/bytom/protocol/vm/vmutil"

	"github.com/vapor-sdk/util"
)

// BytomDecodeRawTx decode raw transaction
func BytomDecodeRawTx(rawTransaction string) []byte {
	var rawTx types.Tx
	if err := rawTx.UnmarshalText([]byte(rawTransaction)); err != nil {
		return nil
	}

	tx := &util.Transaction{
		TxID:      rawTx.ID.String(),
		Version:   int64(rawTx.Version),
		Size:      int64(rawTx.SerializedSize),
		TimeRange: int64(rawTx.TimeRange),
		Inputs:    []util.AnnotatedInput{},
		Outputs:   []util.AnnotatedOutput{},
		Fee:       int64(txbuilder.CalculateTxFee(&rawTx)),
	}

	for i := range rawTx.Inputs {
		tx.Inputs = append(tx.Inputs, buildAnnotatedInput(&rawTx, uint32(i)))
	}
	for i := range rawTx.Outputs {
		tx.Outputs = append(tx.Outputs, buildAnnotatedOutput(&rawTx, i))
	}

	jsonTx, err := json.Marshal(tx)
	if err != nil {
		return nil
	}
	return jsonTx
}

// buildAnnotatedInput build the annotated input.
func buildAnnotatedInput(tx *types.Tx, i uint32) util.AnnotatedInput {
	orig := tx.Inputs[i]
	in := util.AnnotatedInput{}
	if orig.InputType() != types.CoinbaseInputType {
		assetID := orig.AssetID()
		in.AssetID = assetID.String()
		in.Amount = int64(orig.Amount())
		signData := tx.SigHash(i)
		in.SignData = signData.String()
	} else {
		in.AssetID = consensus.BTMAssetID.String()
	}

	id := tx.Tx.InputIDs[i]
	in.InputID = id.String()
	e := tx.Entries[id]
	switch e := e.(type) {
	case *bc.Spend:
		in.Type = "spend"
		controlProgram := orig.ControlProgram()
		in.ControlProgram = hex.EncodeToString(controlProgram)
		in.Address = getAddressFromControlProgram(controlProgram)
		in.SpentOutputID = e.SpentOutputId.String()
		arguments := orig.Arguments()
		for _, arg := range arguments {
			in.WitnessArguments = append(in.WitnessArguments, hex.EncodeToString(arg))
		}

	case *bc.Issuance:
		in.Type = "issue"
		issuanceProgram := orig.IssuanceProgram()
		in.IssuanceProgram = hex.EncodeToString(issuanceProgram)
		arguments := orig.Arguments()
		for _, arg := range arguments {
			in.WitnessArguments = append(in.WitnessArguments, hex.EncodeToString(arg))
		}
		if assetDefinition := orig.AssetDefinition(); isValidJSON(assetDefinition) {
			in.AssetDefinition = hex.EncodeToString(assetDefinition)
		}

	case *bc.Coinbase:
		in.Type = "coinbase"
		in.Arbitrary = hex.EncodeToString(e.Arbitrary)
	}
	return in
}

// buildAnnotatedOutput build the annotated output.
func buildAnnotatedOutput(tx *types.Tx, idx int) util.AnnotatedOutput {
	orig := tx.Outputs[idx]
	outid := tx.OutputID(idx)
	out := util.AnnotatedOutput{
		OutputID:       outid.String(),
		Position:       idx,
		AssetID:        orig.AssetId.String(),
		Amount:         int64(orig.Amount),
		ControlProgram: hex.EncodeToString(orig.ControlProgram),
		Address:        getAddressFromControlProgram(orig.ControlProgram),
	}

	if vmutil.IsUnspendable(orig.ControlProgram) {
		out.Type = "retire"
	} else {
		out.Type = "control"
	}
	return out
}

func isValidJSON(b []byte) bool {
	var v interface{}
	err := json.Unmarshal(b, &v)
	return err == nil
}

func getAddressFromControlProgram(prog []byte) string {
	if segwit.IsP2WPKHScript(prog) {
		if pubHash, err := segwit.GetHashFromStandardProg(prog); err == nil {
			return buildP2PKHAddress(pubHash)
		}
	} else if segwit.IsP2WSHScript(prog) {
		if scriptHash, err := segwit.GetHashFromStandardProg(prog); err == nil {
			return buildP2SHAddress(scriptHash)
		}
	}
	return ""
}

func buildP2PKHAddress(pubHash []byte) string {
	address, err := common.NewAddressWitnessPubKeyHash(pubHash, &consensus.MainNetParams)
	if err != nil {
		return ""
	}
	return address.EncodeAddress()
}

func buildP2SHAddress(scriptHash []byte) string {
	address, err := common.NewAddressWitnessScriptHash(scriptHash, &consensus.MainNetParams)
	if err != nil {
		return ""
	}
	return address.EncodeAddress()
}
