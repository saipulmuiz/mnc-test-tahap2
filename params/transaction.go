package params

type Topup struct {
	Amount float64 `json:"amount" validate:"required"`
}

type TopupResponse struct {
	TopupID       string  `json:"top_up_id"`
	AmountTopup   float64 `json:"amount_top_up"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	CreatedDate   string  `json:"created_date"`
}

type GetTransactionRes struct {
	ProductID    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	Description  string  `json:"description"`
	ProductPrice float64 `json:"product_price"`
	Stock        int     `json:"stock"`
	MinStock     int     `json:"min_stock"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
