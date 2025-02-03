package debug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

// Dumps the underlying fields and values of a struct
func DumpStruct(s any) {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}

// Dumps the JSON body as string
func DumpJsonBody(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	rawJSON := string(body)

	fmt.Println("Raw JSON request body:", rawJSON)
}

// Dumps the request body as string
func DumpRequestBody(r *http.Request) {
	buf, _ := io.ReadAll(r.Body)
	fmt.Println("Request body:", string(buf))
	r.Body = io.NopCloser(bytes.NewBuffer(buf))
}

// Dumps the request body as a map
func RequestBodyMap(w http.ResponseWriter, r *http.Request) {
	var dataMap map[string]any
	err := json.NewDecoder(r.Body).Decode(&dataMap)
	if err != nil {
		fmt.Printf("JSON decoding error: %v\n", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	fmt.Printf("Decoded map: %+v\n", dataMap)
}

// DumpRequest extracts fields and method-returned values from *http.Request
func DumpRequest(req *http.Request) map[string]any {
	result := make(map[string]any)

	// Basic request info
	result["Method"] = req.Method
	result["URL"] = map[string]any{
		"Scheme":   req.URL.Scheme,
		"Opaque":   req.URL.Opaque,
		"User":     req.URL.User,
		"Host":     req.URL.Host,
		"Path":     req.URL.Path,
		"RawQuery": req.URL.RawQuery,
		"Fragment": req.URL.Fragment,
	}
	result["Proto"] = req.Proto
	result["ProtoMajor"] = req.ProtoMajor
	result["ProtoMinor"] = req.ProtoMinor
	result["Header"] = req.Header
	result["Host"] = req.Host
	result["RemoteAddr"] = req.RemoteAddr
	result["RequestURI"] = req.RequestURI
	result["ContentLength"] = req.ContentLength
	result["TransferEncoding"] = req.TransferEncoding
	result["Close"] = req.Close
	result["Trailer"] = req.Trailer

	// Query parameters
	result["QueryParams"] = req.URL.Query()

	// Parse form data
	if err := req.ParseForm(); err == nil {
		result["Form"] = req.Form
		result["PostForm"] = req.PostForm
	}

	// Multipart form (manually included)
	result["MultipartForm"] = req.MultipartForm

	// Call methods for unexported fields
	result["Cookies"] = req.Cookies() // Fetches stored cookies
	result["Context"] = req.Context() // Fetches context

	// TLS info (if available)
	result["TLS"] = req.TLS

	// Placeholder for the body
	result["Body"] = "<nil>"

	// Response and Cancel fields (unexported)
	result["Response"] = req.Response
	result["Cancel"] = "<nil>" // Deprecated in Go 1.17, always nil

	return result
}
