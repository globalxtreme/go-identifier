package middleware

import (
	"github.com/globalxtreme/go-identifier/data"
	"github.com/globalxtreme/gobaseconf/response/error"
	"net/http"
)

func SuperadminAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		employee := r.Context().Value("IDENTIFIER").(data.EmployeeIdentifierData)
		if !employee.Superadmin {
			error.ErrXtremeUnauthenticated("Your access must be superadmin")
		}

		next.ServeHTTP(w, r)
	})
}
