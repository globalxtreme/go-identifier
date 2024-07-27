package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/globalxtreme/go-identifier/data"
	"net/http"
)

func MyGXIdentifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var decryption DecryptionIdentifier

		decryption.Token = r.Header.Get("IDENTIFIER")
		if len(decryption.Token) == 0 {
			xtremeres.ErrXtremeUnauthenticated("IDENTIFIER not found")
		}

		myGX := data.MyGXIdentifierData{}
		mygxData := decryption.Decrypt()
		err := json.Unmarshal(mygxData, &myGX)
		if err != nil {
			xtremeres.ErrXtremeUnauthenticated(fmt.Sprintf("Unable to decode My GX data json: %s", err))
		}

		ctx := context.WithValue(r.Context(), "IDENTIFIER", myGX)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
