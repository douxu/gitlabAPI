package httppara

// HTTPResponse struct of htttp response
// info and header info
type HTTPResponse struct {
	HTTPHeader   HTTPHeader // http header data
	HTTPResponse []byte     // http request result
}
