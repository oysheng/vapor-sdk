package transaction

import (
	"encoding/hex"
	"encoding/json"

	"github.com/bytom/vapor/common"
	"github.com/bytom/vapor/common/arithmetic"
	"github.com/bytom/vapor/consensus"
	"github.com/bytom/vapor/consensus/segwit"
	"github.com/bytom/vapor/protocol/bc"
	"github.com/bytom/vapor/protocol/bc/types"

	"github.com/vapor-sdk/util"
)

// VaporDecodeRawTx decode raw transaction
func VaporDecodeRawTx(rawTransaction string) []byte {
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
	}

	for i := range rawTx.Inputs {
		tx.Inputs = append(tx.Inputs, buildAnnotatedInput(&rawTx, uint32(i)))
	}
	for i := range rawTx.Outputs {
		tx.Outputs = append(tx.Outputs, buildAnnotatedOutput(&rawTx, i))
	}

	txFee, err := arithmetic.CalculateTxFee(&rawTx)
	if err != nil {
		return nil
	}
	tx.Fee = int64(txFee)
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
		if vetoInput, ok := orig.TypedInput.(*types.VetoInput); ok {
			in.Vote = hex.EncodeToString(vetoInput.Vote)
		}
	} else {
		in.AssetID = consensus.BTMAssetID.String()
	}

	id := tx.Tx.InputIDs[i]
	in.InputID = id.String()
	e := tx.Entries[id]
	switch e := e.(type) {
	case *bc.VetoInput:
		in.Type = "veto"
		controlProgram := orig.ControlProgram()
		in.ControlProgram = hex.EncodeToString(controlProgram)
		in.Address = getAddressFromControlProgram(controlProgram, false)
		in.SpentOutputID = e.SpentOutputId.String()
		arguments := orig.Arguments()
		for _, arg := range arguments {
			in.WitnessArguments = append(in.WitnessArguments, hex.EncodeToString(arg))
		}

	case *bc.CrossChainInput:
		in.Type = "cross_chain_in"
		controlProgram := orig.ControlProgram()
		in.ControlProgram = hex.EncodeToString(controlProgram)
		in.Address = getAddressFromControlProgram(controlProgram, true)
		in.SpentOutputID = e.MainchainOutputId.String()
		arguments := orig.Arguments()
		for _, arg := range arguments {
			in.WitnessArguments = append(in.WitnessArguments, hex.EncodeToString(arg))
		}

	case *bc.Spend:
		in.Type = "spend"
		controlProgram := orig.ControlProgram()
		in.ControlProgram = hex.EncodeToString(controlProgram)
		in.Address = getAddressFromControlProgram(controlProgram, false)
		in.SpentOutputID = e.SpentOutputId.String()
		arguments := orig.Arguments()
		for _, arg := range arguments {
			in.WitnessArguments = append(in.WitnessArguments, hex.EncodeToString(arg))
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
		AssetID:        orig.AssetAmount().AssetId.String(),
		Amount:         int64(orig.AssetAmount().Amount),
		ControlProgram: hex.EncodeToString(orig.ControlProgram()),
	}

	var isMainchainAddress bool
	switch e := tx.Entries[*outid].(type) {
	case *bc.IntraChainOutput:
		out.Type = "control"
		isMainchainAddress = false

	case *bc.CrossChainOutput:
		out.Type = "cross_chain_out"
		isMainchainAddress = true

	case *bc.VoteOutput:
		out.Type = "vote"
		out.Vote = hex.EncodeToString(e.Vote)
		isMainchainAddress = false
	}

	out.Address = getAddressFromControlProgram(orig.ControlProgram(), isMainchainAddress)
	return out
}

func getAddressFromControlProgram(prog []byte, isMainchain bool) string {
	netParams := &consensus.MainNetParams
	if isMainchain {
		netParams = consensus.BytomMainNetParams(&consensus.MainNetParams)
	}
	if segwit.IsP2WPKHScript(prog) {
		if pubHash, err := segwit.GetHashFromStandardProg(prog); err == nil {
			return buildP2PKHAddress(pubHash, netParams)
		}
	} else if segwit.IsP2WSHScript(prog) {
		if scriptHash, err := segwit.GetHashFromStandardProg(prog); err == nil {
			return buildP2SHAddress(scriptHash, netParams)
		}
	}
	return ""
}

func buildP2PKHAddress(pubHash []byte, netParams *consensus.Params) string {
	address, err := common.NewAddressWitnessPubKeyHash(pubHash, netParams)
	if err != nil {
		return ""
	}
	return address.EncodeAddress()
}

func buildP2SHAddress(scriptHash []byte, netParams *consensus.Params) string {
	address, err := common.NewAddressWitnessScriptHash(scriptHash, netParams)
	if err != nil {
		return ""
	}
	return address.EncodeAddress()
}
