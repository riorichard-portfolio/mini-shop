package fiberhttp

type ProductItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BrowseProductsResp []ProductItem
