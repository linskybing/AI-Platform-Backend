package user

// Repository defines data access interface for users
type Repository interface {
	Create(user *User) error
	GetByID(uid uint) (*User, error)
	FindByID(uid uint) (*User, error) // Alias for GetByID
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	List() ([]User, error)
	Update(user *User) error
	UpdateStatus(uid uint, status UserStatus) error
	Delete(uid uint) error
	GetUsersByRole(role UserRole) ([]User, error)
}
