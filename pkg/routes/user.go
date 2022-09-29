package routes

import (
	"context"
	"log"

	"github.com/lerryjay/auth-grpc-service/pkg/helpers"
	"github.com/lerryjay/auth-grpc-service/pkg/pb"
	models "github.com/lerryjay/auth-grpc-service/pkg/pb/model"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*models.User, error) {

	req.User.Id = uuid.New().String()
	hashPassword, err := helpers.HashPassword(req.User.Password)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal,
			"Could not generate new user password hash")
	}

	req.User.Password = hashPassword
	userOrm, err := req.User.ToORM(ctx)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal,
			"Unable to create user. DB failed to insert %s", err)
	}

	query := h.DB.Create(&userOrm)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.Internal,
			"Unable to create user. DB failed to insert")
	}

	return req.User, nil

}

func (h *Handler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*models.User, error) {

	var user models.UserORM
	query := h.DB.First(&user, "id = ?", req.Id)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.NotFound,
			"User not found")
	}

	userData, err := user.ToPB(ctx)
	if err != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.Internal,
			"Unable to convert user ")
	}

	return &userData, nil

}

func (h *Handler) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {

	// TODO: Implement search filters
	var usersORMList []*models.UserORM
	var users []*models.User

	query := h.DB.Find(&usersORMList)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.Internal,
			"Unable to find users ")
	}

	for _, obj := range usersORMList {
		userObj, err := obj.ToPB(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal,
				"Could not convert user %s", err)
		}

		users = append(users, &userObj)
	}

	return &pb.ListUsersResponse{
		Users: users,
		Limit: 10,
		Page:  1,
	}, nil

}

func (h *Handler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*models.User, error) {

	var user models.UserORM
	query := h.DB.First(&user, "id = ?", req.User.Id)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.NotFound,
			"User not found")
	}

	userData, err := req.User.ToORM(ctx)
	if err != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.NotFound,
			"User not found")
	}

	h.DB.Save(userData)

	return req.User, nil

}

func (h *Handler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {

	// TODO: Implement soft Delete
	query := h.DB.Delete(&models.UserORM{}, req.Id)
	if query.Error != nil {
		log.Println(query.Error)
		return &emptypb.Empty{}, status.Errorf(codes.NotFound,
			"User not found!")
	}

	return &emptypb.Empty{}, nil

}
