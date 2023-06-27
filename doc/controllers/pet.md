# Pet

Everything about your Pets

Find out more: [http://swagger.io](http://swagger.io)

```go
petController := client.PetController
```

## Class Name

`PetController`

## Methods

* [Update Pet](../../doc/controllers/pet.md#update-pet)
* [Add Pet](../../doc/controllers/pet.md#add-pet)
* [Find Pets by Status](../../doc/controllers/pet.md#find-pets-by-status)
* [Find Pets by Tags](../../doc/controllers/pet.md#find-pets-by-tags)
* [Get Pet by Id](../../doc/controllers/pet.md#get-pet-by-id)
* [Update Pet With Form](../../doc/controllers/pet.md#update-pet-with-form)
* [Delete Pet](../../doc/controllers/pet.md#delete-pet)
* [Upload File](../../doc/controllers/pet.md#upload-file)


# Update Pet

Update an existing pet by Id

```go
UpdatePet(
    name string,
    photoUrls []string,
    id *int64,
    category *models.Category,
    tags *[]models.Tag,
    petStatus *models.PetStatusEnum) (
    https.ApiResponse[models.Pet],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `name` | `string` | Form, Required | - |
| `photoUrls` | `[]string` | Form, Required | - |
| `id` | `*int64` | Form, Optional | - |
| `category` | [`*models.Category`](../../doc/models/category.md) | Form, Optional | - |
| `tags` | [`*[]models.Tag`](../../doc/models/tag.md) | Form, Optional | - |
| `petStatus` | [`*models.PetStatusEnum`](../../doc/models/pet-status-enum.md) | Form, Optional | pet status in the store |

## Response Type

[`models.Pet`](../../doc/models/pet.md)

## Example Usage

```go
name := "doggie"

photoUrls := []string{"photoUrls5", "photoUrls6", "photoUrls7"}

id := int64(10)

category := models.Category{}
categoryId := int64(1)
category.Id = &categoryId

categoryName := "Dogs"
category.Name = &categoryName


tags0 := models.Tag{}
tags := []models.Tag{tags0}


apiResponse, err := petController.UpdatePet(name, photoUrls, &id, &category, &tags, nil)
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
| 404 | Pet not found | `ApiError` |
| 405 | Validation exception | `ApiError` |


# Add Pet

Add a new pet to the store

```go
AddPet(
    name string,
    photoUrls []string,
    id *int64,
    category *models.Category,
    tags *[]models.Tag,
    petStatus *models.PetStatusEnum) (
    https.ApiResponse[models.Pet],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `name` | `string` | Form, Required | - |
| `photoUrls` | `[]string` | Form, Required | - |
| `id` | `*int64` | Form, Optional | - |
| `category` | [`*models.Category`](../../doc/models/category.md) | Form, Optional | - |
| `tags` | [`*[]models.Tag`](../../doc/models/tag.md) | Form, Optional | - |
| `petStatus` | [`*models.PetStatusEnum`](../../doc/models/pet-status-enum.md) | Form, Optional | pet status in the store |

## Response Type

[`models.Pet`](../../doc/models/pet.md)

## Example Usage

```go
name := "doggie"

photoUrls := []string{"photoUrls5", "photoUrls6", "photoUrls7"}

id := int64(10)

category := models.Category{}
categoryId := int64(1)
category.Id = &categoryId

categoryName := "Dogs"
category.Name = &categoryName


tags0 := models.Tag{}
tags := []models.Tag{tags0}


apiResponse, err := petController.AddPet(name, photoUrls, &id, &category, &tags, nil)
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


# Find Pets by Status

Multiple status values can be provided with comma separated strings

```go
FindPetsByStatus(
    status *models.StatusEnum) (
    https.ApiResponse[[]models.Pet],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `status` | [`*models.StatusEnum`](../../doc/models/status-enum.md) | Query, Optional | Status values that need to be considered for filter<br>**Default**: `"available"` |

## Response Type

[`[]models.Pet`](../../doc/models/pet.md)

## Example Usage

```go
status := models.StatusEnum("available")

apiResponse, err := petController.FindPetsByStatus(&status)
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
| 400 | Invalid status value | `ApiError` |


# Find Pets by Tags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

```go
FindPetsByTags(
    tags *[]string) (
    https.ApiResponse[[]models.Pet],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `tags` | `*[]string` | Query, Optional | Tags to filter by |

## Response Type

[`[]models.Pet`](../../doc/models/pet.md)

## Example Usage

```go
apiResponse, err := petController.FindPetsByTags(nil)
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
| 400 | Invalid tag value | `ApiError` |


# Get Pet by Id

Returns a single pet

```go
GetPetById(
    petId int64) (
    https.ApiResponse[models.Pet],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `petId` | `int64` | Template, Required | ID of pet to return |

## Response Type

[`models.Pet`](../../doc/models/pet.md)

## Example Usage

```go
petId := int64(152)

apiResponse, err := petController.GetPetById(petId)
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
| 404 | Pet not found | `ApiError` |


# Update Pet With Form

Updates a pet in the store with form data

```go
UpdatePetWithForm(
    petId int64,
    name *string,
    status *string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `petId` | `int64` | Template, Required | ID of pet that needs to be updated |
| `name` | `*string` | Query, Optional | Name of pet that needs to be updated |
| `status` | `*string` | Query, Optional | Status of pet that needs to be updated |

## Response Type

``

## Example Usage

```go
petId := int64(152)



resp, err := petController.UpdatePetWithForm(petId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 405 | Invalid input | `ApiError` |


# Delete Pet

delete a pet

```go
DeletePet(
    petId int64,
    apiKey *string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `petId` | `int64` | Template, Required | Pet id to delete |
| `apiKey` | `*string` | Header, Optional | - |

## Response Type

``

## Example Usage

```go
petId := int64(152)


resp, err := petController.DeletePet(petId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Invalid pet value | `ApiError` |


# Upload File

uploads an image

```go
UploadFile(
    petId int64,
    additionalMetadata *string,
    body *https.FileWrapper) (
    https.ApiResponse[models.PetImage],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `petId` | `int64` | Template, Required | ID of pet to update |
| `additionalMetadata` | `*string` | Query, Optional | Additional Metadata |
| `body` | `*https.FileWrapper` | Form, Optional | - |

## Response Type

[`models.PetImage`](../../doc/models/pet-image.md)

## Example Usage

```go
petId := int64(152)




apiResponse, err := petController.UploadFile(petId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

