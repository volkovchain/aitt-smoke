package main

import (
	"net"
	"strings"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestVersionRoute(t *testing.T) {
	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	srv := &fasthttp.Server{Handler: newHandler()}
	go func() {
		_ = srv.Serve(ln)
	}()

	client := &fasthttp.Client{
		Dial: func(addr string) (net.Conn, error) {
			return ln.Dial()
		},
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI("http://example.com/version")

	if err := client.DoTimeout(req, resp, 5*time.Second); err != nil {
		t.Fatalf("GET /version: %v", err)
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		t.Errorf("status = %d, want %d", resp.StatusCode(), fasthttp.StatusOK)
	}

	ct := string(resp.Header.Peek("Content-Type"))
	if !strings.Contains(ct, "text/plain") {
		t.Errorf("Content-Type = %q, want it to contain %q", ct, "text/plain")
	}

	body := string(resp.Body())
	if body != "v1" {
		t.Errorf("body = %q, want %q", body, "v1")
	}
}
