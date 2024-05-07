package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/manish-neemnarayan/go-mongodb-server/service"
	"github.com/manish-neemnarayan/go-mongodb-server/types"
)

func MemoryDBHandler(svc *service.MemoryDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		switch r.Method {
		// to insert data if method is post
		case http.MethodPost:
			fmt.Println("in post memroy handler")
			var params *types.MemoryData
			if err := json.NewDecoder(body).Decode(&params); err != nil {
				http.Error(w, "requested params are wrong", http.StatusBadRequest)
				return
			}

			response, err := svc.Post(params)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}

			WriteJson(w, http.StatusOK, response)

		// to get data if method is get
		case http.MethodGet:
			key := r.URL.Query()["key"][0]

			fmt.Println(key)
			response, err := svc.Get(key)
			if err != nil {
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
				return
			}

			WriteJson(w, http.StatusOK, response)

		// bad http method
		default:
			http.Error(w, fmt.Sprintf("%s method is not allowed", r.Method), http.StatusBadGateway)
		}
	}
}
