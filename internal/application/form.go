package application

import (
	"github.com/linskybing/platform-go/internal/domain/form"
	"github.com/linskybing/platform-go/internal/repository"
)

type FormService struct {
	repo *repository.FormRepository
}

func NewFormService(repo *repository.FormRepository) *FormService {
	return &FormService{repo: repo}
}

func (s *FormService) CreateForm(userID uint, input form.CreateFormDTO) (*form.Form, error) {
	f := &form.Form{
		UserID:      userID,
		ProjectID:   input.ProjectID,
		Title:       input.Title,
		Description: input.Description,
		Status:      form.FormStatusPending,
	}
	return f, s.repo.Create(f)
}

func (s *FormService) GetAllForms() ([]form.Form, error) {
	return s.repo.FindAll()
}

func (s *FormService) GetUserForms(userID uint) ([]form.Form, error) {
	return s.repo.FindByUserID(userID)
}

func (s *FormService) UpdateFormStatus(id uint, status string) (*form.Form, error) {
	f, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	f.Status = form.FormStatus(status)
	return f, s.repo.Update(f)
}
