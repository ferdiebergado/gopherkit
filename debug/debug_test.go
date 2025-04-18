package debug_test

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/ferdiebergado/gopherkit/debug"
	ghttp "github.com/ferdiebergado/gopherkit/http"
)

type contextKey string

const testUrl = "https://example.com"

const userID contextKey = "userID"

// Helper function to create a test request
func newTestRequest(method, urlStr string, body string) *http.Request {
	url, _ := url.Parse(urlStr)
	req := &http.Request{
		Method:     method,
		URL:        url,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       url.Host,
		Body:       nil,
	}
	if body != "" {
		req.Body = http.NoBody // Simulate an empty body without reading it
		req.ContentLength = int64(len(body))
	}
	return req
}

// Test basic GET request
func TestDumpRequestGetRequest(t *testing.T) {
	method := http.MethodGet
	req := newTestRequest(method, testUrl+"/path?query=value#fragment", "")

	req.Header.Set(ghttp.HeaderContentType, ghttp.MimeJSON)
	req.Header.Set("X-Custom-Header", "some value")
	req.RemoteAddr = "192.0.2.1:1234"

	result := debug.DumpRequest(req)

	// Validate key request fields
	if result["Method"] != method {
		t.Errorf("Expected Method=GET, got %v", result["Method"])
	}

	if result["Host"] != "example.com" {
		t.Errorf("Expected Host=example.com, got %v", result["Host"])
	}

	if result["RemoteAddr"] != "192.0.2.1:1234" {
		t.Errorf("Expected RemoteAddr=192.0.2.1:1234, got %v", result["RemoteAddr"])
	}

	headers, ok := result["Header"].(http.Header)
	if !ok || headers.Get(ghttp.HeaderContentType) != ghttp.MimeJSON {
		t.Errorf("Expected Content-Type=application/json; charset=utf-8, got %v", headers)
	}

	if req.URL.Path != "/path" {
		t.Errorf("Expected Path=/path, got %v", req.URL.Path)
	}
}

// Test POST request with form data
func TestDumpRequestPostRequest(t *testing.T) {
	body := "name=John&age=30"
	method := http.MethodPost
	req := newTestRequest(method, testUrl+"/submit", body)
	req.Header.Set(ghttp.HeaderContentType, ghttp.MimeFormUrlEncoded)

	// Attach body and parse form
	req.Body = io.NopCloser(strings.NewReader(body))
	req.ContentLength = int64(len(body))
	if err := req.ParseForm(); err != nil {
		t.Fatal("failed to parse form", err)
	} // REQUIRED to populate req.Form

	result := debug.DumpRequest(req)

	// Check method
	if result["Method"] != method {
		t.Errorf("Expected Method=POST, got %v", result["Method"])
	}

	// Check content length
	if result["ContentLength"] != int64(len(body)) {
		t.Errorf("Expected ContentLength=%d, got %v", len(body), result["ContentLength"])
	}

	// Assert form values properly
	form, ok := result["Form"].(map[string][]string) // Correct type assertion
	if !ok {
		t.Fatalf("Expected Form to be map[string][]string, got %T", result["Form"])
	}

	if form["name"][0] != "John" || form["age"][0] != "30" {
		t.Errorf("Expected Form[name]=John and Form[age]=30, got %v", form)
	}
}

// Test request with empty values
func TestDumpRequestEmptyRequest(t *testing.T) {
	method := http.MethodGet
	req := newTestRequest(method, testUrl, "")

	result := debug.DumpRequest(req)

	if result["Method"] != method {
		t.Errorf("Expected Method=GET, got %v", result["Method"])
	}

	if result["Host"] != "example.com" {
		t.Errorf("Expected Host=example.com, got %v", result["Host"])
	}

	if result["RequestURI"] != "" {
		t.Errorf("Expected RequestURI='', got %v", result["RequestURI"])
	}

	if len(result["Header"].(http.Header)) != 0 {
		t.Errorf("Expected empty headers, got %v", result["Header"])
	}
}

// Test request with cookies
func TestDumpRequestCookies(t *testing.T) {
	req := newTestRequest(http.MethodGet, testUrl, "")
	req.AddCookie(&http.Cookie{Name: "session", Value: "abc123"})

	result := debug.DumpRequest(req)

	cookies, ok := result["Cookies"].([]*http.Cookie)
	if !ok || len(cookies) != 1 || cookies[0].Name != "session" || cookies[0].Value != "abc123" {
		t.Errorf("Expected Cookies=[{session abc123}], got %v", cookies)
	}
}

// Test request with a custom context
func TestDumpRequestContext(t *testing.T) {
	req := newTestRequest(http.MethodGet, testUrl, "")
	ctx := context.WithValue(context.Background(), userID, 42)
	req = req.WithContext(ctx)

	result := debug.DumpRequest(req)

	if result["Context"] == nil {
		t.Errorf("Expected non-nil context, got %v", result["Context"])
	}
}

// Test request with URL containing no scheme or host
func TestDumpRequestEmptyURL(t *testing.T) {
	req := newTestRequest(http.MethodGet, "/local/path", "")

	result := debug.DumpRequest(req)

	urlData, ok := result["URL"].(map[string]interface{})
	if !ok || urlData["Path"] != "/local/path" {
		t.Errorf("Expected URL[Path]=/local/path, got %v", urlData)
	}
}
