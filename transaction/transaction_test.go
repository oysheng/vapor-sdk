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
						SignData:         "96b1454d0ca5fd05f321345149ab526ad14be9ae364fdb6e6bda5825b4e1c388",
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
						SignData:         "6e142176d423e825f27971c928ca09e174fc7e8134428c19b3a33e7d6a7abfac",
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
		{
			rawTransaction: `07010001016401628a3e00e2f6cfe2765fd0b51201d3d5e44ba461aa3cd57306068b7bdf0d4a105dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8080b0e8d3eb94d5010101160014973616e27ba7468f3a54820c97ab1b22094bd42d630240d8f36726bf7e69a01afdf05251a2338fb8c2595d881898b5903302d32619185f41c90990e7160593fd4dc416fb38b3845f32277685028e52f01fa98a4d121a0720fbbb8233f1435c2c0ab26ee4aeb94e534490c65a48e253a5dc64cad835462d290201430041ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8080c7e4a28dfed401011600140bcc5b6e8f2cb3390cf6d45fca37ed86062536010001820102409742a39a0bcfb5b7ac8f56f1894fbb694b53ebf58f9a032c36cc22d57a06e49e94ff7199063fb7a78190624fa3530f611404b56fc9af91dcaf4639614512cb643fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8080e983b1de16011600143eb3371ee17bfa7d1e6af07c2e1fc08b3b1177ad00`,
			wantTx: &Transaction{
				TxID:      "4b08a9a705bc83aa4015f7682d054603e6d335a39cdee27baba23681014ce5dd",
				Version:   1,
				Size:      411,
				TimeRange: 0,
				Inputs: []annotatedInput{
					annotatedInput{
						Type:           "spend",
						InputID:        "645045dc9e8bee31738f1d30f702e6678533e215c9175920b3582db9d8026eeb",
						AssetID:        "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
						Amount:         120000000000000000,
						ControlProgram: "0014973616e27ba7468f3a54820c97ab1b22094bd42d",
						Address:        "vp1qjumpdcnm5arg7wj5sgxf02cmygy5h4pde4aynj",
						SpentOutputID:  "fcf9d0fae86697cd396d81a60cbd296f74ba337d76240d12f7baf3f1e548f771",
						WitnessArguments: []string{
							"d8f36726bf7e69a01afdf05251a2338fb8c2595d881898b5903302d32619185f41c90990e7160593fd4dc416fb38b3845f32277685028e52f01fa98a4d121a07",
							"fbbb8233f1435c2c0ab26ee4aeb94e534490c65a48e253a5dc64cad835462d29",
						},
						SignData: "ddd8e2eb9290b4ff95777a823c3193655e16314b37037145369768dc000fe9b8",
					},
				},
				Outputs: []annotatedOutput{
					annotatedOutput{
						Type:           "control",
						OutputID:       "7f78999d9e8c99a0f4ad763da4475c45bd920153a0e02a38b3b9dfdfac4f84a3",
						Position:       0,
						AssetID:        "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
						Amount:         119900000000000000,
						ControlProgram: "00140bcc5b6e8f2cb3390cf6d45fca37ed8606253601",
						Address:        "vp1qp0x9km509jenjr8k630u5dldscrz2dsp8vafzs",
					},
					annotatedOutput{
						Type:           "vote",
						OutputID:       "3b4b503d394a598c88268eba18626d6868dbbf66abd98485573f7f334a3d0124",
						Position:       1,
						AssetID:        "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
						Amount:         100000000000000,
						ControlProgram: "00143eb3371ee17bfa7d1e6af07c2e1fc08b3b1177ad",
						Address:        "vp1q86enw8hp00a868n27p7zu87q3va3zaady6805f",
						Vote:           "9742a39a0bcfb5b7ac8f56f1894fbb694b53ebf58f9a032c36cc22d57a06e49e94ff7199063fb7a78190624fa3530f611404b56fc9af91dcaf4639614512cb64",
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
