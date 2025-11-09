package dto

type CreatePublisher struct {
	Name string `json:"name" validate:"required"`
	City string `json:"city" validate:"required"`
}

type UpdatePublisher struct {
	Name string `json:"name" validate:"required"`
	City string `json:"city" validate:"required"`
}
