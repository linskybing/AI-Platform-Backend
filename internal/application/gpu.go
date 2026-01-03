package application

import (
	"errors"
	"time"

	"github.com/linskybing/platform-go/internal/domain/gpu"
	"github.com/linskybing/platform-go/internal/repository"
)

type GPURequestService struct {
	Repos *repository.Repos
}

func NewGPURequestService(repos *repository.Repos) *GPURequestService {
	return &GPURequestService{
		Repos: repos,
	}
}

func (s *GPURequestService) CreateRequest(projectID uint, userID uint, input gpu.CreateGPURequestDTO) (gpu.GPURequest, error) {
	req := gpu.GPURequest{
		ProjectID:   projectID,
		RequesterID: userID,
		Type:        gpu.GPURequestType(input.Type),
		Reason:      input.Reason,
		Status:      gpu.GPURequestStatusPending,
	}

	if input.RequestedQuota != nil {
		req.RequestedQuota = *input.RequestedQuota
	}
	if input.RequestedAccessType != nil {
		req.RequestedAccessType = *input.RequestedAccessType
	}

	err := s.Repos.GPURequest.Create(&req)
	return req, err
}

func (s *GPURequestService) ListByProject(projectID uint) ([]gpu.GPURequest, error) {
	return s.Repos.GPURequest.ListByProjectID(projectID)
}

func (s *GPURequestService) ListPending() ([]gpu.GPURequest, error) {
	return s.Repos.GPURequest.ListPending()
}

func (s *GPURequestService) ProcessRequest(requestID uint, status string) (gpu.GPURequest, error) {
	req, err := s.Repos.GPURequest.GetByID(requestID)
	if err != nil {
		return gpu.GPURequest{}, err
	}

	if req.Status != gpu.GPURequestStatusPending {
		return req, errors.New("request is not pending")
	}

	req.Status = gpu.GPURequestStatus(status)
	req.UpdatedAt = time.Now()

	if status == string(gpu.GPURequestStatusApproved) {
		// Update Project
		project, err := s.Repos.Project.GetProjectByID(req.ProjectID)
		if err != nil {
			return req, err
		}

		switch req.Type {
		case gpu.GPURequestTypeQuota:
			project.GPUQuota = req.RequestedQuota
		case gpu.GPURequestTypeAccess:
			project.GPUAccess = req.RequestedAccessType
		}

		if err := s.Repos.Project.UpdateProject(&project); err != nil {
			return req, err
		}
	}

	if err := s.Repos.GPURequest.Update(&req); err != nil {
		return req, err
	}

	return req, nil
}
