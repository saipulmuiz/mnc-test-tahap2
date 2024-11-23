package params

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload,omitempty"`
}

type ResponseSuccessLogin struct {
	Status       string `json:"status"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ResponseSuccessMessage struct {
	Message string `json:"message"`
}

type ResponseFailMessage struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type ReponseUnauthorized struct {
	Message     string `json:"message"`
	Error       string `json:"error"`
	IsLoggedOut bool   `json:"is_logged_out,omitempty"`
}

type ResponseSuccessLogout struct {
	Message     string `json:"message"`
	IsLoggedOut bool   `json:"is_logged_out"`
}

type ResponseErrorMessage struct {
	Error string `json:"error"`
}

type ResponseUrl struct {
	Url string `json:"url"`
}

type PaginationDefault struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalPage   int `json:"total_page"`
	Total       int `json:"total"`
}

type PaginationResponse struct {
	CurrentPage  int   `json:"current_page"`
	PageSize     int   `json:"page_size"`
	TotalCount   int64 `json:"total_count"`
	TotalPages   int   `json:"total_pages"`
	FirstPage    int   `json:"first_page"`
	NextPage     int   `json:"next_page"`
	LastPage     int   `json:"last_page"`
	CurrentCount int   `json:"current_count"`
}

type ResponseWithPagination struct {
	Message    string             `json:"message"`
	Data       interface{}        `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}

type ResponseSuccess struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
