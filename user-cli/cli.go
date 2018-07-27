package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	pb "github.com/ichiaohsu/readr-micro/users/proto"
// 	microclient "github.com/micro/go-micro/client"
// 	"github.com/micro/go-micro/cmd"
// )

// func main() {
// 	cmd.Init()

// 	client := pb.NewUserServiceClient("go.micro.srv.users", microclient.DefaultClient)
// 	params := pb.GetRequest{Id: 648}
// 	r, err := client.Get(context.TODO(), &params)
// 	if err != nil {
// 		log.Fatalf("Cannot list user: %v\n", err)
// 	}
// 	for _, v := range r.Users {
// 		fmt.Println(v)
// 	}
// }
