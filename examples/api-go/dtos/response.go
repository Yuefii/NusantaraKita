package dtos

type ApiDTO struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ResponseDTO struct {
	Pagination PaginationDTO `json:"pagination"`
	Data       interface{}   `json:"data"`
}
