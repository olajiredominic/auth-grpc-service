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

func (h *Handler) CreateUser(ctx context.Context, req *models.User) (*models.User, error) {
	var user models.UserORM
	//  Checks auth fields exist
	checkUser := h.DB.First(&user, "email = ? OR (username = ? AND username != '') OR telephone = ?", req.Email, req.Username, req.Telephone)
	if checkUser.Error == nil {
		log.Println("Tried creating user. User exists")
		return nil, status.Errorf(codes.AlreadyExists,
			"User exists")
	}

	req.Id = uuid.New().String()
	hashPassword, err := helpers.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal,
			"Could not generate new user password hash")
	}

	req.Password = hashPassword
	userOrm, err := req.ToORM(ctx)
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

	return req, nil
}

func (h *Handler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*models.User, error) {
	// Define a struct to hold the specific columns you want to retrieve
	// type UserColumns struct {
	// 	Id        string
	// 	Email     string
	// 	Firstname string
	// 	Lastname  string
	// 	Role      string
	// 	ImaageUrl string
	// 	Username  string
	// }

	var user models.UserORM
	query := h.DB.Select("id, email, firstname, lastname, role, image_url, username, bio, verification_status").First(&user, "id = ?", req.Id)
	if query.Error != nil {
		log.Println(query.Error)
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	// Convert the UserColumns struct to a pb.User object
	userData := &models.User{
		Id:                 user.Id,
		Email:              user.Email,
		Firstname:          user.Firstname,
		Lastname:           user.Lastname,
		Role:               user.Role,
		ImageUrl:           user.ImageUrl,
		Username:           user.Username,
		Bio:                user.Bio,
		VerificationStatus: models.VerificationStatus(user.VerificationStatus),
	}

	return userData, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	// TODO: Implement search filters
	var usersORMList []*models.UserORM
	var users []*models.User
	var totalCount int64
	var totalPages int64
	var nextPage int64

	// Set up pagination options
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Page == 0 {
		req.Page = 1
	}

	offset := (req.Page - 1) * req.Limit
	query := h.DB.Model(&models.UserORM{})

	if req.Role != nil && *req.Role != "" {
		query = query.Where("role = ?", req.Role)
	}

	if req.Name != nil && *req.Name != "" {
		query = query.Where("firstname LIKE ? OR lastname LIKE ?", "%"+*req.Name+"%", "%"+*req.Name+"%")
	}

	if req.Email != nil && *req.Email != "" {
		query = query.Where("email  LIKE ? ", "%"+*req.Email+"%")
	}
	// Query total count of users
	if err := query.Count(&totalCount).Error; err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal,
			"Could not convert user %s", err)
	}

	// Calculate total pages
	totalPages = totalCount / int64(req.Limit)
	if totalCount%int64(req.Limit) != 0 {
		totalPages++
	}

	query.Offset(int(offset)).Limit(int(req.Limit)).Find(&usersORMList)
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

	if req.Page < int32(totalPages) {
		nextPage = int64(req.Page) + 1
	} else {
		nextPage = 1
	}

	return &pb.ListUsersResponse{
		Users: users,
		Meta: &models.Meta{
			Limit:       req.Limit,
			CurrentPage: req.Page,
			TotalCount:  int32(totalCount),
			TotalPages:  int32(totalPages),
			NextPage:    int32(nextPage),
		},
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *models.User) (*models.User, error) {

	var user models.UserORM
	query := h.DB.First(&user, "id = ?", req.Id)
	if query.Error != nil {
		log.Println("Error fetching user ", req.Id, query.Error)
		return nil, status.Errorf(codes.NotFound,
			"User not found")
	}

	userData, err := req.ToORM(ctx)
	if err != nil {
		log.Println("Unable to convert to ORM ", req, err)
		return nil, status.Errorf(codes.NotFound,
			"User not found")
	}

	userData.Password = user.Password
	userData.CreatedAt = user.CreatedAt
	userData.UpdatedAt = user.UpdatedAt
	userData.Role = user.Role
	userData.Email = user.Email
	h.DB.Save(&userData)

	return req, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*emptypb.Empty, error) {

	// TODO: Implement soft Delete
	query := h.DB.Where("id = ?", req.GetId()).Delete(&models.UserORM{})
	if query.Error != nil {
		log.Println(query.Error)
		return &emptypb.Empty{}, status.Errorf(codes.NotFound,
			"User not found!")
	}

	return &emptypb.Empty{}, nil

}

func (h *Handler) VerifyUser(ctx context.Context, req *models.UserVerification) (*emptypb.Empty, error) {
	// Check if user with the provided ID already exists
	var existingUser models.UserORM
	query := h.DB.First(&existingUser, "id = ?", req.User)
	if query.Error != nil {
		log.Println("User not found for verification:", query.Error)
		return nil, status.Errorf(codes.NotFound, "User not found for verification")
	}

	// Perform additional verification logic based on IdType
	switch req.IdType {
	case models.IdType_DRIVERS_LICENCE:
		// Perform verification logic for driver's license
		// ...

	case models.IdType_PASSPORT:
		// Perform verification logic for passport
		// ...

	case models.IdType_IDENTITY_CARD:
		// Perform verification logic for identity card
		// ...

	default:
		// Invalid IdType
		return nil, status.Errorf(codes.InvalidArgument, "Invalid IdType")
	}

	// Perform additional verification logic based on CountryCode
	// ...

	// Update user verification status in the database
	// For example, you might set a field like IsVerified to true in the UserORM model
	h.DB.Save(&req)
	existingUser.VerificationStatus = int32(models.VerificationStatus_VERIFIED)
	updateQuery := h.DB.Save(&existingUser)
	if updateQuery.Error != nil {
		log.Println("Failed to update user verification status:", updateQuery.Error)
		return nil, status.Errorf(codes.Internal, "Failed to update user verification status")
	}

	return &emptypb.Empty{}, nil
}
