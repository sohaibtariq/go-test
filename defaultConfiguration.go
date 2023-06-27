package swaggerpetstoreopenapi30

import (
    "github.com/apimatic/go-core-runtime/https"
    "net/http"
)

func DefaultRetryConfiguration() https.RetryConfiguration {
    return https.NewRetryConfiguration(
        https.WithMaxRetryAttempts(0),
        https.WithRetryOnTimeout(true),
        https.WithRetryInterval(1),
        https.WithMaximumRetryWaitTime(0),
        https.WithBackoffFactor(2),
        https.WithHttpStatusCodesToRetry([]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524}),
        https.WithHttpMethodsToRetry([]string{"GET", "PUT"}),
    )
}

func DefaultHttpConfiguration() https.HttpConfiguration {
    return https.NewHttpConfiguration(
        https.WithTimeout(0),
        https.WithTransport(http.DefaultTransport),
        https.WithRetryConfiguration(DefaultRetryConfiguration()),
    )
}

func DefaultConfiguration() Configuration {
    return newConfiguration(
        WithEnvironment(PRODUCTION),
        WithApiKey(""),
        WithHttpConfiguration(DefaultHttpConfiguration()),
    )
}
