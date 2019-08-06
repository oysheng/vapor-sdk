package transaction

import (
	"encoding/hex"

	"github.com/vapor/common"
	"github.com/vapor/common/arithmetic"
	"github.com/vapor/consensus"
	"github.com/vapor/consensus/segwit"
	"github.com/vapor/protocol/bc"
	"github.com/vapor/protocol/bc/types"
)

// Transaction is the annotated transaction
type Transaction struct {
	TxID      string            `json:"tx_id"`
	Version   uint64            `json:"version"`
	Size      uint64            `json:"size"`
	TimeRange uint64            `json:"time_range"`
	Inputs    []annotatedInput  `json:"inputs"`
	Outputs   []annotatedOutput `json:"outputs"`
	Fee       uint64            `json:"fee"`
}

//annotatedInput means an annotated transaction input.
type annotatedInput struct {
	Type             string   `json:"type"`
	InputID          string   `json:"input_id"`
	AssetID          string   `json:"asset_id"`
	Amount           uint64   `json:"amount"`
	ControlProgram   string   `json:"control_program,omitempty"`
	Address          string   `json:"address,omitempty"`
	SpentOutputID    string   `json:"spent_output_id,omitempty"`
	Arbitrary        string   `json:"arbitrary,omitempty"`
	WitnessArguments []string `json:"witness_arguments"`
}

//annotatedOutput means an annotated transaction output.
type annotatedOutput struct {
	Type           string `json:"type"`
	OutputID       string `json:"output_id"`
	Position       int    `json:"position"`
	AssetID        string `json:"asset_id"`
	Amount         uint64 `json:"amount"`
	ControlProgram string `json:"control_program"`
	Address        string `json:"address,omitempty"`
	Vote           string `json:"vote,omitempty"`
}

// DecodeRawTransaction decode raw transaction
func DecodeRawTransaction(rawTransaction string) (*Transaction, error) {
	var rawTx types.Tx
	if err := rawTx.UnmarshalText([]byte(rawTransaction)); err != nil {
		return nil, err
	}

	tx := &Transaction{
		TxID:      rawTx.ID.String(),
		Version:   rawTx.Version,
		Size:      rawTx.SerializedSize,
		TimeRange: rawTx.TimeRange,
		Inputs:    []annotatedInput{},
		Outputs:   []annotatedOutput{},
	}

	for i := range rawTx.Inputs {
		tx.Inputs = append(tx.Inputs, buildAnnotatedInput(&rawTx, uint32(i)))
	}
	for i := range rawTx.Outputs {
		tx.Outputs = append(tx.Outputs, buildAnnotatedOutput(&rawTx, i))
	}

	tx.Fee, _ = arithmetic.CalculateTxFee(&rawTx)
	return tx, nil
}

// buildAnnotatedInput build the annotated input.
func buildAnnotatedInput(tx *types.Tx, i uint32) annotatedInput {
	orig := tx.Inputs[i]
	in := annotatedInput{}
	if orig.InputType() != types.CoinbaseInputType {
		assetID := orig.AssetID()
		in.AssetID = assetID.String()
		in.Amount = orig.Amount()
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
func buildAnnotatedOutput(tx *types.Tx, idx int) annotatedOutput {
	orig := tx.Outputs[idx]
	outid := tx.OutputID(idx)
	out := annotatedOutput{
		OutputID:       outid.String(),
		Position:       idx,
		AssetID:        orig.AssetAmount().AssetId.String(),
		Amount:         orig.AssetAmount().Amount,
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
	netParams := &consensus.ActiveNetParams
	if isMainchain {
		netParams = consensus.BytomMainNetParams(&consensus.ActiveNetParams)
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
