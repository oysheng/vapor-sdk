# vapor-sdk

## `DecodeRawTransaction`

Decode a serialized transaction hex string into a JSON object describing the transaction.

### Parameters

`Object`:

- `String` - *raw_transaction*, hexstring of raw transaction.

### Returns

`Object`:

- `String` - *tx_id*, transaction ID.
- `Integer` - *version*, version of transaction.
- `String` - *size*, size of transaction.
- `String` - *time_range*, time range of transaction.
- `String` - *fee*, fee for sending transaction.
- `Array of Object` - *inputs*, object of inputs for the transaction.
  - `String` - *type*, the type of input action, available option include: 'veto', 'cross_chain_in', 'spend', 'issue', 'coinbase'.
  - `String` - *input_id*, hash of input action.
  - `String` - *asset*, asset id.
  - `Integer` - *amount*, amount of asset.
  - `String` - *script*, control program of account, it only exist when type is 'veto', 'spend', 'cross_chain_in'.
  - `String` - *address*, address of account, it only exist when type is 'veto', 'spend', 'cross_chain_in'.
  - `String` - *issuance_program*, issuance program, it only exist when type is 'issue'.
  - `String` - *asset_definition*, asset definition, it only exist when type is 'issue'.
  - `String` - *spent_output_id*, the front of outputID to be spent in this input, it only exist when type is 'veto', 'spend', 'cross_chain_in'.
  - `String` - *arbitrary*, arbitrary infomation can be set by miner, it only exist when type is 'coinbase'.
  - `Array of String` - *arguments*, witness arguments.
  - `String` - *vote*, vote xpub, it only exist when type is 'veto'.
  - `String` - *sign_data*, sign data, it only exist when type is 'veto', 'cross_chain_in', 'spend', 'issue'.
- `Array of Object` - *outputs*, object of outputs for the transaction.
  - `String` - *type*, the type of output action, available option include: 'control', 'cross_chain_out', 'vote', 'retire'.
  - `String` - *utxo_id*, outputid related to utxo.
  - `Integer` - *position*, position of outputs.
  - `String` - *asset*, asset id.
  - `Integer` - *amount*, amount of asset.
  - `String` - *script*, control program of account.
  - `String` - *address*, address of account.
  - `String` - *vote*, vote xpub, it only exist when type is 'vote'.

### Example
#### bytom transaction
```js
// Request
{
  "raw_transaction": "070100010161015fc8215913a270d3d953ef431626b19a89adf38e2486bb235da732f0afed515299ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8099c4d59901000116001456ac170c7965eeac1cc34928c9f464e3f88c17d8630240b1e99a3590d7db80126b273088937a87ba1e8d2f91021a2fd2c36579f7713926e8c7b46c047a43933b008ff16ecc2eb8ee888b4ca1fe3fdf082824e0b3899b02202fb851c6ed665fcd9ebc259da1461a1e284ac3b27f5e86c84164aa518648222602013effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80bbd0ec980101160014c3d320e1dc4fe787e9f13c1464e3ea5aae96a58f00013cffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8084af5f01160014bb93cdb4eca74b068321eeb84ac5d33686281b6500"
}

// Result
{
  "tx_id": "4c97d7412b04d49acc33762fc748cd0780d8b44086c229c1a6d0f2adfaaac2db",
  "version": 1,
  "size": 332,
  "time_range": 0,
  "fee": 20000000,
  "inputs": [
    {
      "type": "spend",
      "input_id": "9963265eb601df48501cc240e1480780e9ed6e0c8f18fd7dd57954068c5dfd02",
      "asset": "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
      "amount": 41250000000,
      "script": "001456ac170c7965eeac1cc34928c9f464e3f88c17d8",
      "address": "bm1q26kpwrrevhh2c8xrfy5vnaryu0ugc97c3j896t",
      "spent_output_id": "01bb3309666618a1507cb5be845b17dee5eb8028ee7e71b17d74b4dc97085bc8",
      "witness_arguments": [
        "b1e99a3590d7db80126b273088937a87ba1e8d2f91021a2fd2c36579f7713926e8c7b46c047a43933b008ff16ecc2eb8ee888b4ca1fe3fdf082824e0b3899b02",
        "2fb851c6ed665fcd9ebc259da1461a1e284ac3b27f5e86c84164aa5186482226"
      ],
      "sign_data": "8d2bb534c819464472a94b41cea788e97a2c9dae09a6cb3b7024a44ce5a27835"
    }
  ],
  "outputs": [
    {
      "type": "control",
      "utxo_id": "567b34857614d16292220beaca16ce34b939c75023a49cc43fa432fff51ca0dd",
      "position": 0,
      "asset": "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
      "amount": 41030000000,
      "script": "0014c3d320e1dc4fe787e9f13c1464e3ea5aae96a58f",
      "address": "bm1qc0fjpcwuflnc06038s2xfcl2t2hfdfv07hgf77"
    },
    {
      "type": "control",
      "id": "a8069d412e48c2b2994d2816758078cff46b215421706b4bad41f72a32928d92",
      "position": 1,
      "asset": "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
      "amount": 200000000,
      "script": "0014bb93cdb4eca74b068321eeb84ac5d33686281b65",
      "address": "bm1qhwfumd8v5a9sdqepa6uy43wnx6rzsxm9cp6j43"
    }
  ]
}
```

