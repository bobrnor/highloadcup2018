package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/bobrnor/highloadcup2018/pkg/filtering"
)

type response struct {
	accounts []filtering.Account `json:"accounts"`
}

func New(f filtering.Service) http.Handler {
	router := httprouter.New()

	router.GET("/accounts/filter/", filter(f))

	return router
}

func filter(f filtering.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if err := r.ParseForm(); err != nil {
			panic("can't parse form")
		}

		var limit int
		var filters []filtering.FieldFilter
		for key, value := range r.Form {
			if key == "limit" {
				l, err := strconv.Atoi(value[0])
				if err != nil {
					panic("bad limit")
				}
				limit = l
				continue
			}

			parts := strings.Split(key, "_")
			if len(parts) != 2 {
				panic("bad key")
			}

			filters = append(filters, filtering.FieldFilter{
				Field:     parts[0],
				Operation: parts[1],
				Value:     value[0],
			})
		}

		accounts, err := f.Fetch(filters, limit)
		if err != nil {
			panic(err.Error())
		}

		if err := json.NewEncoder(w).Encode(response{accounts: accounts}); err != nil {
			panic(err.Error())
		}
	}
}
