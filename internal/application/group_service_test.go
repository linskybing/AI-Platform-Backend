package application_test

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/linskybing/platform-go/internal/application"
	"github.com/linskybing/platform-go/internal/domain/group"
	"github.com/linskybing/platform-go/internal/repository"
	"github.com/linskybing/platform-go/internal/repository/mock"
	"github.com/linskybing/platform-go/pkg/utils"
)

func setupGroupMocks(t *testing.T) (*application.GroupService, *mock.MockGroupRepo, *mock.MockAuditRepo, *gin.Context) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() { ctrl.Finish() })

	mockGroup := mock.NewMockGroupRepo(ctrl)
	mockAudit := mock.NewMockAuditRepo(ctrl)
	repos := &repository.Repos{
		Group: mockGroup,
		Audit: mockAudit,
	}

	svc := application.NewGroupService(repos)
	c, _ := gin.CreateTestContext(nil)

	// Mock audit log globally
	utils.LogAuditWithConsole = func(c *gin.Context, action, resourceType, resourceID string, oldData, newData interface{}, msg string, repos repository.AuditRepo) {
	}

	return svc, mockGroup, mockAudit, c
}

func TestGroupServiceCRUD(t *testing.T) {
	svc, mockGroup, _, c := setupGroupMocks(t)

	t.Run("ListGroups success", func(t *testing.T) {
		mockGroup.EXPECT().GetAllGroups().Return([]group.Group{
			{GID: 1, GroupName: "dev"},
			{GID: 2, GroupName: "ops"},
		}, nil)

		groups, err := svc.ListGroups()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(groups) != 2 {
			t.Fatalf("expected 2 groups, got %d", len(groups))
		}
	})

	t.Run("GetGroup success", func(t *testing.T) {
		mockGroup.EXPECT().GetGroupByID(uint(1)).Return(group.Group{GID: 1, GroupName: "dev"}, nil)
		group, err := svc.GetGroup(1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if group.GroupName != "dev" {
			t.Fatalf("expected dev, got %s", group.GroupName)
		}
	})

	t.Run("GetGroup not found", func(t *testing.T) {
		mockGroup.EXPECT().GetGroupByID(uint(99)).Return(group.Group{}, errors.New("not found"))
		_, err := svc.GetGroup(99)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("CreateGroup success", func(t *testing.T) {
		input := group.GroupCreateDTO{GroupName: "qa", Description: nil}
		mockGroup.EXPECT().CreateGroup(gomock.Any()).Return(nil)

		group, err := svc.CreateGroup(c, input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if group.GroupName != "qa" {
			t.Fatalf("expected qa, got %s", group.GroupName)
		}
	})

	t.Run("CreateGroup reserved name", func(t *testing.T) {
		input := group.GroupCreateDTO{GroupName: "super"}
		_, err := svc.CreateGroup(c, input)
		if !errors.Is(err, application.ErrReservedGroupName) {
			t.Fatalf("expected reserved group name error, got %v", err)
		}
	})

	t.Run("UpdateGroup success", func(t *testing.T) {
		old := group.Group{GID: 1, GroupName: "dev"}
		mockGroup.EXPECT().GetGroupByID(uint(1)).Return(old, nil)
		mockGroup.EXPECT().UpdateGroup(gomock.Any()).Return(nil)

		newName := "devops"
		newDesc := "updated description"
		input := group.GroupUpdateDTO{GroupName: &newName, Description: &newDesc}
		group, err := svc.UpdateGroup(c, 1, input)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if group.GroupName != "devops" || group.Description != "updated description" {
			t.Fatalf("unexpected updated group %+v", group)
		}
	})

	t.Run("UpdateGroup reserved name", func(t *testing.T) {
		old := group.Group{GID: 1, GroupName: "dev"}
		mockGroup.EXPECT().GetGroupByID(uint(1)).Return(old, nil)
		newName := "super"
		input := group.GroupUpdateDTO{GroupName: &newName}
		_, err := svc.UpdateGroup(c, 1, input)
		if !errors.Is(err, application.ErrReservedGroupName) {
			t.Fatalf("expected reserved group name error, got %v", err)
		}
	})

	t.Run("UpdateGroup not found", func(t *testing.T) {
		mockGroup.EXPECT().GetGroupByID(uint(99)).Return(group.Group{}, errors.New("not found"))
		newName := "newname"
		input := group.GroupUpdateDTO{GroupName: &newName}
		_, err := svc.UpdateGroup(c, 99, input)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("DeleteGroup success", func(t *testing.T) {
		existing := group.Group{GID: 1, GroupName: "dev"}
		mockGroup.EXPECT().GetGroupByID(uint(1)).Return(existing, nil)
		mockGroup.EXPECT().DeleteGroup(uint(1)).Return(nil)

		err := svc.DeleteGroup(c, 1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("DeleteGroup not found", func(t *testing.T) {
		mockGroup.EXPECT().GetGroupByID(uint(99)).Return(group.Group{}, errors.New("not found"))
		err := svc.DeleteGroup(c, 99)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}
