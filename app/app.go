package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/Piramind/Testforgo/internal/config"
	"github.com/Piramind/Testforgo/internal/tokens"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	pathConf string
)

func init() {
	flag.StringVar(&pathConf, pathConf, "config/conf.yaml", "path to configuration file")
	flag.Parse()
}

func main() {
	config, err := config.NewConfig(pathConf)
	if err != nil {
		panic(err)
	}

	db, err := mongo.NewClient(options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		panic(err)
	}
	err = db.Connect(context.Background())

	if err != nil {
		panic(err)
	}

	token := tokens.NewToken(db, config.SecretKey)

	router := gin.Default()
	handler := tokens.NewHandler(token)
	handler.Register(router)
	start(router)
}

func start(router *gin.Engine) {
	fmt.Println("Start server at :8080")
	router.Run(":8080")
}
