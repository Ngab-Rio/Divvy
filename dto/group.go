package dto

import (
	"time"
)

type GroupResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Created_by string `json:"created_by"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type CreateGroupRequest struct {
	Name string `json:"name" validate:"required,min=4,max=100"`
}

type GroupWithUserResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Created_by UserResponse `json:"created_by"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}