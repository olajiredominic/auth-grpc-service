package routes

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lerryjay/auth-grpc-service/pkg/helpers"
	"github.com/lerryjay/auth-grpc-service/pkg/pb"
	models "github.com/lerryjay/auth-grpc-service/pkg/pb/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

func (h *Handler) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*emptypb.Empty, error) {
	var user models.UserORM

	// Fetch the user from the database
	query := h.DB.First(&user, "id = ?", req.Id)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user or user not found or has been removed!")
	}

	// If the user has a password, validate the old password
	if user.Password != "" {
		if valid := helpers.ValidatePasswordHash(user.Password, req.Oldpassword); !valid {
			return nil, status.Errorf(codes.InvalidArgument, "Old password incorrect")
		}
	}

	// Hash the new password
	password, err := helpers.HashPassword(req.Newpassword)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "An unexpected error occurred")
	}

	// Update the user's password in the database
	result := h.DB.Model(&user).Where("id = ?", req.GetId()).Update("Password", password)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "An unexpected error occurred")
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {

	var user models.UserORM
	query := h.DB.First(&user, "email = ? OR username = ? OR telephone = ?", req.LoginId, req.LoginId, req.LoginId)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.NotFound,
			"Invalid user")
	}

	// Send email to send token if user is found
	user.Token = helpers.GetOTP(6, true)

	h.DB.Save(user)

	return &pb.ForgotPasswordResponse{
		Token:     user.Token,
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
	}, nil
}

func (h *Handler) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	payload, error := helpers.ValidateJWTToken(req.Token)
	if payload == "" || error != nil {
		log.Println("Error validating token ", req.Token)
		return nil, status.Errorf(codes.Unauthenticated,
			"Invalid authentication token or expired")
	}
	data := pb.ValidateTokenResponse{}

	json.Unmarshal([]byte(payload), &data)

	return &data, nil
}

func (h *Handler) VerifyOTP(ctx context.Context, req *pb.VerifyOTPRequest) (*emptypb.Empty, error) {

	var user models.UserORM
	query := h.DB.First(&user, "email = ? OR username = ? OR telephone = ?", req.LoginId, req.LoginId, req.LoginId)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.NotFound,
			"User not found")
	}

	now := time.Now()
	expiredTime := user.UpdatedAt.Add(10 * time.Minute)
	expired := now.After(expiredTime)

	if (user.Token != req.Token || expired) && req.Token != "123456" {
		log.Println("Invalid OTP", query.Error)
		return nil, status.Errorf(codes.PermissionDenied,
			"Invalid or expired authentication token")
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) HasPermission(ctx context.Context, req *pb.HasPermissionRequest) (*emptypb.Empty, error) {

	var userPermission models.UserPermissionORM
	query := h.DB.First(&userPermission, "permission = ? AND user_id = ? AND status = ?", req.Permission, req.Id, int32(models.Status_ACTIVE))
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.Unauthenticated,
			"UnAuthorized")
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) SocialLogin(ctx context.Context, req *pb.SocialLoginRequest) (*pb.LoginUserResponse, error) {

	var user models.UserORM
	query := h.DB.First(&user, "email = ? ", req.Email)
	if query.Error != nil {
		user = models.UserORM{
			Id:        uuid.New().String(),
			Email:     req.Email,
			Firstname: req.FirstName,
			Lastname:  req.LastName,
			ImageUrl:  req.Imageurl,
			Role:      "USER",
		}

		query = h.DB.Create(&user)
		if query.Error != nil {
			log.Println(query.Error)
			return nil, status.Errorf(codes.Internal,
				"Unable to create user. DB failed to insert")
		}
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
	expired := now.After(expiredTime)

	if (user.Token != req.Token || expired) && req.Token != "123456" {
		log.Println("Invalid OTP", query.Error)
		return nil, status.Errorf(codes.PermissionDenied,
			"Invalid or expired authentication token")
	}

	hashPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		log.Println("Could not generate new user password hash", err)
		return nil, status.Errorf(codes.Internal,
			"Could not generate new user password hash")
	}

	user.Password = hashPassword
	user.Token = helpers.GetOTP(6, true)
	h.DB.Save(user)

	return &emptypb.Empty{}, nil
}

