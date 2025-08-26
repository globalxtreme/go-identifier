package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/globalxtreme/go-identifier/data"
	"net/http"
)

func PartnerIdentifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var decryption DecryptionIdentifier

		decryption.Token = r.Header.Get("IDENTIFIER")
		if len(decryption.Token) == 0 {
			xtremeres.ErrXtremeUnauthenticated("IDENTIFIER not found")
		}

		partner := data.PartnerIdentifierData{}
		partnerData := decryption.Decrypt()
		err := json.Unmarshal(partnerData, &partner)
		if err != nil {
			xtremeres.ErrXtremeUnauthenticated(fmt.Sprintf("Unable to decode employee data json: %s", err))
		}

		ctx := context.WithValue(r.Context(), "IDENTIFIER", partner)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
