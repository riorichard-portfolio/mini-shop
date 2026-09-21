package fiberhttp

type MakeOrderReq struct {
	ProductID  string `json:"product_id" validate:"required"`
	Quantity   int `json:"quantity" validate:"required"` 
}

type OrderListReq struct{
	Limit    int `query:"limit" validate:"required"`
	Offset   int `query:"offset" validate:"gte=0"`
}

type CompleteOrderReq struct {
	OrderID  string `json:"order_id" validate:"required"`
}

type CancelOrderReq struct {
	OrderID  string `json:"order_id" validate:"required"`
}