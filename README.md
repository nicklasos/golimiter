# Go Rate Limiter
Thanks to Alex Edwards for great [article](https://www.alexedwards.net/blog/how-to-rate-limit-http-requests) about rate limits on Go!

### Usage
More about rate limit in [golang.org/x/time/rate](https://godoc.org/golang.org/x/time/rate)
```go

limit := golimiter.NewLimiter(5, 2) // rate and bucket

if limit.Allow("user string idendifier") == false {
    // You have been banned!
}
```