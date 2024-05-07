package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	db "github.com/manish-neemnarayan/go-mongodb-server/config"
	"github.com/manish-neemnarayan/go-mongodb-server/handler"
	"github.com/manish-neemnarayan/go-mongodb-server/repository"
	"github.com/manish-neemnarayan/go-mongodb-server/service"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbname   = "getircase-study"
	collName = "records"
)

func main() {
	var (
		client    *mongo.Client
		mongoRepo service.DBServicer
		svc       *service.DBService
	)

	//env setup
	envVar := flag.String("uri", "", "Mongodb URI value")
	flag.Parse()
	os.Setenv("uri_env", *envVar)
	uri := os.Getenv("uri_env")

	//connection to mongodb
	client = db.NewDB(uri)
	mongoRepo = repository.NewMongoDBRepo(client, dbname, collName)
	svc = service.NewDBService(mongoRepo) // service for mongo db
	memorySVC := service.NewMemoryDB()    // service for in-memory db

	http.HandleFunc("/health", HealthHandler())
	http.HandleFunc("/mongo", handler.MongoHandler(svc))
	http.HandleFunc("/in-memory", handler.MemoryDBHandler(memorySVC))

	fmt.Println("Server started on port :9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}

func HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler.IsPostHTTPMethod(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		} else {
			handler.WriteJson(w, 200, map[string]any{"code": 0, "msg": "healthy server", "records": []int{}})
		}
	}
}
