
# Customer

## Structure

`Customer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*int64` | Optional | - |
| `username` | `*string` | Optional | - |
| `address` | [`*[]models.Address`](../../doc/models/address.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": 100000,
  "username": "fehguy",
  "address": [
    {
      "street": "street4",
      "city": "city4",
      "state": "state0",
      "zip": "zip8"
    }
  ]
}
```

