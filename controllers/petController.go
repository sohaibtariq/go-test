package controllers

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/apiError"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "net/http"
    "swaggerpetstoreopenapi30/models"
)

type PetController struct {
    baseController
}

func NewPetController(baseController baseController) *PetController {
    petController := PetController{baseController: baseController}
    return &petController
}

// Update an existing pet by Id
func (p *PetController) UpdatePet(
    name string,
    photoUrls []string,
    id *int64,
    category *models.Category,
    tags *[]models.Tag,
    petStatus *models.PetStatusEnum) (
    https.ApiResponse[models.Pet],
    error) {
    req := p.prepareRequest("PUT", "/pet")
    req.Authenticate(true)
    req.Header("Content-Type", "application/x-www-form-urlencoded")
    req.FormParam("name", name)
    req.FormParam("photoUrls", photoUrls)
    if id != nil {
        req.FormParam("id", *id)
    }
    if category != nil {
        req.FormParam("category", *category)
    }
    if tags != nil {
        req.FormParam("tags", *tags)
    }
    if petStatus != nil {
        req.FormParam("petStatus", *petStatus)
    }
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[models.Pet]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[models.Pet]{ Response: resp}, err
    }
    
    var result models.Pet
    result, err = utilities.DecodeResults[models.Pet](decoder)
    if err != nil {
        return https.ApiResponse[models.Pet]{ Response: resp}, err
    }
    
    if resp.StatusCode == 400 {
        		err = apiError.NewApiError(400, "Invalid ID supplied")
    }
    if resp.StatusCode == 404 {
        		err = apiError.NewApiError(404, "Pet not found")
    }
    if resp.StatusCode == 405 {
        		err = apiError.NewApiError(405, "Validation exception")
    }
    return https.NewApiResponse(result, resp), err
}

// Add a new pet to the store
func (p *PetController) AddPet(
    name string,
    photoUrls []string,
    id *int64,
    category *models.Category,
    tags *[]models.Tag,
    petStatus *models.PetStatusEnum) (
    https.ApiResponse[models.Pet],
    error) {
    req := p.prepareRequest("POST", "/pet")
    req.Authenticate(true)
    req.Header("Content-Type", "application/x-www-form-urlencoded")
    req.FormParam("name", name)
    req.FormParam("photoUrls", photoUrls)
    if id != nil {
        req.FormParam("id", *id)
    }
    if category != nil {
        req.FormParam("category", *category)
    }
    if tags != nil {
        req.FormParam("tags", *tags)
    }
    if petStatus != nil {
        req.FormParam("petStatus", *petStatus)
    }
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[models.Pet]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[models.Pet]{ Response: resp}, err
    }
    
    var result models.Pet
    result, err = utilities.DecodeResults[models.Pet](decoder)
    if err != nil {
        return https.ApiResponse[models.Pet]{ Response: resp}, err
    }
    
    if resp.StatusCode == 405 {
        		err = apiError.NewApiError(405, "Invalid input")
    }
    return https.NewApiResponse(result, resp), err
}

// Multiple status values can be provided with comma separated strings
func (p *PetController) FindPetsByStatus(status *models.StatusEnum) (
    https.ApiResponse[[]models.Pet],
    error) {
    req := p.prepareRequest("GET", "/pet/findByStatus")
    req.Authenticate(true)
    if status != nil {
        req.QueryParam("status", *status)
    }
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[[]models.Pet]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[[]models.Pet]{ Response: resp}, err
    }
    
    var result []models.Pet
    result, err = utilities.DecodeResults[[]models.Pet](decoder)
    if err != nil {
        return https.ApiResponse[[]models.Pet]{ Response: resp}, err
    }
    
    if resp.StatusCode == 400 {
        		err = apiError.NewApiError(400, "Invalid status value")
    }
    return https.NewApiResponse(result, resp), err
}

// Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
func (p *PetController) FindPetsByTags(tags *[]string) (
    https.ApiResponse[[]models.Pet],
    error) {
    req := p.prepareRequest("GET", "/pet/findByTags")
    req.Authenticate(true)
    if tags != nil {
        req.QueryParam("tags", *tags)
    }
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[[]models.Pet]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[[]models.Pet]{ Response: resp}, err
    }
    
    var result []models.Pet
    result, err = utilities.DecodeResults[[]models.Pet](decoder)
    if err != nil {
        return https.ApiResponse[[]models.Pet]{ Response: resp}, err
    }
    
    if resp.StatusCode == 400 {
        		err = apiError.NewApiError(400, "Invalid tag value")
    }
    return https.NewApiResponse(result, resp), err
}

// Returns a single pet
func (p *PetController) GetPetById(petId int64) (
    https.ApiResponse[models.Pet],
    error) {
    req := p.prepareRequest("GET", fmt.Sprintf("/pet/%s", petId))
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[models.Pet]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[models.Pet]{ Response: resp}, err
    }
    
    var result models.Pet
    result, err = utilities.DecodeResults[models.Pet](decoder)
    if err != nil {
        return https.ApiResponse[models.Pet]{ Response: resp}, err
    }
    
    if resp.StatusCode == 400 {
        		err = apiError.NewApiError(400, "Invalid ID supplied")
    }
    if resp.StatusCode == 404 {
        		err = apiError.NewApiError(404, "Pet not found")
    }
    return https.NewApiResponse(result, resp), err
}

// Updates a pet in the store with form data
func (p *PetController) UpdatePetWithForm(
    petId int64,
    name *string,
    status *string) (
    *http.Response,
    error) {
    req := p.prepareRequest("POST", fmt.Sprintf("/pet/%s", petId))
    req.Authenticate(true)
    if name != nil {
        req.QueryParam("name", *name)
    }
    if status != nil {
        req.QueryParam("status", *status)
    }
    
    context, err := req.Call()
    err = validateResponse(*context.Response)
    if err != nil {
        return context.Response, err
    }
    if context.Response.StatusCode == 405 {
        		err = apiError.NewApiError(405, "Invalid input")
    }
    return context.Response, err
}

// delete a pet
func (p *PetController) DeletePet(
    petId int64,
    apiKey *string) (
    *http.Response,
    error) {
    req := p.prepareRequest("DELETE", fmt.Sprintf("/pet/%s", petId))
    req.Authenticate(true)
    if apiKey != nil {
        req.Header("api_key", *apiKey)
    }
    
    context, err := req.Call()
    err = validateResponse(*context.Response)
    if err != nil {
        return context.Response, err
    }
    if context.Response.StatusCode == 400 {
        		err = apiError.NewApiError(400, "Invalid pet value")
    }
    return context.Response, err
}

// uploads an image
func (p *PetController) UploadFile(
    petId int64,
    additionalMetadata *string,
    body *https.FileWrapper) (
    https.ApiResponse[models.PetImage],
    error) {
    req := p.prepareRequest("POST", fmt.Sprintf("/pet/%s/uploadImage", petId))
    req.Authenticate(true)
    req.Header("Content-Type", "application/octet-stream")
    if additionalMetadata != nil {
        req.QueryParam("additionalMetadata", *additionalMetadata)
    }
    formFields := []https.FormParam{}
    if body != nil {
        bodyParam := https.FormParam{Key: "body", Value: *body, Headers: http.Header{}}
        formFields = append(formFields, bodyParam)
    }
    req.FormData(formFields)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[models.PetImage]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[models.PetImage]{ Response: resp}, err
    }
    
    var result models.PetImage
    result, err = utilities.DecodeResults[models.PetImage](decoder)
    if err != nil {
        return https.ApiResponse[models.PetImage]{ Response: resp}, err
    }
    
    return https.NewApiResponse(result, resp), err
}
