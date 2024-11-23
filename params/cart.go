package params

type AddProductToCart struct {
	ProductID int `json:"product_id" validate:"required"`
	Qty       int `json:"qty" validate:"required"`
}

type UpdatedCartReq struct {
	Qty int `json:"qty" validate:"required"`
}

type GetCartRes struct {
	CartID      int     `json:"cart_id"`
	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
	Total       float64 `json:"total"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}
