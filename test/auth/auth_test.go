package auth

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	authhttp "github.com/Lemuriets/diary/internal/auth/http"
)

var authHandler *authhttp.Handler = authhttp.NewHandler(&UsecaseMock{})

func TestHandler_signIn(t *testing.T) {
	tests := []struct {
		name     string
		login    string
		password string
		want     map[string]string
	}{
		{
			name:     "valid",
			login:    "1",
			password: "1",
			want: map[string]string{
				"access":  "access token",
				"refresh": "refresh token",
			},
		},
	}

	handler := http.HandlerFunc(authHandler.SignIn)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "auth/sign-in", nil)
			// need to add post-data later

			if err != nil {
				t.Fatal(err)
			}

			// I'll to do it later, because with testify
			handler.ServeHTTP(rec, req)
			if !reflect.DeepEqual(rec.Body.Bytes(), tc.want) {
				t.Errorf("")
			}
		})
	}
}
