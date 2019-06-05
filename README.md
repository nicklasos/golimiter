# Go Rate Limiter
Thanks to Alex Edwards for great [article](https://www.alexedwards.net/blog/how-to-rate-limit-http-requests) about rate limits on Go!

### Usage
More about rate limit in [golang.org/x/time/rate](https://godoc.org/golang.org/x/time/rate)
```go

limit := golimiter.NewLimiter(5, 2) // rate and bucket

if limit.IsBanned("user string identified") {
    w.WriteHeader(http.StatusTooManyRequests)
    w.Write([]byte("429 - Too many requests"))
    retur
}

if limit.Allow("user string identifier") == false {
    // User has reached its quota

    limit.Ban("user string identifier", time.Second*60) // You can also ban user
}
```