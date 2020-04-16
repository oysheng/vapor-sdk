package transaction

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/bytom/bytom/testutil"

	"github.com/vapor-sdk/util"
)

func TestBytomDecodeRawTx(t *testing.T) {
	cases := []struct {
		desc           string
		rawTransaction string
		wantTx         *util.Transaction
	}{
		{
			rawTransaction: `070100010161015fc8215913a270d3d953ef431626b19a89adf38e2486bb235da732f0afed515299ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8099c4d59901000116001456ac170c7965eeac1cc34928c9f464e3f88c17d8630240b1e99a3590d7db80126b273088937a87ba1e8d2f91021a2fd2c36579f7713926e8c7b46c047a43933b008ff16ecc2eb8ee888b4ca1fe3fdf082824e0b3899b02202fb851c6ed665fcd9ebc259da1461a1e284ac3b27f5e86c84164aa518648222602013effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80bbd0ec980101160014c3d320e1dc4fe787e9f13c1464e3ea5aae96a58f00013cffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8084af5f01160014bb93cdb4eca74b068321eeb84ac5d33686281b6500`,
			wantTx: &util.Transaction{
				TxID:      "4c97d7412b04d49acc33762fc748cd0780d8b44086c229c1a6d0f2adfaaac2db",
				Version:   1,
				Size:      332,
				TimeRange: 0,
				Inputs: []util.AnnotatedInput{
					util.AnnotatedInput{
						Type:           "spend",
						InputID:        "9963265eb601df48501cc240e1480780e9ed6e0c8f18fd7dd57954068c5dfd02",
						AssetID:        "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
						Amount:         41250000000,
						ControlProgram: "001456ac170c7965eeac1cc34928c9f464e3f88c17d8",
						Address:        "bm1q26kpwrrevhh2c8xrfy5vnaryu0ugc97c3j896t",
						SpentOutputID:  "01bb3309666618a1507cb5be845b17dee5eb8028ee7e71b17d74b4dc97085bc8",
						WitnessArguments: []string{
							"b1e99a3590d7db80126b273088937a87ba1e8d2f91021a2fd2c36579f7713926e8c7b46c047a43933b008ff16ecc2eb8ee888b4ca1fe3fdf082824e0b3899b02",
							"2fb851c6ed665fcd9ebc259da1461a1e284ac3b27f5e86c84164aa5186482226",
						},
						SignData: "8d2bb534c819464472a94b41cea788e97a2c9dae09a6cb3b7024a44ce5a27835",
					},
				},
				Outputs: []util.AnnotatedOutput{
					util.AnnotatedOutput{
						Type:           "control",
						OutputID:       "567b34857614d16292220beaca16ce34b939c75023a49cc43fa432fff51ca0dd",
						Position:       0,
						AssetID:        "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
						Amount:         41030000000,
						ControlProgram: "0014c3d320e1dc4fe787e9f13c1464e3ea5aae96a58f",
						Address:        "bm1qc0fjpcwuflnc06038s2xfcl2t2hfdfv07hgf77",
					},
					util.AnnotatedOutput{
						Type:           "control",
						OutputID:       "a8069d412e48c2b2994d2816758078cff46b215421706b4bad41f72a32928d92",
						Position:       1,
						AssetID:        "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
						Amount:         200000000,
						ControlProgram: "0014bb93cdb4eca74b068321eeb84ac5d33686281b65",
						Address:        "bm1qhwfumd8v5a9sdqepa6uy43wnx6rzsxm9cp6j43",
					},
				},
				Fee: 20000000,
			},
		},
	}

	for i, c := range cases {
		jsonTx := BytomDecodeRawTx(c.rawTransaction)
		if jsonTx == nil {
			t.Fatal(errors.New("error"))
		}

		gotTx := &util.Transaction{}
		if err := json.Unmarshal(jsonTx, gotTx); err != nil {
			t.Fatal(err)
		}

		if !testutil.DeepEqual(gotTx, c.wantTx) {
			t.Errorf("case #%d, annotated transaction got=%#v, want=%#v", i, gotTx, c.wantTx)
		}
	}
}
