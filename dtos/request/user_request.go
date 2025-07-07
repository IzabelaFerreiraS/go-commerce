package request

import (
	"fmt"
	"go-commerce/utils"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email  string `json:"email"`
	Password string `json:"password"`
	Role   string  `json:"role"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" && r.Email == "" && r.Password == "" && r.Role == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Name == "" {
		return utils.ErrParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return utils.ErrParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return utils.ErrParamIsRequired("password", "string")
	}
	if r.Role == "" {
		return utils.ErrParamIsRequired("role", "string")
	}
	return nil
}

type UpdatedUserRequest struct {
	Name     string `json:"name"`
	Email  string `json:"email"`
	Password string `json:"password"`
	Role   string  `json:"role"`
}

func (r *UpdatedUserRequest) Validate() error {
	if r.Name != "" || r.Email != "" || r.Password != "" || r.Role != ""{
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}