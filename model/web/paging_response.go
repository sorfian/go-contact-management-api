package web

type PagingResponse struct {
	Page      int `json:"page"`
	Size      int `json:"size"`
	TotalPage int `json:"total_page"`
	TotalItem int `json:"total_item"`
}
