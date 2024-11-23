package params

type CheckoutOrderReq struct {
	CartIDs []int `json:"cart_ids" validate:"required"`
}

type GetOrderRes struct {
	OrderID     int                 `json:"order_id"`
	Status      string              `json:"status"`
	TotalPrice  float64             `json:"total_price"`
	CreatedAt   string              `json:"created_at"`
	UpdatedAt   string              `json:"updated_at"`
	OrderDetail []GetOrderDetailRes `json:"order_details"`
}

type GetOrderDetailRes struct {
	OrderDetailID int     `json:"order_detail_id"`
	ProductID     int     `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Price         float64 `json:"price"`
	Qty           int     `json:"qty"`
	Total         float64 `json:"total"`
	CreatedAt     string  `json:"created_at,omitempty"`
	UpdatedAt     string  `json:"updated_at,omitempty"`
}
