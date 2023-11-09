package request

type ProductCategoryRequest struct {
	ParentCategoryId int    `json:"parent_category_id" validate:"required"`
	Name             string `json:"name" validate:"required"`
	Slug             string `json:"slug"`
}