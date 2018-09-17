package main

import (
    "net/http"

    "golang.org/x/time/rate"
)


func limit(conf *config, next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    limiter := rate.NewLimiter(rate.Limit(conf.Max_rate), conf.Max_rate * 3)
        if limiter.Allow() == false {
            http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
            return
        }

        next.ServeHTTP(w, r)
    })
}

