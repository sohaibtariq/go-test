
# Order

## Structure

`Order`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*int64` | Optional | - |
| `petId` | `*int64` | Optional | - |
| `quantity` | `*int` | Optional | - |
| `shipDate` | `*time.Time` | Optional | - |
| `orderStatus` | [`*models.OrderStatusEnum`](../../doc/models/order-status-enum.md) | Optional | Order Status |
| `complete` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "id": 10,
  "petId": 198772,
  "quantity": 7,
  "orderStatus": "approved",
  "shipDate": "2016-03-13T12:52:32.123Z"
}
```

