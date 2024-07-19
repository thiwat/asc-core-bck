package order

type PlaceOrderInput struct {
	Event    string `json:"event"`
	Quantity int32  `json:"quantity"`
}

type UploadSlipInput struct {
	OrderNo string `json:"order_no"`
	SlipUrl string `json:"slip_url"`
}

type ApprovePaymentInput struct {
	OrderNo string `json:"order_no"`
}

type ListOutput struct {
	Rows      []Order `json:"rows"`
	Page      int64   `json:"page"`
	PageSize  int64   `json:"page_size"`
	Total     int64   `json:"total"`
	TotalPage int64   `json:"total_page"`
}
