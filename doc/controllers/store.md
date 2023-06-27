# Store

Access to Petstore orders

Find out more about our store: [http://swagger.io](http://swagger.io)

```go
storeController := client.StoreController
```

## Class Name

`StoreController`

## Methods

* [Get Inventory](../../doc/controllers/store.md#get-inventory)
* [Place Order](../../doc/controllers/store.md#place-order)
* [Get Order by Id](../../doc/controllers/store.md#get-order-by-id)
* [Delete Order](../../doc/controllers/store.md#delete-order)


# Get Inventory

Returns a map of status codes to quantities

```go
GetInventory() (
    https.ApiResponse[map[string]int],
    error)
```

## Response Type

`map[string]int`

## Example Usage

```go
apiResponse, err := storeController.GetInventory()
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Place Order

Place a new order in the store

```go
PlaceOrder(
    id *int64,
    petId *int64,
    quantity *int,
    shipDate *time.Time,
    orderStatus *models.OrderStatusEnum,
    complete *bool) (
    https.ApiResponse[models.Order],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*int64` | Form, Optional | - |
| `petId` | `*int64` | Form, Optional | - |
| `quantity` | `*int` | Form, Optional | - |
| `shipDate` | `*time.Time` | Form, Optional | - |
| `orderStatus` | [`*models.OrderStatusEnum`](../../doc/models/order-status-enum.md) | Form, Optional | Order Status |
| `complete` | `*bool` | Form, Optional | - |

## Response Type

[`models.Order`](../../doc/models/order.md)

## Example Usage

```go
id := int64(10)

petId := int64(198772)

quantity := 7





apiResponse, err := storeController.PlaceOrder(&id, &petId, &quantity, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 405 | Invalid input | `ApiError` |


# Get Order by Id

For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.

```go
GetOrderById(
    orderId int64) (
    https.ApiResponse[models.Order],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `int64` | Template, Required | ID of order that needs to be fetched |

## Response Type

[`models.Order`](../../doc/models/order.md)

## Example Usage

```go
orderId := int64(62)

apiResponse, err := storeController.GetOrderById(orderId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Invalid ID supplied | `ApiError` |
| 404 | Order not found | `ApiError` |


# Delete Order

For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors

```go
DeleteOrder(
    orderId int64) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `int64` | Template, Required | ID of the order that needs to be deleted |

## Response Type

``

## Example Usage

```go
orderId := int64(62)

resp, err := storeController.DeleteOrder(orderId)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Invalid ID supplied | `ApiError` |
| 404 | Order not found | `ApiError` |