#### vapor transaction
```js
// Request
{
  "raw_transaction": "07010001015f015d13c41cc617304ba0866fa59f07d7bb2bcab60c43e5cc79bb75a4dd97471cdcbaffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80ade20400011600144b6995dc11354d44c6e382c19d6b92bdbbd3aea1010002013e003cffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc096b102011600149682e64b2114f7c2581ab1ba0c67315d06aaea8200013e003cffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc096b10201160014da26416fa79947ec6a569e0493dbffec1a3f223400"
}

// Result
{
  "tx_id": "ab3130d01b1d41c4d772f258fc5d2b38660d5d41e44d107fe935eb2b85015990",
  "version": 1,
  "size": 234,
  "time_range": 0,
  "fee": 0,
  "inputs": [
    {
      "type": "spend",
      "input_id": "84d1a1ebf56015eec08b4f21556ad5ee5faed86e0f3b6b0d2366f8ea4280959d",
      "asset": "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
      "amount": 10000000,
      "script": "00144b6995dc11354d44c6e382c19d6b92bdbbd3aea1",
      "address": "vp1qfd5ethq3x4x5f3hrstqe66ujhkaa8t4p8vud4p",
      "spent_output_id": "bf099ee53337ce14411e31d392fcf65f25ab7bc993f10ca06fbaf7d74a6d9c5d",
      "sign_data": "8d30d28aeed8c8b69a941ecc1e0da8dafd8b022b48c82c0d5eb8283870569aa8",
      "arguments": null
    }
  ],
  "outputs": [
    {
      "type": "control",
      "utxo_id": "ff91113d21213c9a2338b512257345be91f85e403d8df51f57e06673bd231c49",
      "position": 0,
      "asset": "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
      "amount": 5000000,
      "script": "00149682e64b2114f7c2581ab1ba0c67315d06aaea82",
      "address": "vp1qj6pwvjepznmuykq6kxaqcee3t5r2465z0hmr70"
    },
    {
      "type": "control",
      "utxo_id": "7a16706c5deabee77c3626eaf7e71e1ccedda771e2f2aedcb009ebb044c367d1",
      "position": 1,
      "asset": "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
      "amount": 5000000,
      "script": "0014da26416fa79947ec6a569e0493dbffec1a3f2234",
      "address": "vp1qmgnyzma8n9r7c6jknczf8kllasdr7g35whjwpg"
    }
  ]
}
```

