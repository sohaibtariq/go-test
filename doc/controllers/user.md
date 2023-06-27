# User

Operations about user

```go
userController := client.UserController
```

## Class Name

`UserController`

## Methods

* [Create User](../../doc/controllers/user.md#create-user)
* [Create Users With List Input](../../doc/controllers/user.md#create-users-with-list-input)
* [Login User](../../doc/controllers/user.md#login-user)
* [Logout User](../../doc/controllers/user.md#logout-user)
* [Get User by Name](../../doc/controllers/user.md#get-user-by-name)
* [Update User](../../doc/controllers/user.md#update-user)
* [Delete User](../../doc/controllers/user.md#delete-user)


# Create User

This can only be done by the logged in user.

```go
CreateUser(
    id *int64,
    username *string,
    firstName *string,
    lastName *string,
    email *string,
    password *string,
    phone *string,
    userStatus *int) (
    https.ApiResponse[models.User],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `*int64` | Form, Optional | - |
| `username` | `*string` | Form, Optional | - |
| `firstName` | `*string` | Form, Optional | - |
| `lastName` | `*string` | Form, Optional | - |
| `email` | `*string` | Form, Optional | - |
| `password` | `*string` | Form, Optional | - |
| `phone` | `*string` | Form, Optional | - |
| `userStatus` | `*int` | Form, Optional | User Status |

## Response Type

[`models.User`](../../doc/models/user.md)

## Example Usage

```go
id := int64(10)

username := "theUser"

firstName := "John"

lastName := "James"

email := "john@email.com"

password := "12345"

phone := "12345"

userStatus := 1

apiResponse, err := userController.CreateUser(&id, &username, &firstName, &lastName, &email, &password, &phone, &userStatus)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Users With List Input

Creates list of users with given input array

```go
CreateUsersWithListInput(
    body *[]models.User) (
    https.ApiResponse[models.User],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*[]models.User`](../../doc/models/user.md) | Body, Optional | - |

## Response Type

[`models.User`](../../doc/models/user.md)

## Example Usage

```go
body0 := models.User{}
body0Id := int64(10)
body0.Id = &body0Id

body0Username := "theUser"
body0.Username = &body0Username

body0FirstName := "John"
body0.FirstName = &body0FirstName

body0LastName := "James"
body0.LastName = &body0LastName

body0Email := "john@email.com"
body0.Email = &body0Email

body0Password := "12345"
body0.Password = &body0Password

body0Phone := "12345"
body0.Phone = &body0Phone

body0UserStatus := 1
body0.UserStatus = &body0UserStatus

body := []models.User{body0}

apiResponse, err := userController.CreateUsersWithListInput(&body)
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
| Default | successful operation | `ApiError` |


# Login User

Logs user into the system

```go
LoginUser(
    username *string,
    password *string) (
    https.ApiResponse[string],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `username` | `*string` | Query, Optional | The user name for login |
| `password` | `*string` | Query, Optional | The password for login in clear text |

## Response Type

`string`

## Example Usage

```go
apiResponse, err := userController.LoginUser(nil, nil)
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
| 400 | Invalid username/password supplied | `ApiError` |


# Logout User

Logs out current logged in user session

```go
LogoutUser() (
    http.Response,
    error)
```

## Response Type

``

## Example Usage

```go
resp, err := userController.LogoutUser()
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Get User by Name

Get user by user name

```go
GetUserByName(
    name string) (
    https.ApiResponse[models.User],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `name` | `string` | Template, Required | The name that needs to be fetched. Use user1 for testing. |

## Response Type

[`models.User`](../../doc/models/user.md)

## Example Usage

```go
name := "name0"

apiResponse, err := userController.GetUserByName(name)
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
| 400 | Invalid username supplied | `ApiError` |
| 404 | User not found | `ApiError` |


# Update User

This can only be done by the logged in user.

```go
UpdateUser(
    name string,
    id *int64,
    username *string,
    firstName *string,
    lastName *string,
    email *string,
    password *string,
    phone *string,
    userStatus *int) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `name` | `string` | Template, Required | name that need to be deleted |
| `id` | `*int64` | Form, Optional | - |
| `username` | `*string` | Form, Optional | - |
| `firstName` | `*string` | Form, Optional | - |
| `lastName` | `*string` | Form, Optional | - |
| `email` | `*string` | Form, Optional | - |
| `password` | `*string` | Form, Optional | - |
| `phone` | `*string` | Form, Optional | - |
| `userStatus` | `*int` | Form, Optional | User Status |

## Response Type

``

## Example Usage

```go
name := "name0"

id := int64(10)

username := "theUser"

firstName := "John"

lastName := "James"

email := "john@email.com"

password := "12345"

phone := "12345"

userStatus := 1

resp, err := userController.UpdateUser(name, &id, &username, &firstName, &lastName, &email, &password, &phone, &userStatus)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Delete User

This can only be done by the logged in user.

```go
DeleteUser(
    name string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `name` | `string` | Template, Required | The name that needs to be deleted |

## Response Type

``

## Example Usage

```go
name := "name0"

resp, err := userController.DeleteUser(name)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Invalid username supplied | `ApiError` |
| 404 | User not found | `ApiError` |

