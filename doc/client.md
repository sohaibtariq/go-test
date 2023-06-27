
# Client Class Documentation

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

## Swagger Petstore - OpenAPI 3.0 Client

The gateway for the SDK. This class acts as a factory for the Controllers and also holds the configuration of the SDK.

## Controllers

| Name | Description |
|  --- | --- |
| pet | Gets PetController |
| store | Gets StoreController |
| user | Gets UserController |

