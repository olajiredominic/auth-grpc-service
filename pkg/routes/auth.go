package routes

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/lerryjay/auth-grpc-service/pkg/helpers"
	"github.com/lerryjay/auth-grpc-service/pkg/pb"
	models "github.com/lerryjay/auth-grpc-service/pkg/pb/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*emptypb.Empty, error) {

	var user models.UserORM

	query := h.DB.First(&user, "id = ?", req.Id)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.InvalidArgument,
			"Invalid user or user not found or has been removed!")
	}

	if valid := helpers.ValidatePasswordHash(user.Password, req.Oldpassword); !valid {
		return nil, status.Errorf(codes.InvalidArgument,
			"Old password incorrect")
	}

	password, err := helpers.HashPassword(req.Newpassword)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal,
			"An unexpected error occured")
	}

	result := h.DB.Model(&user).Where("id = ?", req.GetId()).Update("Password", password)
	// Append to the Books table
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal,
			"An unexpected error occured")
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*emptypb.Empty, error) {

	var user models.UserORM
	query := h.DB.First(&user, "email = ? OR username = ? OR telephone = ?", req.LoginId, req.LoginId, req.LoginId)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.NotFound,
			"Invalid user")
	}

	// Send email to send token if user is found
	user.Token = helpers.GetOTP(6)

	h.DB.Save(user)

	return &emptypb.Empty{}, nil
}

func (h *Handler) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {

	payload, error := helpers.ValidateJWTToken(req.Token)
	if payload == "" || error != nil {
		log.Println("Error validating token ", req.Token, error.Error())
		return nil, status.Errorf(codes.Unauthenticated,
			"Invalid authentication token or expired")
	}

	data := pb.ValidateTokenResponse{}
	json.Unmarshal([]byte(payload), &data)

	return &data, nil
}

func (h *Handler) HasPermission(ctx context.Context, req *pb.HasPermissionRequest) (*emptypb.Empty, error) {

	var userPermission models.UserPermissionsORM
	query := h.DB.First(&userPermission, "permission = ? AND userId = ?", req.Permission, req.Id)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.Unauthenticated,
			"Invalid username or password")
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	var user models.UserORM
	query := h.DB.First(&user, "email = ? OR username = ? OR telephone = ?", req.LoginId, req.LoginId, req.LoginId)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.InvalidArgument,
			"Invalid username or password")
	}

	if valid := helpers.ValidatePasswordHash(user.Password, *req.Password); !valid {
		return nil, status.Errorf(codes.InvalidArgument,
			"Invalid username or password")
	}

	claims := map[string]string{
		"Id":   user.Id,
		"Role": user.Role,
	}

	tokenString, err := helpers.GenerateToken(claims)
	if err != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.Internal,
			"Error authenticating user!")
	}

	return &pb.LoginUserResponse{
		Token:   tokenString,
		Message: "Login Successful",
	}, nil
}

func (h *Handler) ResetPassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*emptypb.Empty, error) {

	var user models.UserORM
	query := h.DB.First(&user, "email = ? OR username = ? OR telephone = ?", req.LoginId, req.LoginId, req.LoginId)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.NotFound,
			"User not found")
	}

	now := time.Now()
	expiredTime := user.UpdatedAt.Add(10 * time.Minute)
	expired := expiredTime.Before(now)

	if user.Token != req.Token || expired {
		log.Println(query.Error)
		return nil, status.Errorf(codes.PermissionDenied,
			"Invalid or expired authentication token")
	}

	hashPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal,
			"Could not generate new user password hash")
	}

	user.Password = hashPassword
	user.Token = helpers.GetOTP(6)
	h.DB.Save(user)

	return &emptypb.Empty{}, nil
}
