package service

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	pb "github.com/venomuz/project3/UserService/genproto"
	l "github.com/venomuz/project3/UserService/pkg/logger"
	cl "github.com/venomuz/project3/UserService/service/grpc_client"
	"github.com/venomuz/project3/UserService/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  cl.GrpcClientI
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger, client cl.GrpcClientI) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	id1 := uuid.NewV4()
	id2 := uuid.NewV4()
	req.Id = id1.String()
	req.Address.Id = id2.String()
	user, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("Error while inserting user info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert user")
	}
	return user, err
}
func (s *UserService) GetByID(ctx context.Context, req *pb.GetIdFromUserID) (*pb.User, error) {
	user, err := s.storage.User().GetByID(req.Id)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while getting user info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert user")
	}

	return user, err
}
func (s *UserService) DeleteByID(ctx context.Context, req *pb.GetIdFromUserID) (*pb.GetIdFromUserID, error) {
	user, err := s.storage.User().DeleteByID(req.Id)
	if err != nil {
		s.logger.Error("Error while getting user info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert user")
	}
	return user, err
}
func (s *UserService) GetAllByUserId(ctx context.Context, req *pb.GetIdFromUser) (*pb.Post, error) {
	post, err := s.client.PostService().PostGetByID(ctx, req)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while getting post info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert post")
	}

	return post, err
}
func (s *UserService) GetAllUserFromDb(ctx context.Context, req *pb.Empty) (*pb.AllUser, error) {
	users, err := s.storage.User().GetAllUserFromDb(req)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while getting post info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert post")
	}

	user := users.Users
	for _, usr := range user {
		aa := pb.GetIdFromUser{}
		aa.Id = usr.Id
		post, err := s.client.PostService().PostGetAllPosts(ctx, &aa)
		if err != nil {
			fmt.Println(err)
			s.logger.Error("Error while getting post info", l.Error(err))
			return nil, status.Error(codes.Internal, "Error insert post")
		}
		usr.Posts = post.Posts

	}
	users.Users = user

	return users, err
}
