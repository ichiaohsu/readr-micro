package main

import (
	"log"

	pb "github.com/ichiaohsu/readr-micro/users/proto"
	"github.com/jmoiron/sqlx"
)

type UserInterface interface {
	Get() ([]*pb.User, error)
}

type User struct {
	db *sqlx.DB
}

func (u *User) Get(req *pb.GetRequest) (result []*pb.User, err error) {
	err = u.db.Select(&result, u.db.Rebind(`SELECT id, name, nickname FROM members WHERE id = ?;`), req.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Printf("length of result:%v", len(result))
	return result, nil
}
