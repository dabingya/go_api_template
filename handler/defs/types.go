package defs

type User struct {
	Name string `json:"name" form:"name" validate:"required,gt=3"`
}