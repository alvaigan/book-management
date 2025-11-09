package dto

type CreateBook struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	PublisherId int    `json:"publisher_id" validate:"required"`
	AuthorId    int    `json:"author_id" validate:"required"`
}

type UpdateBook struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	PublisherId int    `json:"publisher_id" validate:"required"`
	AuthorId    int    `json:"author_id" validate:"required"`
}
