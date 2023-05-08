package server

type SuccessResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

type PaginationResponseMixin struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}
