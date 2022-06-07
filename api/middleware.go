package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func (api *API) AllowOrigin(w http.ResponseWriter, req *http.Request) {
	// localhost:9000 origin mendapat ijin akses
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// allow cookie
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
}

func (api *API) AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.AllowOrigin(w, r)
		encoder := json.NewEncoder(w)
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada

		 cookie, err := r.Cookie("token")
		 if err != nil {
		 	encoder.Encode(AuthErrorResponse{err.Error()})
		 	if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
		 		return
		 	}
		 	w.WriteHeader(http.StatusBadRequest)
		 	return
		}

		// Task: Ambil value dari cookie token

		value := cookie.Value

		// Task: Deklarasi variable claim

		claim := Claims{}

		// Task: 1. parse JWT token ke dalam claim
		//       2. return unauthorized ketika signature invalid
		//       3. return bad request ketika field token tidak ada
		//       4. return unauthorized ketika token sudah tidak valid (biasanya karna token expired)

		cookies, err := jwt.ParseWithClaims(value, &claim, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		});
		if err != nil {
			encoder.Encode(AuthErrorResponse{err.Error()})
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !cookies.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(map[string]string{"message": "unauthorized"})
			return
		}

		// Task: Validasi

		cntx := context.WithValue(r.Context(), "user", claim)
		next.ServeHTTP(w, r.WithContext(cntx))
	})
}