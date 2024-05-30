package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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

	payload := map[string]interface{}{
		"firstname": req.Firstname,
		"lastname":  req.Lastname,
		// "middlename": req.Middlename,
		// "dob":        req.Dob,
		// "phone":      req.Phone,
		// "email":      req.Email,
		// "gender":     req.Gender,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to marshal request payload: %v", err)
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create HTTP request: %v", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to make HTTP request: %v", err)
	}
	defer httpResp.Body.Close()

	var nimcResp NIMCResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&nimcResp); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to decode response: %v", err)
	}

	if nimcResp.Status.Status != "verified" {
		return nil, status.Errorf(codes.InvalidArgument, "NIN verification failed")
	}

	grpcResp := &pb.VerifyNINResponse{
		Id: nimcResp.ID,
		Applicant: &pb.VerifyNINResponse_Applicant{
			Firstname: nimcResp.Applicant.Firstname,
			Lastname:  nimcResp.Applicant.Lastname,
		},
		Summary: &pb.VerifyNINResponse_Summary{
			NinCheck: &pb.VerifyNINResponse_Summary_NinCheck{
				Status: nimcResp.Summary.NINCheck.Status,
				FieldMatches: &pb.VerifyNINResponse_Summary_NinCheck_FieldMatches{
					Firstname: nimcResp.Summary.NINCheck.FieldMatches.Firstname,
					Lastname:  nimcResp.Summary.NINCheck.FieldMatches.Lastname,
				},
			},
		},
		Status: &pb.VerifyNINResponse_Status{
			State:  nimcResp.Status.State,
			Status: nimcResp.Status.Status,
		},
		Nin: &pb.VerifyNINResponse_Nin{
			Nin:        nimcResp.NIN.NIN,
			Firstname:  nimcResp.NIN.Firstname,
			Lastname:   nimcResp.NIN.Lastname,
			Middlename: nimcResp.NIN.Middlename,
			Phone:      nimcResp.NIN.Phone,
			Gender:     nimcResp.NIN.Gender,
			Birthdate:  nimcResp.NIN.Birthdate,
			Photo:      nimcResp.NIN.Photo,
			Address:    nimcResp.NIN.Address,
		},
	}

	return grpcResp, nil
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

	payload := map[string]interface{}{
		"firstname": req.Firstname,
		"lastname":  req.Lastname,
		// "middlename": req.Middlename,
		// "dob":        req.Dob,
		// "phone":      req.Phone,
		// "email":      req.Email,
		// "gender":     req.Gender,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to marshal request payload: %v", err)
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create HTTP request: %v", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))

	client := &http.Client{}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to make HTTP request: %v", err)
	}
	defer httpResp.Body.Close()

	var nimcResp NIMCResponse
	if err := json.NewDecoder(httpResp.Body).Decode(&nimcResp); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to decode response: %v", err)
	}

	if nimcResp.Status.Status != "verified" {
		return nil, status.Errorf(codes.InvalidArgument, "NIN verification failed")
	}

	grpcResp := &pb.VerifyNINResponse{
		Id: nimcResp.ID,
		Applicant: &pb.VerifyNINResponse_Applicant{
			Firstname: nimcResp.Applicant.Firstname,
			Lastname:  nimcResp.Applicant.Lastname,
		},
		Summary: &pb.VerifyNINResponse_Summary{
			NinCheck: &pb.VerifyNINResponse_Summary_NinCheck{
				Status: nimcResp.Summary.NINCheck.Status,
				FieldMatches: &pb.VerifyNINResponse_Summary_NinCheck_FieldMatches{
					Firstname: nimcResp.Summary.NINCheck.FieldMatches.Firstname,
					Lastname:  nimcResp.Summary.NINCheck.FieldMatches.Lastname,
				},
			},
		},
		Status: &pb.VerifyNINResponse_Status{
			State:  nimcResp.Status.State,
			Status: nimcResp.Status.Status,
		},
		Nin: &pb.VerifyNINResponse_Nin{
			Nin:        nimcResp.NIN.NIN,
			Firstname:  nimcResp.NIN.Firstname,
			Lastname:   nimcResp.NIN.Lastname,
			Middlename: nimcResp.NIN.Middlename,
			Phone:      nimcResp.NIN.Phone,
			Gender:     nimcResp.NIN.Gender,
			Birthdate:  nimcResp.NIN.Birthdate,
			Photo:      nimcResp.NIN.Photo,
			Address:    nimcResp.NIN.Address,
		},
	}

	return grpcResp, nil
}

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	payload := map[string]string{
		"clientId": h.ClientID,
		"secret":   h.SecretKey,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to marshal request payload: %v", err)
	}

	httpReq, err := http.NewRequest("POST", h.TokenURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create HTTP request: %v", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to make HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, status.Errorf(codes.Internal, "Received non-200 response: %v", resp.Status)
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to decode response: %v", err)
	}

	return &pb.LoginResponse{Token: authResp.AccessToken}, nil

}
