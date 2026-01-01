package application

import (
	"errors"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/linskybing/platform-go/internal/domain/group"
	"github.com/linskybing/platform-go/internal/repository"
	"github.com/linskybing/platform-go/pkg/utils"
)

var ErrReservedUser = errors.New("cannot modify reserved user & group 'admin & super'")

type UserGroupService struct {
	Repos *repository.Repos
}

func NewUserGroupService(repos *repository.Repos) *UserGroupService {
	return &UserGroupService{
		Repos: repos,
	}
}

func (s *UserGroupService) AllocateGroupResource(gid uint, userName string) error {
	projects, err := s.Repos.Project.ListProjectsByGroup(gid)

	if err != nil {
		return err
	}

	for _, project := range projects {
		ns := utils.FormatNamespaceName(project.PID, userName)
		if err := utils.CreateNamespace(ns); err != nil {
			return err
		}
		// if err := utils.CreatePVC(ns, config.DefaultStorageName, config.DefaultStorageClassName, config.DefaultStorageSize); err != nil {
		// 	return err
		// }
	}

	return nil
}

func (s *UserGroupService) RemoveGroupResource(gid uint, userName string) error {
	projects, err := s.Repos.Project.ListProjectsByGroup(gid)

	if err != nil {
		return err
	}

	for _, project := range projects {
		ns := utils.FormatNamespaceName(project.PID, userName)
		if err := utils.DeleteNamespace(ns); err != nil {
			return err
		}
	}

	return nil
}

func (s *UserGroupService) CreateUserGroup(c *gin.Context, userGroup *group.UserGroup) (*group.UserGroup, error) {
	if err := s.Repos.UserGroup.CreateUserGroup(userGroup); err != nil {
		return nil, err
	}

	uesrName, err := s.Repos.User.GetUsernameByID(userGroup.UID)

	if err != nil {
		return nil, err
	}

	if err := s.AllocateGroupResource(userGroup.GID, uesrName); err != nil {
		return nil, err
	}

	go utils.LogAuditWithConsole(c, "create", "user_group",
		fmt.Sprintf("u_id=%d,g_id=%d", userGroup.UID, userGroup.GID),
		nil, *userGroup, "", s.Repos.Audit)

	return userGroup, nil
}

func (s *UserGroupService) UpdateUserGroup(c *gin.Context, userGroup *group.UserGroup, existing group.UserGroup) (*group.UserGroup, error) {
	if err := s.Repos.UserGroup.UpdateUserGroup(userGroup); err != nil {
		return nil, err
	}

	go utils.LogAuditWithConsole(c, "update", "user_group",
		fmt.Sprintf("u_id=%d,g_id=%d", userGroup.UID, userGroup.GID),
		existing, *userGroup, "", s.Repos.Audit)

	return userGroup, nil
}

func (s *UserGroupService) DeleteUserGroup(c *gin.Context, uid, gid uint) error {
	oldUserGroup, err := s.Repos.UserGroup.GetUserGroup(uid, gid)
	if err != nil {
		return err
	}

	// Check if this is the admin user or super group (no GroupName in group.UserGroup, so we need to query the group)
	if uid == 1 && gid == 1 {
		return ErrReservedUser
	}

	log.Printf("sddf")
	if err := s.Repos.UserGroup.DeleteUserGroup(uid, gid); err != nil {
		return err
	}

	uesrName, err := s.Repos.User.GetUsernameByID(uid)
	if err != nil {
		return err
	}

	if err := s.RemoveGroupResource(gid, uesrName); err != nil {
		return err
	}

	go utils.LogAuditWithConsole(c, "delete", "user_group",
		fmt.Sprintf("u_id=%d,g_id=%d", uid, gid),
		oldUserGroup, nil, "", s.Repos.Audit)

	return nil
}

func (s *UserGroupService) GetUserGroup(uid, gid uint) (group.UserGroup, error) {
	return s.Repos.UserGroup.GetUserGroup(uid, gid)
}

func (s *UserGroupService) GetUserGroupsByUID(uid uint) ([]group.UserGroup, error) {
	return s.Repos.UserGroup.GetUserGroupsByUID(uid)
}

func (s *UserGroupService) GetUserGroupsByGID(gid uint) ([]group.UserGroup, error) {
	return s.Repos.UserGroup.GetUserGroupsByGID(gid)
}

func (s *UserGroupService) FormatByUID(records []group.UserGroup) map[uint]map[string]interface{} {
	result := make(map[uint]map[string]interface{})

	for _, r := range records {
		if g, exists := result[r.UID]; exists {
			// Convert to map[string]interface{} for flexibility
			g["gid"] = r.GID
			g["role"] = r.Role
			result[r.UID] = g
		} else {
			result[r.UID] = map[string]interface{}{
				"uid":  r.UID,
				"gid":  r.GID,
				"role": r.Role,
			}
		}
	}
	return result
}

func (s *UserGroupService) FormatByGID(records []group.UserGroup) map[uint]map[string]interface{} {
	result := make(map[uint]map[string]interface{})

	for _, r := range records {
		if g, exists := result[r.GID]; exists {
			g["uid"] = r.UID
			g["role"] = r.Role
			result[r.GID] = g
		} else {
			result[r.GID] = map[string]interface{}{
				"gid":  r.GID,
				"uid":  r.UID,
				"role": r.Role,
			}
		}
	}
	return result
}
