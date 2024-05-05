package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	db "github.com/manish-neemnarayan/go-mongodb-server/config"
	"github.com/manish-neemnarayan/go-mongodb-server/repository"
	"github.com/manish-neemnarayan/go-mongodb-server/service"
	"github.com/manish-neemnarayan/go-mongodb-server/types"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	uri      = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true"
	dbname   = "getircase-study"
	collName = "records"
)

func main() {
	var (
		client    *mongo.Client
		mongoRepo service.DBServicer
		svc       *service.DBService
	)

	client = db.NewDB(uri)
	mongoRepo = repository.NewMongoDBRepo(client, dbname, collName)
	svc = service.NewDBService(mongoRepo)
	_ = svc

	http.HandleFunc("/health", HealthHandler())
	http.HandleFunc("/mongo", MongoHandler(svc))

	fmt.Println("Server started on port :9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}

func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := isPostHTTPMethod(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		} else {
			writeJson(w, 200, map[string]any{"code": 0, "msg": "healthy server", "records": []int{}})
		}
	}
}

func MongoHandler(svc *service.DBService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := isPostHTTPMethod(r)
		query := r.URL.Query()

		if err != nil {
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)

		} else {
			filters, err := extractFilters(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			res := svc.FetchRecords(*filters)

			writeJson(w, http.StatusOK, res)
		}

	}
}

// /// utility functions --------------------------------
func writeJson(w http.ResponseWriter, statusCode int, msg any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(msg)
}

func isPostHTTPMethod(r *http.Request) error {
	switch r.Method {
	case http.MethodPost:
		return nil
	default:
		return fmt.Errorf("method: %s not allowed", r.Method)
	}
}

func extractFilters(query url.Values) (filters *types.FilterOptions, err error) {
	filters = &types.FilterOptions{}

	if minCountStr := query["minCount"]; minCountStr[0] != "" {
		minCount, err := strconv.Atoi(minCountStr[0])
		if err != nil {
			return nil, fmt.Errorf("invalid minCount param: %v", err)
		}
		filters.MinCount = minCount
	}

	if maxCountStr := query["maxCount"]; maxCountStr[0] != "" {
		maxCount, err := strconv.Atoi(maxCountStr[0])
		if err != nil {
			return nil, fmt.Errorf("invalid maxCount param: %v", err)
		}
		filters.MaxCount = maxCount
	}

	if startDateStr := query["startDate"]; startDateStr[0] != "" {
		startDate, err := time.Parse(time.DateOnly, startDateStr[0])
		if err != nil {
			return nil, fmt.Errorf("invalid start date param: %v", err)
		}
		filters.StartDate = startDate
	}

	if endDateStr := query["endDate"]; endDateStr[0] != "" {
		endDate, err := time.Parse(time.DateOnly, endDateStr[0])
		if err != nil {
			return nil, fmt.Errorf("invalid end date param: %v", err)
		}
		filters.EndDate = endDate
	}

	return filters, nil
}
