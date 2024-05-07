package types

import "time"

type FilterOptions struct {
	MinCount  int       `json:"minCount"`
	MaxCount  int       `json:"maxCount"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

type FetchResponse struct {
	Key        string    `bson:"key" json:"key"`
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
	TotalCount int       `bson:"totalCount" json:"totalCount"`
}

type UserResponse struct {
	Code    uint            `json:"code"`
	Msg     string          `json:"msg"`
	Records []FetchResponse `json:"records"`
}

type MemoryData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
