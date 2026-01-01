package application

import (
	"github.com/linskybing/platform-go/internal/config"
	"github.com/linskybing/platform-go/internal/repository"
	"github.com/linskybing/platform-go/pkg/utils"
)

type AdminService struct {
	Repos *repository.Repos
}

func NewAdminService(repos *repository.Repos) *AdminService {
	return &AdminService{Repos: repos}
}

// EnsureAllUserPV creates PV/PVC for all users who do not have one
func (s *AdminService) EnsureAllUserPV() (int, error) {
	users, err := s.Repos.User.GetAllUsers()
	if err != nil {
		return 0, err
	}
	created := 0
	for _, user := range users {
		pvName := "pv-user-" + user.Username
		pvcName := "pvc-user-" + user.Username
		// TODO: check if PV/PVC already exists (call k8s API)
		// If not exists, create
		errPV := utils.CreatePV(pvName, config.DefaultStorageClassName, config.UserPVSize, user.Username)
		errPVC := utils.CreatePVC("default", pvcName, config.DefaultStorageClassName, config.UserPVSize)
		if errPV == nil && errPVC == nil {
			created++
		}
	}
	return created, nil
}