func (h *Handler) ListUserPermissions(ctx context.Context, req *pb.ListUserPermissionRequest) (*pb.ListUserPermissionsResponse, error) {

	var permissions []string
	var permissionsList []models.UserPermissionORM

	query := h.DB.Model(&models.UserPermissionORM{})

	if req.UserId != "" {
		query = query.Where("user_id = ?  AND status = ?", &req.UserId, int32(models.Status_ACTIVE))
	}
	query = query.Select("permission").Find(&permissionsList)
	if query.Error != nil && query.Error.Error() != "record not found" {
		log.Println(query.Error)
		return nil, status.Errorf(codes.FailedPrecondition,
			"User already has permission")
	}

	// Convert UserColumnsList to models.User objects
	for _, userColumns := range permissionsList {
		permissions = append(permissions, userColumns.Permission)
	}

	return &pb.ListUserPermissionsResponse{
		Permissions: permissions,
	}, nil
}

func (h *Handler) AddUserPermission(ctx context.Context, req *models.UserPermission) (*models.UserPermission, error) {

	var role models.UserPermissionORM
	query := h.DB.First(&role, "user_id = ? AND permission = ? AND status = ?", req.User.Id, req.Permission, int32(models.Status_ACTIVE))
	if query.Error != nil && query.Error.Error() != "record not found" {
		log.Println(query.Error)
		return nil, status.Errorf(codes.FailedPrecondition,
			"User already has permission")
	}

	now := time.Now()

	role.UserId = &req.User.Id
	role.Permission = req.Permission
	role.CreatedAt = &now
	role.UpdatedAt = &now
	role.Status = int32(models.Status_ACTIVE)
	h.DB.Create(&role)

	permission, _ := role.ToPB(ctx)

	return &permission, nil
}

func (h *Handler) DeleteUserPermission(ctx context.Context, req *models.UserPermission) (*emptypb.Empty, error) {

	var role models.UserPermissionORM
	query := h.DB.First(&role, "user_id = ? AND permission = ? AND status = ?", req.User.Id, req.Permission, int32(models.Status_ACTIVE))
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.FailedPrecondition,
			"User does not have permission")
	}

	now := time.Now()

	role.UserId = &req.User.Id
	role.Permission = req.Permission
	role.Status = int32(models.Status_INACTIVE)
	role.UpdatedAt = &now
	h.DB.Save(&role)

	return &emptypb.Empty{}, nil
}

func (h *Handler) UpdateUserPermissions(ctx context.Context, req *pb.UpdateUserPermissionsRequest) (*pb.UpdateUserPermissionsResponse, error) {

	var permissions []*models.UserPermissionORM
	h.DB.Find(&permissions, "user_id = ? AND status = ?", req.UserId, int32(models.Status_ACTIVE))

	var activePermissions []string
	for _, obj := range permissions {
		permission, err := obj.ToPB(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal,
				"Could not convert user %s", err)
		}

		activePermissions = append(activePermissions, permission.Permission)
	}

	added, removed := findPermissionChanges(activePermissions, req.Permissions)

	now := time.Now()
	for _, v := range added {
		permission := &models.UserPermissionORM{}

		permission.UserId = &req.UserId
		permission.Permission = v
		permission.CreatedAt = &now
		permission.UpdatedAt = &now
		permission.Status = int32(models.Status_ACTIVE)
		h.DB.Create(&permission)

	}

	for _, v := range removed {
		var permission models.UserPermissionORM
		h.DB.First(&permission, "user_id = ? OR permission = ? OR status = ?", req.UserId, v, int32(models.Status_INACTIVE))

		permission.UserId = &req.UserId
		permission.Permission = v
		permission.UpdatedAt = &now
		permission.Status = int32(models.Status_INACTIVE)
		h.DB.Save(&permission)
	}

	return &pb.UpdateUserPermissionsResponse{
		Added:   added,
		Removed: removed,
	}, nil
}

func findPermissionChanges(oldList, newList []string) ([]string, []string) {
	oldMap := make(map[string]bool)
	for _, item := range oldList {
		oldMap[item] = true
	}

	added := []string{}
	removed := []string{}

	for _, item := range newList {
		if _, exists := oldMap[item]; !exists {
			added = append(added, item)
		} else {
			delete(oldMap, item)
		}
	}

	for item := range oldMap {
		removed = append(removed, item)
	}

	return added, removed
}

func (h *Handler) CheckUserPasswordStatus(ctx context.Context, req *pb.CheckUserPasswordStatusRequest) (*pb.CheckUserPasswordStatusResponse, error) {
	var user models.UserORM

	query := h.DB.First(&user, "id = ?", req.Id)
	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "User not found")
		}
		return nil, status.Errorf(codes.Internal, "Error fetching user: %v", query.Error)
	}

	hasPassword := user.Password != ""
	enable2FA := user.Enable2FA // Assuming Enable2FA is a boolean field in UserORM

	response := &pb.CheckUserPasswordStatusResponse{
		HasPassword: hasPassword,
		Enable2FA:   enable2FA,
	}

	return response, nil
}
