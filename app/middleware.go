package app

import (
	"net/http"

	"github.com/Lemuriets/diary/pkg/httpjson"
	"github.com/Lemuriets/diary/pkg/token"
)

func Authorization(handler http.HandlerFunc, permissions int8) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authValue := r.Header["Authorization"]

		if len(authValue) == 0 {
			httpjson.UnauthorizedResponse(w)
			return
		}
		accessToken, err := token.GetToken(authValue[0])

		if !token.CheckValid(accessToken) || err != nil {
			httpjson.UnauthorizedResponse(w)
			return
		}
		claims := token.GetPayload(accessToken)
		if claims == nil {
			httpjson.UnauthorizedResponse(w)
			return
		}
		receivedPerm, _ := claims["userPermissions"].(float64)
		if int8(receivedPerm) > permissions {
			httpjson.ForbiddenResponse(w)
			return
		}

		handler(w, r)
	}
}
