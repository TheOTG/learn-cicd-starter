package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header1 := make(http.Header)
	header2 := make(http.Header)
	header3 := make(http.Header)
	header4 := make(http.Header)

	header1.Set("Authorization", "ApiKey test123")
	header3.Set("Authorization", "Bearer aasd")
	header4.Set("Authorization", "ApiKey")

	tests := []struct {
		name    string
		header  http.Header
		wantKey string
		wantErr bool
	}{
		{
			name:    "Correct header",
			header:  header1,
			wantKey: "test123",
			wantErr: false,
		},
		{
			name:    "Wrong header",
			header:  header2,
			wantErr: true,
		},
		{
			name:    "Malformed header",
			header:  header3,
			wantErr: true,
		},
		{
			name:    "Fail test on purpose",
			header:  header4,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
			}
			if key != tt.wantKey {
				t.Errorf("GetAPIKey() key = %v, wantKey %v", key, tt.wantKey)
			}
		})
	}
}
