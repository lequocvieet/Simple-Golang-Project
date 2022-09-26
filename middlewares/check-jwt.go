package middlewares

import (
	"errors"
	jwt "golang_standard/auth"
	res "golang_standard/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func CheckJwt(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := jwt.Verify(r)

		if err != nil {
			res.ERROR(w, 401, errors.New("Unauthorized"))
			return
		}

		next(w, r)
	}
}
