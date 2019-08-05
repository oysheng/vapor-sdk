package transaction

import (
	"testing"

	"github.com/bytom/testutil"
	chainjson "github.com/vapor/encoding/json"
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
						ControlProgram:   chainjson.HexBytes{0x0, 0x14, 0x4b, 0x69, 0x95, 0xdc, 0x11, 0x35, 0x4d, 0x44, 0xc6, 0xe3, 0x82, 0xc1, 0x9d, 0x6b, 0x92, 0xbd, 0xbb, 0xd3, 0xae, 0xa1},
						Address:          "vp1qfd5ethq3x4x5f3hrstqe66ujhkaa8t4p8vud4p",
						SpentOutputID:    "873cd20c2cd260e1d2902f173bbc32490a9aa184b8e47aaedf3f37d7bf5225dd",
						Arbitrary:        nil,
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
						ControlProgram: chainjson.HexBytes{0x0, 0x14, 0x96, 0x82, 0xe6, 0x4b, 0x21, 0x14, 0xf7, 0xc2, 0x58, 0x1a, 0xb1, 0xba, 0xc, 0x67, 0x31, 0x5d, 0x6, 0xaa, 0xea, 0x82},
						Address:        "vp1qj6pwvjepznmuykq6kxaqcee3t5r2465z0hmr70",
						Vote:           nil,
					},
					annotatedOutput{
						Type:           "control",
						OutputID:       "6d1116c361d8001e5f2d491c796eb36b5cb53dd630c1310a9be13742fd6e9cbc",
						Position:       1,
						AssetID:        "bb16babcc936f9a7467bc9f615be17cb69809aa7cefd4287d4098690585b3991",
						Amount:         5000000,
						ControlProgram: chainjson.HexBytes{0x0, 0x14, 0xda, 0x26, 0x41, 0x6f, 0xa7, 0x99, 0x47, 0xec, 0x6a, 0x56, 0x9e, 0x4, 0x93, 0xdb, 0xff, 0xec, 0x1a, 0x3f, 0x22, 0x34},
						Address:        "vp1qmgnyzma8n9r7c6jknczf8kllasdr7g35whjwpg",
						Vote:           nil,
					},
				},
				Fee: 0,
			},
		},
	}

	for i, c := range cases {
		gotTx, err := DecodeRawTransaction(c.rawTransaction)
		if err != nil {
			t.Fatal(err)
		}

		if !testutil.DeepEqual(gotTx, c.wantTx) {
			t.Errorf(`case #%d, annotated transaction got=%#v;\n want=%#v`, i, gotTx, c.wantTx)
		}
	}
}
