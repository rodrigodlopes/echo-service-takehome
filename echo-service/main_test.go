package main

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerEchoesRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "/foo/bar?a=1", strings.NewReader("hello"))
	req.Header.Set("X-Test", "value")
	rec := httptest.NewRecorder()

	handler(rec, req)

	if rec.Code != 200 {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var got Echo
	if err := json.NewDecoder(rec.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if got.Path != "/foo/bar" {
		t.Errorf("expected path /foo/bar, got %q", got.Path)
	}
	if got.Body != "hello" {
		t.Errorf("expected body %q, got %q", "hello", got.Body)
	}
	if got.Params["a"][0] != "1" {
		t.Errorf("expected param a=1, got %v", got.Params["a"])
	}
	if got.Headers["X-Test"][0] != "value" {
		t.Errorf("expected header X-Test=value, got %v", got.Headers["X-Test"])
	}
}
