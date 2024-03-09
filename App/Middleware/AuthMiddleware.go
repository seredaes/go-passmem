package Middleware

import (
	"net/http"
	"seredaes/go-passmem/App/JWT"
	"seredaes/go-passmem/App/Response"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// ---------------------------------------------------------------------
		// Check Content-Type in header
		// ---------------------------------------------------------------------
		contentType := r.Header.Get("Content-Type")
		if strings.Compare(contentType, "application/json") != 0 {
			Response.RenderResponse(w, false, "Content type: application/json is required", nil, 422)
			return
		}

		// ---------------------------------------------------------------------
		// CHECK IF ROUTE PATH CAN BE EXCLUDED
		// ---------------------------------------------------------------------
		route := r.URL.Path
		var disabledAuth []string = []string{"/api/registration", "/api/auth"}
		for _, routePath := range disabledAuth {
			if strings.Compare(routePath, route) == 0 {
				next.ServeHTTP(w, r)
				return
			}
		}

		// ---------------------------------------------------------------------
		// CHECK TOKEN IF ROUTE IS NOT EXCLUDED
		// ---------------------------------------------------------------------
		token := r.Header.Get("Authorization")

		if token == "" {
			Response.RenderResponse(w, false, "Auth is required", nil, 401)
			return
		}

		token = strings.Replace(token, "Bearer ", "", -1)
		if !JWT.CheckJWT(token) {
			Response.RenderResponse(w, false, "Auth is required", nil, 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}
