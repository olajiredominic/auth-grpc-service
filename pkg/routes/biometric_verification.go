package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lerryjay/auth-grpc-service/pkg/helpers"
	"github.com/lerryjay/auth-grpc-service/pkg/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) VerifyNINImage(ctx context.Context, req *pb.VerifyIdentityImageRequest) (*pb.VerifyIdentityImageResponse, error) {
	// Call the Login function to get the bearer token
	token, err := h.Login(ctx, &pb.LoginRequest{
		ClientId:  h.ClientID,
		SecretKey: h.SecretKey,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to login: %v", err)
	}

	// Get the image verification URL part from the idType
	imageVerificationType, err := helpers.GetImageVerificationURL(req.IdType)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid identity type: %v", err)
	}

	// Fetch the baseURL from the config
	baseURL := h.BiometricQoreidBaseURL

	// Construct the URL for NIN image verification
	url := fmt.Sprintf("%s%s", baseURL, imageVerificationType)

	// Create the payload for NIN image verification
	payload := map[string]string{
		"photoUrl":    req.PhotoUrl,
		"photoBase64": req.PhotoBase64,
	}

	// Create the headers with the authorization token
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token.Token),
	}

	// Make the POST request for NIN image verification
	resp, err := helpers.PostRequest(url, payload, header)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, status.Errorf(codes.Internal, "Received non-200 response: %v", resp.Status)
	}

	// Decode the response into the VerifyIdentityImageResponse
	var imageVerificationResp pb.VerifyIdentityImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&imageVerificationResp); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to decode response: %v", err)
	}

	// Check if the verification status is "verified"
	if imageVerificationResp.Status.Status != "verified" {
		return nil, status.Errorf(codes.InvalidArgument, "NIN image verification failed")
	}

	// Return the successful verification response
	return &imageVerificationResp, nil
}
