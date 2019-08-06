package transaction

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/vapor/testutil"
)

func TestDecodeRawTransaction(t *testing.T) {
	cases := []struct {
		desc           string
		rawTransaction string
		wantTx         *Transaction
	}{
		{
			rawTransaction: `07010001015f015d13c41cc617304ba0866fa59f07d7bb2bcab60c43e5cc79bb75a4dd97471cdcbabb16babcc936f9a7467bc9f615be17cb69809aa7cefd4287d4098690585b399180ade20400011600144b6995dc11354d44c6e382c19d6b92bdbbd3aea1010002013e003cbb16babcc936f9a7467bc9f615be17cb69809aa7cefd4287d4098690585b3991c096b102011600149682e64b2114f7c2581ab1ba0c67315d06aaea8200013e003cbb16babcc936f9a7467bc9f615be17cb69809aa7cefd4287d4098690585b3991c096b10201160014da26416fa79947ec6a569e0493dbffec1a3f223400`,
			wantTx: &Transaction{
				TxID:      "ab3130d01b1d41c4d772f258fc5d2b38660d5d41e44d107fe935eb2b85015990",
				Version:   1,
				Size:      234,
				TimeRange: 0,
				Inputs: []annotatedInput{
					annotatedInput{
						Type:             "spend",
						InputID:          "7d9c3a6481fd249c78f6037d3c999a3fe753882bd13e072cecc8ce92fbbbb41b",
						AssetID:          "bb16babcc936f9a7467bc9f615be17cb69809aa7cefd4287d4098690585b3991",
						Amount:           10000000,
						ControlProgram:   "00144b6995dc11354d44c6e382c19d6b92bdbbd3aea1",
						Address:          "vp1qfd5ethq3x4x5f3hrstqe66ujhkaa8t4p8vud4p",
						SpentOutputID:    "873cd20c2cd260e1d2902f173bbc32490a9aa184b8e47aaedf3f37d7bf5225dd",
						Arbitrary:        "",
						WitnessArguments: nil,
					},
				},
				Outputs: []annotatedOutput{
					annotatedOutput{
						Type:           "control",
						OutputID:       "1df78ee679f30bb4597e1c3e459a0cd0429e69de875dd85a57fa34f94a59aba4",
						Position:       0,
						AssetID:        "bb16babcc936f9a7467bc9f615be17cb69809aa7cefd4287d4098690585b3991",
						Amount:         5000000,
						ControlProgram: "00149682e64b2114f7c2581ab1ba0c67315d06aaea82",
						Address:        "vp1qj6pwvjepznmuykq6kxaqcee3t5r2465z0hmr70",
						Vote:           "",
					},
					annotatedOutput{
						Type:           "control",
						OutputID:       "6d1116c361d8001e5f2d491c796eb36b5cb53dd630c1310a9be13742fd6e9cbc",
						Position:       1,
						AssetID:        "bb16babcc936f9a7467bc9f615be17cb69809aa7cefd4287d4098690585b3991",
						Amount:         5000000,
						ControlProgram: "0014da26416fa79947ec6a569e0493dbffec1a3f2234",
						Address:        "vp1qmgnyzma8n9r7c6jknczf8kllasdr7g35whjwpg",
						Vote:           "",
					},
				},
				Fee: 0,
			},
		},
		{
			rawTransaction: `07010001015d015bbfa8cb0c58b545bf844dd642b6b5333ac76b4b789b3795a129a93a9fe47c3227ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff904e0101160014d66216efa3177397973c6e173f8f7f17a7b64b81010001013c003affffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff904e01160014d66216efa3177397973c6e173f8f7f17a7b64b8100`,
			wantTx: &Transaction{
				TxID:      "8d0010cb3cd757d6c2dbe0864f18c0651c9c1cbdc4ca68219481b182aef47527",
				Version:   1,
				Size:      165,
				TimeRange: 0,
				Inputs: []annotatedInput{
					annotatedInput{
						Type:             "spend",
						InputID:          "e8bed028eadf67a683f9a4ccfac2bbd385b0e5abf8d30a9f8d4f1b814d536402",
						AssetID:          "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
						Amount:           10000,
						ControlProgram:   "0014d66216efa3177397973c6e173f8f7f17a7b64b81",
						Address:          "vp1q6e3pdmarzaee09eudctnlrmlz7nmvjup8wtqxd",
						SpentOutputID:    "933d1e2e7a1317f25ee1f75de6abf93867100c4190a9e3d2c4abe3485ebe63b7",
						Arbitrary:        "",
						WitnessArguments: nil,
					},
				},
				Outputs: []annotatedOutput{
					annotatedOutput{
						Type:           "control",
						OutputID:       "a92271244e13fee0720385b27444fee7cffa51810d42999b8c9f01088fedf9c5",
						Position:       0,
						AssetID:        "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
						Amount:         10000,
						ControlProgram: "0014d66216efa3177397973c6e173f8f7f17a7b64b81",
						Address:        "vp1q6e3pdmarzaee09eudctnlrmlz7nmvjup8wtqxd",
						Vote:           "",
					},
				},
				Fee: 0,
			},
		},
	}

	for i, c := range cases {
		jsonTx := DecodeRawTransaction(c.rawTransaction)
		if jsonTx == nil {
			t.Fatal(errors.New("error"))
		}

		gotTx := &Transaction{}
		if err := json.Unmarshal(jsonTx, gotTx); err != nil {
			t.Fatal(err)
		}

		if !testutil.DeepEqual(gotTx, c.wantTx) {
			t.Errorf("case #%d, annotated transaction got=%#v, want=%#v", i, gotTx, c.wantTx)
		}
	}
}
