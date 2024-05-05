package repository

import (
	"context"
	"log"

	"github.com/manish-neemnarayan/go-mongodb-server/service"
	"github.com/manish-neemnarayan/go-mongodb-server/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDBRepo struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoDBRepo(client *mongo.Client, dbname string, collName string) service.DBServicer {
	coll := client.Database(dbname).Collection(collName)

	return &mongoDBRepo{
		client: client,
		coll:   coll,
	}
}

func (m *mongoDBRepo) Fetch(filters types.FilterOptions) *[]types.FetchResponse {
	ctx := context.Background()

	pipeline := []bson.M{
		{
			"$match": bson.M{
				"$and": []bson.M{
					{"createdAt": bson.M{"$gte": filters.StartDate}},
					{"createdAt": bson.M{"$lte": filters.EndDate}},
				},
			},
		},
		{
			"$addFields": bson.M{
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"$and": []bson.M{
					{"totalCount": bson.M{"$gte": filters.MinCount}},
					{"totalCount": bson.M{"$lt": filters.MaxCount}},
				},
			},
		},
		{"$project": bson.M{
			"key":        1,
			"createdAt":  1,
			"totalCount": 1,
		}},
	}

	cur, err := m.coll.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatalf("Error while fetching data: %v", err)
	}
	defer cur.Close(ctx)

	var results []types.FetchResponse
	if err := cur.All(ctx, &results); err != nil {
		log.Fatalf("Error while getting data: %v", err)
	}

	return &results
}
