package dtos

type ResponseDTO struct {
	Pagination PaginationDTO `json:"pagination"`
	Data       interface{}   `json:"data"`
}
