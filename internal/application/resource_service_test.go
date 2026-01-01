package application

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/linskybing/platform-go/internal/domain/resource"
	"github.com/linskybing/platform-go/internal/repository"
	"github.com/linskybing/platform-go/internal/repository/mock"
	"github.com/linskybing/platform-go/pkg/utils"
)

func setupResourceMocks(t *testing.T) (*ResourceService,
	*mock.MockResourceRepo,
	*mock.MockAuditRepo,
	*gin.Context) {

	ctrl := gomock.NewController(t)
	t.Cleanup(func() { ctrl.Finish() })

	mockResource := mock.NewMockResourceRepo(ctrl)
	mockAudit := mock.NewMockAuditRepo(ctrl)

	repos := &repository.Repos{
		Resource: mockResource,
		Audit:    mockAudit,
	}

	svc := NewResourceService(repos)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	utils.LogAuditWithConsole = func(c *gin.Context, action, resourceType, resourceID string, oldData, newData interface{}, msg string, repos repository.AuditRepo) {
	}
	utils.GetUserIDFromContext = func(c *gin.Context) (uint, error) { return 1, nil }

	return svc, mockResource, mockAudit, c
}

func TestResourceCRUD(t *testing.T) {
	svc, mockResource, _, c := setupResourceMocks(t)

	// ----- Create -----
	res := &resource.Resource{
		RID:  1,
		Type: "Pod",
		Name: "res1",
	}
	mockResource.EXPECT().CreateResource(res).Return(nil)

	created, err := svc.CreateResource(c, res)
	if err != nil || created.RID != 1 {
		t.Fatalf("CreateResource failed: %v", err)
	}

	// ----- Get -----
	mockResource.EXPECT().GetResourceByID(uint(1)).Return(res, nil)
	got, err := svc.GetResource(1)
	if err != nil || got.RID != 1 {
		t.Fatalf("GetResource failed: %v", err)
	}

	// ----- Update -----
	newName := "res2"
	updateDTO := resource.ResourceUpdateDTO{Name: &newName}
	mockResource.EXPECT().GetResourceByID(uint(1)).Return(res, nil)
	mockResource.EXPECT().UpdateResource(res).Return(nil)

	updated, err := svc.UpdateResource(c, 1, updateDTO)
	if err != nil || updated.Name != "res2" {
		t.Fatalf("UpdateResource failed: %v", err)
	}

	// ----- Delete -----
	mockResource.EXPECT().GetResourceByID(uint(1)).Return(res, nil)
	mockResource.EXPECT().DeleteResource(uint(1)).Return(nil)

	if err := svc.DeleteResource(c, 1); err != nil {
		t.Fatalf("DeleteResource failed: %v", err)
	}
}

func TestResourceList(t *testing.T) {
	svc, mockResource, _, _ := setupResourceMocks(t)

	t.Run("ListResourcesByConfigFileID", func(t *testing.T) {
		resources := []resource.Resource{{RID: 1, Name: "r1"}}
		mockResource.EXPECT().ListResourcesByConfigFileID(uint(10)).Return(resources, nil)

		res, err := svc.ListResourcesByConfigFileID(10)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(res) != 1 {
			t.Fatalf("expected 1 resource, got %d", len(res))
		}
	})

	t.Run("ListResourcesByProjectID", func(t *testing.T) {
		resources := []resource.Resource{{RID: 1, Name: "r1"}}
		mockResource.EXPECT().ListResourcesByProjectID(uint(20)).Return(resources, nil)

		res, err := svc.ListResourcesByProjectID(20)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(res) != 1 {
			t.Fatalf("expected 1 resource, got %d", len(res))
		}
	})
}
