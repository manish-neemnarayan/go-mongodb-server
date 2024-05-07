package service

import (
	"fmt"

	"github.com/manish-neemnarayan/go-mongodb-server/types"
)

type DBServicer interface {
	Fetch(filters types.FilterOptions) []types.FetchResponse
}

type DBService struct {
	client DBServicer
}

func NewDBService(client DBServicer) *DBService {
	return &DBService{
		client: client,
	}
}

func (d *DBService) FetchRecords(filters types.FilterOptions) *types.UserResponse {
	res := d.client.Fetch(filters)

	fmt.Printf("response is : %+v", res)
	return &types.UserResponse{
		Code:    0,
		Msg:     "successful",
		Records: res,
	}
}
