package handler

import (
	"net/http"

	"github.com/manish-neemnarayan/go-mongodb-server/service"
)

func MongoHandler(svc *service.DBService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := IsPostHTTPMethod(r)
		query := r.URL.Query()

		if err != nil {
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)

		} else {
			filters, err := ExtractFilters(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			res := svc.FetchRecords(*filters)

			WriteJson(w, http.StatusOK, res)
		}

	}
}
