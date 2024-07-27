package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/globalxtreme/go-identifier/data"
	"net/http"
)

func EmployeeIdentifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var decryption DecryptionIdentifier

		decryption.Token = r.Header.Get("IDENTIFIER")
		if len(decryption.Token) == 0 {
			xtremeres.ErrXtremeUnauthenticated("IDENTIFIER not found")
		}

		employee := data.EmployeeIdentifierData{}
		employeeData := decryption.Decrypt()
		err := json.Unmarshal(employeeData, &employee)
		if err != nil {
			xtremeres.ErrXtremeUnauthenticated(fmt.Sprintf("Unable to decode employee data json: %s", err))
		}

		decryption.Token = r.Header.Get("ACCESS")
		if len(decryption.Token) == 0 {
			xtremeres.ErrXtremeUnauthenticated("ACCESS not found")
		}

		access := data.AccessIdentifierData{}
		accessData := decryption.Decrypt()
		err = json.Unmarshal(accessData, &access)
		if err != nil {
			xtremeres.ErrXtremeUnauthenticated(fmt.Sprintf("Unable to decode access data json: %s", err))
		}

		ctx := context.WithValue(r.Context(), "IDENTIFIER", employee)
		ctx = context.WithValue(ctx, "ACCESS", access)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
