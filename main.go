package main

import (
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

func newHandler() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/health":
			ctx.Response.Header.Set("content-type", "application/json")
			fmt.Fprintln(ctx, `{"ok":true}`)
		case "/version":
			ctx.Response.Header.Set("Content-Type", "text/plain")
			ctx.SetStatusCode(fasthttp.StatusOK)
			ctx.WriteString("v1")
		default:
			ctx.SetStatusCode(fasthttp.StatusNotFound)
		}
	}
}

func main() {
	handler := newHandler()
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = ":8081"
	}
	fmt.Println("listening on", addr)
	if err := fasthttp.ListenAndServe(addr, handler); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
