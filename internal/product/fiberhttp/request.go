package fiberhttp

type AddNewProductReq struct {
	Name  string `json:"name" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
}

type BrowseProductsReq struct {
	Limit  int `json:"limit" validate:"required"`
	Offset int `json:"offset" validate:"required"`
}
