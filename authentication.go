package swaggerpetstoreopenapi30

import (
    "github.com/apimatic/go-core-runtime/https"
    "net/http"
)

func CustomHeaderAuthentication(config Configuration) https.Authenticator {
    return func(requiresAuth bool) https.HttpInterceptor {
        if !requiresAuth {
            return https.PassThroughInterceptor
        }
        return func(req *http.Request,
            next https.HttpCallExecutor,
        ) https.HttpContext {
            req.Header.Add("api_key", config.ApiKey())
            return next(req)
        }
    }
}
