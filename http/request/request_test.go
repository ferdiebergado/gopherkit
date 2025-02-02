package request_test

import (
	"net/http/httptest"
	"testing"

	"github.com/ferdiebergado/gopherkit/http/request"
)

func TestGetIPAddress(t *testing.T) {
	tests := []struct {
		name          string
		xRealIP       string
		xForwardedFor string
		remoteAddr    string
		expectedIP    string
	}{
		{
			name:       "X-Real-IP",
			xRealIP:    "192.168.1.100",
			remoteAddr: "10.0.0.1:1234", // Should be ignored
			expectedIP: "192.168.1.100",
		},
		{
			name:          "X-Forwarded-For (single)",
			xForwardedFor: "192.168.2.200",
			remoteAddr:    "10.0.0.2:5678", // Should be ignored
			expectedIP:    "192.168.2.200",
		},
		{
			name:          "X-Forwarded-For (multiple)",
			xForwardedFor: "192.168.2.200, 10.0.0.3",
			remoteAddr:    "10.0.0.2:5678", // Should be ignored
			expectedIP:    "192.168.2.200", // First IP is chosen
		},
		{
			name:          "X-Forwarded-For (multiple with whitespace)",
			xForwardedFor: " 192.168.2.200 , 10.0.0.3 ", // with whitespace
			remoteAddr:    "10.0.0.2:5678",              // Should be ignored
			expectedIP:    "192.168.2.200",              // First IP is chosen, whitespace trimmed
		},
		{
			name:       "RemoteAddr",
			remoteAddr: "10.0.0.4:9012",
			expectedIP: "10.0.0.4",
		},
		{
			name:       "No Headers",
			remoteAddr: "10.0.0.5:3456",
			expectedIP: "10.0.0.5",
		},
		{
			name:          "X-Forwarded-For empty",
			xForwardedFor: "",
			remoteAddr:    "10.0.0.6:7890",
			expectedIP:    "10.0.0.6",
		},
		{
			name:          "X-Forwarded-For with spaces only",
			xForwardedFor: "   ",
			remoteAddr:    "10.0.0.7:1357",
			expectedIP:    "10.0.0.7",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/", nil)
			if tt.xRealIP != "" {
				req.Header.Set("X-Real-IP", tt.xRealIP)
			}
			if tt.xForwardedFor != "" {
				req.Header.Set("X-Forwarded-For", tt.xForwardedFor)
			}
			req.RemoteAddr = tt.remoteAddr

			ip := request.GetIPAddress(req)
			if ip != tt.expectedIP {
				t.Errorf("got IP %q, want %q for test case %q", ip, tt.expectedIP, tt.name)
			}
		})
	}
}
