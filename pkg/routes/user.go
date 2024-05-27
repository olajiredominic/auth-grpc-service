package routes

import (
	"context"
	"errors"
	"log"

	"github.com/lerryjay/auth-grpc-service/pkg/helpers"
	"github.com/lerryjay/auth-grpc-service/pkg/pb"
	models "github.com/lerryjay/auth-grpc-service/pkg/pb/model"
	"gorm.io/gorm"

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

func (h *Handler) UpdateUserIDImage(ctx context.Context, req *pb.UpdateIDImageRequest) (*pb.UpdateIDImageResponse, error) {
	// First, check if the user exists in UserORM
	var user models.UserORM
	query := h.DB.First(&user, "id = ?", req.UserId)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			// User does not exist in UserORM, create a new row
			user.Id = req.UserId
			// Assuming other necessary fields are set here
			if err := h.DB.Create(&user).Error; err != nil {
				log.Println("Error creating user ", req.UserId, err)
				return nil, status.Errorf(codes.Internal, "Database error")
			}
		} else {
			log.Println("Error fetching user ", req.UserId, query.Error)
			return nil, status.Errorf(codes.Internal, "Database error")
		}
	}

	// Next, check if the user exists in UserVerificationORM
	var verification models.UserVerificationORM
	query = h.DB.First(&verification, "user_id = ?", req.UserId)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			// User does not exist in UserVerificationORM, create a new row
			verification.UserId = &req.UserId
			verification.IdFilePath = req.IdImagePath
			if err := h.DB.Create(&verification).Error; err != nil {
				log.Println("Error creating user verification ", req.UserId, err)
				return nil, status.Errorf(codes.Internal, "Database error")
			}
		} else {
			log.Println("Error fetching user verification ", req.UserId, query.Error)
			return nil, status.Errorf(codes.Internal, "Database error")
		}
	} else {
		// User exists in UserVerificationORM, update the necessary fields
		if err := h.DB.Model(&verification).Where("user_id = ?", req.UserId).Updates(models.UserVerificationORM{
			IdFilePath: req.IdImagePath,
		}).Error; err != nil {
			log.Println("Error updating user verification ", req.UserId, err)
			return nil, status.Errorf(codes.Internal, "Database error")
		}
	}

	// Convert the updated UserORM model back to a Pb model
	updatedUser, err := user.ToPB(ctx)
	if err != nil {
		log.Println("Unable to convert UserORM to User model", err)
		return nil, status.Errorf(codes.Internal, "Unable to update user ID image")
	}

	// Create a response object and populate it with the necessary data
	response := &pb.UpdateIDImageResponse{
		User: &updatedUser, // Assuming pb.UpdateIDImageResponse has a User field
	}

	return response, nil
}

func (h *Handler) UpdateUserSelfie(ctx context.Context, req *pb.UpdateSelfieRequest) (*pb.UpdateSelfieResponse, error) {
	// First, check if the user exists in UserORM
	var user models.UserORM
	query := h.DB.First(&user, "id = ?", req.UserId)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			// User does not exist in UserORM, create a new row
			user.Id = req.UserId
			// Assuming other necessary fields are set here
			if err := h.DB.Create(&user).Error; err != nil {
				log.Println("Error creating user ", req.UserId, err)
				return nil, status.Errorf(codes.Internal, "Database error")
			}
		} else {
			log.Println("Error fetching user ", req.UserId, query.Error)
			return nil, status.Errorf(codes.Internal, "Database error")
		}
	} else {
		// User exists, update the user details if necessary
		// Assuming user.IDImagePath or other fields might need updates
		// if err := h.DB.Save(&user).Error; err != nil {
		// 	log.Println("Error updating user ", req.UserId, err)
		// 	return nil, status.Errorf(codes.Internal, "Database error")
		// }
	}

	// Next, check if the user exists in UserVerificationORM
	var verification models.UserVerificationORM
	query = h.DB.First(&verification, "user_id = ?", req.UserId)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			// User does not exist in UserVerificationORM, create a new row
			verification.UserId = &req.UserId
			verification.Selfie = req.SelfiePath
			if err := h.DB.Create(&verification).Error; err != nil {
				log.Println("Error creating user verification ", req.UserId, err)
				return nil, status.Errorf(codes.Internal, "Database error")
			}
		} else {
			log.Println("Error fetching user verification ", req.UserId, query.Error)
			return nil, status.Errorf(codes.Internal, "Database error")
		}
	} else {
		// User exists in UserVerificationORM, update the necessary fields
		if err := h.DB.Model(&verification).Where("user_id = ?", req.UserId).Updates(models.UserVerificationORM{
			Selfie: req.SelfiePath,
		}).Error; err != nil {
			log.Println("Error updating user verification ", req.UserId, err)
			return nil, status.Errorf(codes.Internal, "Database error")
		}
	}

	// Convert the updated UserORM model back to a Pb model
	updatedUser, err := user.ToPB(ctx)
	if err != nil {
		log.Println("Unable to convert UserORM to User model", err)
		return nil, status.Errorf(codes.Internal, "Unable to update user Selfie image")
	}

	// Create a response object and populate it with the necessary data
	response := &pb.UpdateSelfieResponse{
		User: &updatedUser, // Assuming pb.UpdateSelfieResponse has a User field
	}

	return response, nil
}

func (h *Handler) UpdateUserIDNumber(ctx context.Context, req *pb.UpdateIDNumberRequest) (*pb.UpdateIDNumberResponse, error) {
	// First, check if the user exists in UserORM
	var user models.UserORM
	query := h.DB.First(&user, "id = ?", req.UserId)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			// User does not exist in UserORM, create a new row
			user.Id = req.UserId
			// Assuming other necessary fields are set here
			if err := h.DB.Create(&user).Error; err != nil {
				log.Println("Error creating user ", req.UserId, err)
				return nil, status.Errorf(codes.Internal, "Database error")
			}
		} else {
			log.Println("Error fetching user ", req.UserId, query.Error)
			return nil, status.Errorf(codes.Internal, "Database error")
		}
	}

	// Next, check if the user exists in UserVerificationORM
	var verification models.UserVerificationORM
	query = h.DB.First(&verification, "user_id = ?", req.UserId)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			// User does not exist in UserVerificationORM, create a new row
			verification.UserId = &req.UserId
			verification.IdNumber = req.IdNumber
			if err := h.DB.Create(&verification).Error; err != nil {
				log.Println("Error creating user verification ", req.UserId, err)
				return nil, status.Errorf(codes.Internal, "Database error")
			}
		} else {
			log.Println("Error fetching user verification ", req.UserId, query.Error)
			return nil, status.Errorf(codes.Internal, "Database error")
		}
	} else {
		// User exists in UserVerificationORM, update the necessary fields
		if err := h.DB.Model(&verification).Where("user_id = ?", req.UserId).Updates(models.UserVerificationORM{
			IdNumber: req.IdNumber,
		}).Error; err != nil {
			log.Println("Error updating user verification ", req.UserId, err)
			return nil, status.Errorf(codes.Internal, "Database error")
		}
	}

	// Convert the updated UserORM model back to a Pb model
	updatedUser, err := user.ToPB(ctx)
	if err != nil {
		log.Println("Unable to convert UserORM to User model", err)
		return nil, status.Errorf(codes.Internal, "Unable to update user ID number")
	}

	// Create a response object and populate it with the necessary data
	response := &pb.UpdateIDNumberResponse{
		User: &updatedUser, // Assuming pb.UpdateIDNumberResponse has a User field
	}

	return response, nil
}
