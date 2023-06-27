package tests

import (
    "github.com/apimatic/go-core-runtime/testHelper"
    "swaggerpetstoreopenapi30/models"
    "testing"
    "time"
)

func TestStoreControllerTestGetInventory(t *testing.T) {
    apiResponse, err := StoreController.GetInventory()
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}

func TestStoreControllerTestPlaceOrder(t *testing.T) {
    id := int64(10)
    petId := int64(198772)
    quantity := 7
    var shipDate *time.Time = ""
    var orderStatus *models.OrderStatusEnum = ""
    var complete *bool = false
    apiResponse, err := StoreController.PlaceOrder(&id, &petId, &quantity, &shipDate, &orderStatus, &complete)
    if err != nil {
        t.Errorf("Endpoint call failed: %v", err)
    }
    testHelper.CheckResponseStatusCode(t, apiResponse.Response.StatusCode, 200)
    expectedHeaders:= []testHelper.TestHeader{
        testHelper.NewTestHeader(true,"Content-Type","application/json"),
    }
    testHelper.CheckResponseHeaders(t, apiResponse.Response.Header, expectedHeaders, true)
}
