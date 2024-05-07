package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/manish-neemnarayan/go-mongodb-server/types"
)

// /// utility functions --------------------------------
func WriteJson(w http.ResponseWriter, statusCode int, msg any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(msg)
}

func IsPostHTTPMethod(r *http.Request) error {
	switch r.Method {
	case http.MethodPost:
		return nil
	default:
		return fmt.Errorf("method: %s not allowed", r.Method)
	}
}

func ExtractFilters(query url.Values) (filters *types.FilterOptions, err error) {
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
