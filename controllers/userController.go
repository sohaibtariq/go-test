package controllers

import (
    "fmt"
    "github.com/apimatic/go-core-runtime/apiError"
    "github.com/apimatic/go-core-runtime/https"
    "github.com/apimatic/go-core-runtime/utilities"
    "net/http"
    "swaggerpetstoreopenapi30/models"
)

type UserController struct {
    baseController
}

func NewUserController(baseController baseController) *UserController {
    userController := UserController{baseController: baseController}
    return &userController
}

// This can only be done by the logged in user.
func (u *UserController) CreateUser(
    id *int64,
    username *string,
    firstName *string,
    lastName *string,
    email *string,
    password *string,
    phone *string,
    userStatus *int) (
    https.ApiResponse[models.User],
    error) {
    req := u.prepareRequest("POST", "/user")
    req.Authenticate(true)
    req.Header("Content-Type", "application/x-www-form-urlencoded")
    if id != nil {
        req.FormParam("id", *id)
    }
    if username != nil {
        req.FormParam("username", *username)
    }
    if firstName != nil {
        req.FormParam("firstName", *firstName)
    }
    if lastName != nil {
        req.FormParam("lastName", *lastName)
    }
    if email != nil {
        req.FormParam("email", *email)
    }
    if password != nil {
        req.FormParam("password", *password)
    }
    if phone != nil {
        req.FormParam("phone", *phone)
    }
    if userStatus != nil {
        req.FormParam("userStatus", *userStatus)
    }
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[models.User]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[models.User]{ Response: resp}, err
    }
    
    var result models.User
    result, err = utilities.DecodeResults[models.User](decoder)
    if err != nil {
        return https.ApiResponse[models.User]{ Response: resp}, err
    }
    
    return https.NewApiResponse(result, resp), err
}

// Creates list of users with given input array
func (u *UserController) CreateUsersWithListInput(body *[]models.User) (
    https.ApiResponse[models.User],
    error) {
    req := u.prepareRequest("POST", "/user/createWithList")
    req.Authenticate(true)
    req.Header("Content-Type", "application/json")
    if body != nil {
        req.Json(*body)
    }
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[models.User]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[models.User]{ Response: resp}, err
    }
    
    var result models.User
    result, err = utilities.DecodeResults[models.User](decoder)
    if err != nil {
        return https.ApiResponse[models.User]{ Response: resp}, err
    }
    
    if resp.StatusCode == 0 {
        		err = apiError.NewApiError(0, "successful operation")
    }
    return https.NewApiResponse(result, resp), err
}

// Logs user into the system
func (u *UserController) LoginUser(
    username *string,
    password *string) (
    https.ApiResponse[string],
    error) {
    req := u.prepareRequest("GET", "/user/login")
    req.Authenticate(true)
    if username != nil {
        req.QueryParam("username", *username)
    }
    if password != nil {
        req.QueryParam("password", *password)
    }
    str, resp, err := req.CallAsText()
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[string]{ Response: resp}, err
    }
    var result string = str

    if resp.StatusCode == 400 {
        		err = apiError.NewApiError(400, "Invalid username/password supplied")
    }
    return https.NewApiResponse(result, resp), err
}

// Logs out current logged in user session
func (u *UserController) LogoutUser() (
    *http.Response,
    error) {
    req := u.prepareRequest("GET", "/user/logout")
    req.Authenticate(true)
    context, err := req.Call()
    err = validateResponse(*context.Response)
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// Get user by user name
func (u *UserController) GetUserByName(name string) (
    https.ApiResponse[models.User],
    error) {
    req := u.prepareRequest("GET", fmt.Sprintf("/user/%s", name))
    req.Authenticate(true)
    
    decoder, resp, err := req.CallAsJson()
    if err != nil {
        return https.ApiResponse[models.User]{ Response: resp}, err
    }
    err = validateResponse(*resp)
    if err != nil {
        return https.ApiResponse[models.User]{ Response: resp}, err
    }
    
    var result models.User
    result, err = utilities.DecodeResults[models.User](decoder)
    if err != nil {
        return https.ApiResponse[models.User]{ Response: resp}, err
    }
    
    if resp.StatusCode == 400 {
        		err = apiError.NewApiError(400, "Invalid username supplied")
    }
    if resp.StatusCode == 404 {
        		err = apiError.NewApiError(404, "User not found")
    }
    return https.NewApiResponse(result, resp), err
}

// This can only be done by the logged in user.
func (u *UserController) UpdateUser(
    name string,
    id *int64,
    username *string,
    firstName *string,
    lastName *string,
    email *string,
    password *string,
    phone *string,
    userStatus *int) (
    *http.Response,
    error) {
    req := u.prepareRequest("PUT", fmt.Sprintf("/user/%s", name))
    req.Authenticate(true)
    req.Header("Content-Type", "application/x-www-form-urlencoded")
    if id != nil {
        req.FormParam("id", *id)
    }
    if username != nil {
        req.FormParam("username", *username)
    }
    if firstName != nil {
        req.FormParam("firstName", *firstName)
    }
    if lastName != nil {
        req.FormParam("lastName", *lastName)
    }
    if email != nil {
        req.FormParam("email", *email)
    }
    if password != nil {
        req.FormParam("password", *password)
    }
    if phone != nil {
        req.FormParam("phone", *phone)
    }
    if userStatus != nil {
        req.FormParam("userStatus", *userStatus)
    }
    
    context, err := req.Call()
    err = validateResponse(*context.Response)
    if err != nil {
        return context.Response, err
    }
    return context.Response, err
}

// This can only be done by the logged in user.
func (u *UserController) DeleteUser(name string) (
    *http.Response,
    error) {
    req := u.prepareRequest("DELETE", fmt.Sprintf("/user/%s", name))
    req.Authenticate(true)
    
    context, err := req.Call()
    err = validateResponse(*context.Response)
    if err != nil {
        return context.Response, err
    }
    if context.Response.StatusCode == 400 {
        		err = apiError.NewApiError(400, "Invalid username supplied")
    }
    if context.Response.StatusCode == 404 {
        		err = apiError.NewApiError(404, "User not found")
    }
    return context.Response, err
}
