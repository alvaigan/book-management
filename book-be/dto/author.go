package dto

type CreateAuthor struct {
	Name string `json:"name" validate:"required"`
}

type UpdateAuthor struct {
	Name string `json:"name" validate:"required"`
}
