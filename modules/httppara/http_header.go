package httppara

import (
	"net/http"
	"strconv"
)

// HTTPHeader struct of http header value
type HTTPHeader struct {
	TotalPages  int // total pages
	TotalItems  int // total items
	CurrentPage int // current index of page
}

// Init func of init value of http request header
func (h *HTTPHeader) Init(resp *http.Response) (err error) {
	// http header's total pages
	h.TotalPages, err = strconv.Atoi(resp.Header["X-Total-Pages"][0])
	if err != nil {
		return err
	}
	// http header's total items
	h.TotalItems, err = strconv.Atoi(resp.Header["X-Total"][0])
	if err != nil {
		return err
	}
	// http header's current index of page
	h.CurrentPage, err = strconv.Atoi(resp.Header["X-Page"][0])
	if err != nil {
		return err
	}
	return err
}
