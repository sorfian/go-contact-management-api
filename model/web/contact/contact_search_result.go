package contact

import "github.com/sorfian/go-contact-management-api/model/web"

type SearchResult struct {
	Contacts []ContactResponse  `json:"contacts"`
	Paging   web.PagingResponse `json:"paging"`
}
