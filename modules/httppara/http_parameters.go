package httppara

import "gitlabAPI/modules/token"

// HTTPParameters struct of http Parameter
type HTTPParameters struct {
	URL       string              // http request's address
	Types     string              // http request's type
	Parameter token.PersonalToken // http request's header parameter
}

// InitPara func of init http request para
func (h *HTTPParameters) InitPara(url, types string, para token.PersonalToken) {
	h.URL = url
	h.Types = types
	h.Parameter = para
}
