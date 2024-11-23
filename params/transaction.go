package params

type TopupRequest struct {
	Amount float64 `json:"amount" validate:"required"`
}

type PaymentRequest struct {
	Amount  float64 `json:"amount" validate:"required"`
	Remarks string  `json:"remarks"`
}

type TransferRequest struct {
	TargetUser string  `json:"target_user" validate:"required"`
	Amount     float64 `json:"amount" validate:"required"`
	Remarks    string  `json:"remarks"`
}

type TopupResponse struct {
	TopupID       string  `json:"top_up_id"`
	AmountTopup   float64 `json:"amount_top_up"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	CreatedDate   string  `json:"created_date"`
}

type PaymentResponse struct {
	PaymentID     string  `json:"payment_id"`
	Amount        float64 `json:"amount"`
	Remarks       string  `json:"remarks"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	CreatedDate   string  `json:"created_date"`
}

type TransferResponse struct {
	TransferID    string  `json:"transfer_id"`
	Amount        float64 `json:"amount"`
	Remarks       string  `json:"remarks"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	CreatedDate   string  `json:"created_date"`
}

type GetTransactionRes struct {
	Status          string  `json:"status"`
	UserID          string  `json:"user_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	Remarks         string  `json:"remarks"`
	BalanceBefore   float64 `json:"balance_before"`
	BalanceAfter    float64 `json:"balance_after"`
	CreatedDate     string  `json:"created_date"`
}
