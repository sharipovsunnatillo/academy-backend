package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

func Ensure(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, middleware := range middlewares {
			next = middleware(next)
		}
		return next
	}
}
