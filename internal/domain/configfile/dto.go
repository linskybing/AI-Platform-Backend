package configfile

type ConfigFileUpdateDTO struct {
	Filename *string `form:"filename"`
	RawYaml  *string `form:"raw_yaml"`
}

type CreateConfigFileInput struct {
	Filename  string `form:"filename" binding:"required"`
	RawYaml   string `form:"raw_yaml" binding:"required"`
	ProjectID uint   `form:"project_id" binding:"required"`
}

// GetProjectID returns the project ID for GID lookup
func (d CreateConfigFileInput) GetProjectID() uint {
	return d.ProjectID
}

type ProjectGetter interface {
	GetGroupIDByProjectID(projectID uint) uint
}
