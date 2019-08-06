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
  - `String` - *asset_id*, asset id.
  - `Integer` - *amount*, amount of asset.
  - `String` - *control_program*, control program of account, it only exist when type is 'veto', 'spend', 'cross_chain_in'.
  - `String` - *address*, address of account, it only exist when type is 'veto', 'spend', 'cross_chain_in'.
  - `String` - *spent_output_id*, the front of outputID to be spent in this input, it only exist when type is 'veto', 'spend', 'cross_chain_in'.
  - `String` - *arbitrary*, arbitrary infomation can be set by miner, it only exist when type is 'coinbase'.
  - `Array of String` - *witness_arguments*, witness arguments.
- `Array of Object` - *outputs*, object of outputs for the transaction.
  - `String` - *type*, the type of output action, available option include: 'control', 'cross_chain_out', 'vote', 'retire'.
  - `String` - *output_id*, outputid related to utxo.
  - `Integer` - *position*, position of outputs.
  - `String` - *asset_id*, asset id.
  - `Integer` - *amount*, amount of asset.
  - `String` - *control_program*, control program of account.
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
      "input_id": "7d9c3a6481fd249c78f6037d3c999a3fe753882bd13e072cecc8ce92fbbbb41b",
      "asset_id": "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
      "amount": 10000000,
      "control_program": "00144b6995dc11354d44c6e382c19d6b92bdbbd3aea1",
      "address": "sp1qfd5ethq3x4x5f3hrstqe66ujhkaa8t4p0f7942",
      "spent_output_id": "873cd20c2cd260e1d2902f173bbc32490a9aa184b8e47aaedf3f37d7bf5225dd",
      "witness_arguments": null
    }
  ],
  "outputs": [
    {
      "type": "control",
      "output_id": "1df78ee679f30bb4597e1c3e459a0cd0429e69de875dd85a57fa34f94a59aba4",
      "position": 0,
      "asset_id": "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
      "amount": 5000000,
      "control_program": "00149682e64b2114f7c2581ab1ba0c67315d06aaea82",
      "address": "sp1qj6pwvjepznmuykq6kxaqcee3t5r2465z8jet7y"
    },
    {
      "type": "control",
      "output_id": "6d1116c361d8001e5f2d491c796eb36b5cb53dd630c1310a9be13742fd6e9cbc",
      "position": 1,
      "asset_id": "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
      "amount": 5000000,
      "control_program": "0014da26416fa79947ec6a569e0493dbffec1a3f2234",
      "address": "sp1qmgnyzma8n9r7c6jknczf8kllasdr7g35xjsxpr"
    }
  ]
}
```

