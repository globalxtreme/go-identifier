package middleware

import (
	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"github.com/globalxtreme/go-identifier/data"
	"net/http"
)

func SuperadminAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		employee := r.Context().Value("IDENTIFIER").(data.EmployeeIdentifierData)
		if !employee.Superadmin {
			xtremeres.ErrXtremeUnauthenticated("Your access must be superadmin")
		}

		next.ServeHTTP(w, r)
	})
}
