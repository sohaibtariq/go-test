
# Getting Started with Swagger Petstore - OpenAPI 3.0

## Introduction

This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about
Swagger at [https://swagger.io](https://swagger.io). In the third iteration of the pet store, we've switched to the design first approach!
You can now help us improve the API whether it's by making changes to the definition itself or to the code.
That way, with time, we can improve the API in general, and expose some of the new features in OAS3.

_If you're looking for the Swagger 2.0/OAS 2.0 version of Petstore, then click [here](https://editor.swagger.io/?url=https://petstore.swagger.io/v2/swagger.yaml). Alternatively, you can load via the `Edit > Load Petstore OAS 2.0` menu option!_

Some useful links:

- [The Pet Store repository](https://github.com/swagger-api/swagger-petstore)
- [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)

Find out more about Swagger: [http://swagger.io](http://swagger.io)

### Requirements

The SDK requires **Go version 1.18 or above**.

## Building

### Install Dependencies

Resolve all the SDK dependencies, using the `go get` command.

## Installation

The following section explains how to use the swaggerpetstoreopenapi30 library in a new project.

### 1. Add SDK as a Dependency to the Application

- Add the following lines to your application's `go.mod` file:

```go
replace swaggerpetstoreopenapi30 => ".\\Swagger Petstore - OpenAPI 3.0" // local path to the SDK

require swaggerpetstoreopenapi30 v0.0.0
```

- Resolve the dependencies in the updated `go.mod` file, using the `go get` command.

## Test the SDK

`Go Testing` is used as the testing framework. To run the unit tests of the SDK, navigate to the `tests` directory of the SDK and run the following command in the terminal:

```bash
$ go test
```

## Initialize the API Client

**_Note:_** Documentation for the client can be found [here.](doc/client.md)

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `httpConfiguration` | `https.HttpConfiguration` | Configurable http client options. |
| `apiKey` | `string` |  |

The API client can be initialized as follows:

```go
config := swaggerpetstoreopenapi30.ConfigurationFactory(
    swaggerpetstoreopenapi30.WithApiKey("api_key"),
)
config.LoadFromEnvironment()

client := swaggerpetstoreopenapi30.NewClient(config)
```

## Authorization

This API uses `Custom Header Signature`.

## List of APIs

* [Pet](doc/controllers/pet.md)
* [Store](doc/controllers/store.md)
* [User](doc/controllers/user.md)

## Classes Documentation

* [HttpConfiguration](doc/http-configuration.md)
* [RetryConfiguration](doc/retry-configuration.md)

