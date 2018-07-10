package main

import (
	"context"

	pb "github.com/ichiaohsu/readr-micro/users/proto"
)

type service struct {
	user User
}

func (s *service) Get(ctx context.Context, req *pb.GetRequest, res *pb.Users) error {
	// Call UserInterface Get()
	users, err := s.user.Get(req)
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}
