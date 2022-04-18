package api

type CreateProductResponse struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float64 `json:"price"`
}
