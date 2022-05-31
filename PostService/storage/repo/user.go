package repo

import (
	pb "github.com/venomuz/project3/PostService/genproto"
)

//PostStorageI ...
type PostStorageI interface {
	PostCreate(*pb.Post) (*pb.OkBOOL, error)
	PostGetByID(ID string) (*pb.Post, error)
	PostDeleteByID(ID string) (*pb.OkBOOL, error)
}
