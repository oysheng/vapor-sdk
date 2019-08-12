package util

// Transaction is the annotated transaction
type Transaction struct {
	TxID      string            `json:"hash"`
	Version   int64             `json:"version"`
	Size      int64             `json:"size"`
	TimeRange int64             `json:"time_range"`
	Inputs    []AnnotatedInput  `json:"inputs"`
	Outputs   []AnnotatedOutput `json:"outputs"`
	Fee       int64             `json:"fee"`
}

// AnnotatedInput means an annotated transaction input.
type AnnotatedInput struct {
	Type             string   `json:"type"`
	InputID          string   `json:"input_id"`
	AssetID          string   `json:"asset"`
	Amount           int64    `json:"amount"`
	ControlProgram   string   `json:"script,omitempty"`
	Address          string   `json:"address,omitempty"`
	IssuanceProgram  string   `json:"issuance_program,omitempty"`
	AssetDefinition  string   `json:"asset_definition,omitempty"`
	SpentOutputID    string   `json:"spent_output_id,omitempty"`
	Arbitrary        string   `json:"arbitrary,omitempty"`
	WitnessArguments []string `json:"arguments,omitempty"`
	Vote             string   `json:"vote,omitempty"`
	SignData         string   `json:"sign_data,omitempty"`
}

// AnnotatedOutput means an annotated transaction output.
type AnnotatedOutput struct {
	Type           string `json:"type"`
	OutputID       string `json:"utxo_id"`
	Position       int    `json:"position"`
	AssetID        string `json:"asset"`
	Amount         int64  `json:"amount"`
	ControlProgram string `json:"script"`
	Address        string `json:"address,omitempty"`
	Vote           string `json:"vote,omitempty"`
}
