package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/ichiaohsu/readr-micro/users/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	os.Setenv("MICRO_SERVER_ADDRESS", ":50051")
	os.Setenv("MICRO_REGISTRY", "mdns")

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
	srv := micro.NewService(
		micro.Name("go.micro.srv.users"),
		micro.Version("latest"),
	)

	srv.Init()

	pb.RegisterUserServiceHandler(srv.Server(), &service{u})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
