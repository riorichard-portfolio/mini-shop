package fiberhttp

type AddNewProductReq struct {
	Name  string `json:"name" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
}

type BrowseProductsReq struct {
	Limit  int `query:"limit" validate:"required"`
	Offset int `query:"offset" validate:"gte=0"`
}
