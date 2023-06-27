package controllers

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/apiError"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "net/http"
    "swaggerpetstoreopenapi30/models"
    "time"
)

type StoreController struct {
    baseController
}

func NewStoreController(baseController baseController) *StoreController {
    storeController := StoreController{baseController: baseController}
    return &storeController
}

// Returns a map of status codes to quantities
func (s *StoreController) GetInventory() (
    https.ApiResponse[map[string]int],
    error) {
    req := s.prepareRequest("GET", "/store/inventory")
    req.Authenticate(true)
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[map[string]int]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[map[string]int]{ Response: resp}, err
    }
    
    var result map[string]int
    result, err = utilities.DecodeResults[map[string]int](decoder)
    if err != nil {
        return https.ApiResponse[map[string]int]{ Response: resp}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Place a new order in the store
func (s *StoreController) PlaceOrder(
    id *int64,
    petId *int64,
    quantity *int,
    shipDate *time.Time,
    orderStatus *models.OrderStatusEnum,
    complete *bool) (
    https.ApiResponse[models.Order],
    error) {
    req := s.prepareRequest("POST", "/store/order")
    req.Authenticate(true)
    req.Header("Content-Type", "application/x-www-form-urlencoded")
    if id != nil {
        req.FormParam("id", *id)
    }
    if petId != nil {
        req.FormParam("petId", *petId)
    }
    if quantity != nil {
        req.FormParam("quantity", *quantity)
    }
    if shipDate != nil {
        req.FormParam("shipDate", shipDate.Format(time.RFC3339))
    }
    if orderStatus != nil {
        req.FormParam("orderStatus", *orderStatus)
    }
    if complete != nil {
        req.FormParam("complete", *complete)
    }
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[models.Order]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[models.Order]{ Response: resp}, err
    }
    
    var result models.Order
    result, err = utilities.DecodeResults[models.Order](decoder)
    if err != nil {
        return https.ApiResponse[models.Order]{ Response: resp}, err
    }
    
    if resp.StatusCode == 405 {
        		err = apiError.NewApiError(405, "Invalid input")
    }
    return https.NewApiResponse(result, resp), err
}

// For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.
func (s *StoreController) GetOrderById(orderId int64) (
    https.ApiResponse[models.Order],
    error) {
    req := s.prepareRequest("GET", fmt.Sprintf("/store/order/%s", orderId))
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[models.Order]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[models.Order]{ Response: resp}, err
    }
    
    var result models.Order
    result, err = utilities.DecodeResults[models.Order](decoder)
    if err != nil {
        return https.ApiResponse[models.Order]{ Response: resp}, err
    }
    
    if resp.StatusCode == 400 {
        		err = apiError.NewApiError(400, "Invalid ID supplied")
    }
    if resp.StatusCode == 404 {
        		err = apiError.NewApiError(404, "Order not found")
    }
    return https.NewApiResponse(result, resp), err
}

// For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors
func (s *StoreController) DeleteOrder(orderId int64) (
    *http.Response,
    error) {
    req := s.prepareRequest("DELETE", fmt.Sprintf("/store/order/%s", orderId))
    req.Authenticate(true)
    
    context, err := req.Call()
    err = validateResponse(*context.Response)
    if err != nil {
        return context.Response, err
    }
    if context.Response.StatusCode == 400 {
        		err = apiError.NewApiError(400, "Invalid ID supplied")
    }
    if context.Response.StatusCode == 404 {
        		err = apiError.NewApiError(404, "Order not found")
    }
    return context.Response, err
}
