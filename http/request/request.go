package request

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

// Decodes the JSON payload from the request body
func JSON[T any](r *http.Request) (T, error) {
	var v T
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

// getIPAddress extracts the client's IP address from the request.
func GetIPAddress(r *http.Request) string {
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}

	if forwardedFor := r.Header.Values("X-Forwarded-For"); len(forwardedFor) > 0 {
		firstIP := forwardedFor[0]
		ips := strings.Split(firstIP, ",")
		var validIPs []string

		for _, ip := range ips {
			trimmedIP := strings.TrimSpace(ip)
			if trimmedIP != "" {
				validIPs = append(validIPs, trimmedIP)
			}
		}

		if len(validIPs) > 0 {
			return validIPs[0]
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}
