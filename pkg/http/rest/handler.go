package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/bobrnor/highloadcup2018/pkg/filtering"
)

type response struct {
	Accounts []map[string]interface{} `json:"accounts"`
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
		var filters []filtering.Filter
		for key, value := range r.Form {
			if key == "query_id" {
				continue
			}

			if key == "limit" {
				l, err := strconv.Atoi(value[0])
				if err != nil {
					panic("bad limit")
				}
				limit = l
				continue
			}

			filter, err := filtering.Make(key, value[0])
			if err != nil {
				panic(err.Error())
			}

			filters = append(filters, filter)
		}

		accounts, err := f.Fetch(filters, limit)
		if err != nil {
			panic(err.Error())
		}

		resp := response{
			Accounts: make([]map[string]interface{}, 0, len(accounts)),
		}

		for _, acc := range accounts {
			accountMap := map[string]interface{}{
				"id":    acc.ID,
				"email": acc.Email,
			}

			for _, f := range filters {
				f.Fill(acc, accountMap)
			}

			resp.Accounts = append(resp.Accounts, accountMap)
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			panic(err.Error())
		}
	}
}
