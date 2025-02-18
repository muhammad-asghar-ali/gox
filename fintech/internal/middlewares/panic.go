package middlewares

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/types"
)

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			error := recover()

			if error != nil {
				log.Println(error)

				resp := types.ErrResponse{
					Code:    http.StatusInternalServerError,
					Message: "Internal server error",
				}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(resp)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}
