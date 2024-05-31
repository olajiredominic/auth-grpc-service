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

type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int32  `json:"expiresIn"`
	TokenType   string `json:"tokenType"`
}

type NIMCResponse struct {
	ID        int32 `json:"id"`
	Applicant struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"applicant"`
	Summary struct {
		NINCheck struct {
			Status       string `json:"status"`
			FieldMatches struct {
				Firstname bool `json:"firstname"`
				Lastname  bool `json:"lastname"`
			} `json:"fieldMatches"`
		} `json:"nin_check"`
	} `json:"summary"`
	Status struct {
		State  string `json:"state"`
		Status string `json:"status"`
	} `json:"status"`
	NIN struct {
		NIN        string `json:"nin"`
		Firstname  string `json:"firstname"`
		Lastname   string `json:"lastname"`
		Middlename string `json:"middlename"`
		Phone      string `json:"phone"`
		Gender     string `json:"gender"`
		Birthdate  string `json:"birthdate"`
		Photo      string `json:"photo"`
		Address    string `json:"address"`
	} `json:"nin"`
}

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	payload := map[string]string{
		"clientId": h.ClientID,
		"secret":   h.SecretKey,
	}

	resp, _ := helpers.PostRequest(h.TokenURL, payload, nil)
	if resp.StatusCode != http.StatusCreated {
		return nil, status.Errorf(codes.Internal, "Received non-200 response: %v", resp.Status)
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to decode response: %v", err)
	}

	defer resp.Body.Close()

	return &pb.LoginResponse{Token: authResp.AccessToken}, nil

}

// VerifyNIN verifies the NIN with the provided details
func (h *Handler) VerifyNIN(ctx context.Context, req *pb.VerifyNINRequest) (*pb.VerifyNINResponse, error) {
	// Call the Login function to get the bearer token
	token, err := h.Login(ctx, &pb.LoginRequest{
		ClientId:  h.ClientID,
		SecretKey: h.SecretKey,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to login: %v", err)
	}

	// Fetch the baseURL and vninURL from the config
	baseURL := h.QoreidBaseURL
	ninURL := h.NINURL

	// Construct the URL using the values from config
	url := fmt.Sprintf("%s%s/%s", baseURL, ninURL, req.IdNumber)

	payload := map[string]string{
		"firstname": req.Firstname,
		"lastname":  req.Lastname,
		// "middlename": req.Middlename,
		// "dob":        req.Dob,
		// "phone":      req.Phone,
		// "email":      req.Email,
		// "gender":     req.Gender,
	}
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token.Token),
	}
	resp, _ := helpers.PostRequest(url, payload, header)
	if resp.StatusCode != http.StatusCreated {
		return nil, status.Errorf(codes.Internal, "Received non-200 response: %v", resp.Status)
	}

	var nimcResp pb.VerifyNINResponse
	if err := json.NewDecoder(resp.Body).Decode(&nimcResp); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to decode response: %v", err)
	}

	defer resp.Body.Close()

	if nimcResp.Status.Status != "verified" {
		return nil, status.Errorf(codes.InvalidArgument, "NIN verification failed")
	}

	return &nimcResp, nil
}

func (h *Handler) VerifyVNIN(ctx context.Context, req *pb.VerifyNINRequest) (*pb.VerifyNINResponse, error) {
	// Call the Login function to get the bearer token
	token, err := h.Login(ctx, &pb.LoginRequest{
		ClientId:  h.ClientID,
		SecretKey: h.SecretKey,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to login: %v", err)
	}

	// Fetch the baseURL and vninURL from the config
	baseURL := h.QoreidBaseURL
	VninURL := h.VNINURL

	// Construct the URL using the values from config
	url := fmt.Sprintf("%s%s/%s", baseURL, VninURL, req.IdNumber)

	payload := map[string]string{
		"firstname": req.Firstname,
		"lastname":  req.Lastname,
		// "middlename": req.Middlename,
		// "dob":        req.Dob,
		"phone": req.Phone,
		// "email":      req.Email,
		// "gender":     req.Gender,
	}
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token.Token),
	}

	resp, err := helpers.PostRequest(url, payload, header)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to decode response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, status.Errorf(codes.FailedPrecondition, "Received non-200 response: %v", resp.Status)
	}

	var nimcResp pb.VerifyNINResponse
	if err := json.NewDecoder(resp.Body).Decode(&nimcResp); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to decode response: %v", err)
	}

	defer resp.Body.Close()

	if nimcResp.Status.Status != "verified" {
		return nil, status.Errorf(codes.InvalidArgument, "NIN verification failed")
	}

	return &nimcResp, nil
}
