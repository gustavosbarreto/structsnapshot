package example

//go:generate structsnapshot User
type User struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Timezone string `json:"timezone" validate:"timezone"`
}
