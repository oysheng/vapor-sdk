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
  - `String` - *type*, the type of input action, available option include: 'veto', 'cross_chain_in', 'spend', 'coinbase'.
  - `String` - *input_id*, hash of input action.
  - `String` - *asset*, asset id.
  - `Integer` - *amount*, amount of asset.
  - `String` - *script*, control program of account, it only exist when type is 'veto', 'spend', 'cross_chain_in'.
  - `String` - *address*, address of account, it only exist when type is 'veto', 'spend', 'cross_chain_in'.
  - `String` - *spent_output_id*, the front of outputID to be spent in this input, it only exist when type is 'veto', 'spend', 'cross_chain_in'.
  - `String` - *arbitrary*, arbitrary infomation can be set by miner, it only exist when type is 'coinbase'.
  - `Array of String` - *arguments*, witness arguments.
  - `String` - *vote*, vote xpub, it only exist when type is 'veto'.
  - `String` - *sign_data*, sign data, it only exist when type is 'veto', 'cross_chain_in', 'spend'.
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

