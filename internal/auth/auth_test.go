package auth

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		key       string
		value     string
		expect    string
		errExpect string
	}{
		{
			errExpect: "no authorization header included",
		},
		{
			key:       "smth",
			value:     "smth",
			errExpect: "no authorization header included",
		},
		{
			key:    "Authorization",
			value:  "ApiKey FFFFFFF",
			expect: "FFFFFFF",
		},
		{
			key:       "Authorization",
			value:     "ApiKey",
			errExpect: "malformed authorization header",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("TestGetAPIKey Case #%d:", i), func(t *testing.T) {
			header := http.Header{}
			header.Set(tt.key, tt.value)

			got, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), tt.errExpect) {
					return
				}
				t.Errorf("Unexpected: TestGetApiKey:%v\n", err)
				return
			}

			if got != tt.expect {
				t.Errorf("Unexpected: TestGetApiKey:%v\n", got)
				return
			}
		})
	}
}
