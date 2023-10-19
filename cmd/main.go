package main

import (
	"github.com/nskntc/CodePix/application/grpc"
	"github.com/nskntc/CodePix/infrasctructure/db"
	"github.com/jinzhu/gorm"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}