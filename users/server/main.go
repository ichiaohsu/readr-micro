package main

import (
	"fmt"
	"log"

	pb "github.com/ichiaohsu/readr-micro/users/server/proto"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
)

func main() {

	db, err := Connect("host=127.0.0.1 port=5432 user=ich dbname=readr sslmode=disable")
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not establish database connection:%v\n", err)
	}
	u := User{db}
	// host := os.Getenv("DB_HOST")
	// if host == ""
	// 	host = "localhost:5432"
	// }
	srv := grpc.NewService(
		micro.Name("go.micro.srv.users"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{u})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
