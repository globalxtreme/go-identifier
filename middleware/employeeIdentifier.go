package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/globalxtreme/go-identifier/data"
	errResponse "github.com/globalxtreme/gobaseconf/response/error"
	"net/http"
)

func EmployeeIdentifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var decryption DecryptionIdentifier

		decryption.Token = r.Header.Get("IDENTIFIER")
		if len(decryption.Token) == 0 {
			errResponse.ErrXtremeUnauthenticated("IDENTIFIER not found")
		}

		employeeData := decryption.Decrypt()
		err := json.Unmarshal(employeeData, &data.Employee)
		if err != nil {
			errResponse.ErrXtremeUnauthenticated(fmt.Sprintf("Unable to decode employee data json: %s", err))
		}

		decryption.Token = r.Header.Get("ACCESS")
		if len(decryption.Token) == 0 {
			errResponse.ErrXtremeUnauthenticated("ACCESS not found")
		}

		accessData := decryption.Decrypt()
		err = json.Unmarshal(accessData, &data.Access)
		if err != nil {
			errResponse.ErrXtremeUnauthenticated(fmt.Sprintf("Unable to decode access data json: %s", err))
		}

		next.ServeHTTP(w, r)
	})
}
