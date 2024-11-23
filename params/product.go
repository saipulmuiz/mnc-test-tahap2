package params

type CreateProductReq struct {
	ProductName  string  `json:"product_name" validate:"required"`
	Description  string  `json:"description"`
	ProductPrice float64 `json:"product_price" validate:"required"`
	Stock        int     `json:"stock" validate:"required"`
	MinStock     int     `json:"min_stock"`
}

type UpdatedProductReq struct {
	ProductName  string  `json:"product_name" validate:"required"`
	Description  string  `json:"description"`
	ProductPrice float64 `json:"product_price" validate:"required"`
	Stock        int     `json:"stock" validate:"required"`
	MinStock     int     `json:"min_stock"`
}

type UpdatedProductRes struct {
	ProductName  string  `json:"product_name"`
	Description  string  `json:"description"`
	ProductPrice float64 `json:"product_price"`
	Stock        int     `json:"stock"`
	MinStock     int     `json:"min_stock"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type GetProductRes struct {
	ProductID    int     `json:"product_id"`
	ProductName  string  `json:"product_name"`
	Description  string  `json:"description"`
	ProductPrice float64 `json:"product_price"`
	Stock        int     `json:"stock"`
	MinStock     int     `json:"min_stock"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
