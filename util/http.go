package util

import (
	"gitlabAPI/modules/httppara"
	"io/ioutil"
	"net/http"
)

// HTTPRequestWithHeader func of send http request
func HTTPRequestWithHeader(para *httppara.HTTPParameters, from string) (interface{}, error) {
	client := &http.Client{}
	request, err := http.NewRequest(para.Types, para.URL, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(para.Parameter.Name, para.Parameter.Value)
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if from == "issue" {
		// analyze response result's header for
		// get issus's total number and current page index
		httpHeader := new(httppara.HTTPHeader)
		err := httpHeader.Init(resp)
		if err != nil {
			return nil, err
		}
		// 组装值
		issuesData := httppara.HTTPResponse{
			HTTPHeader:   *httpHeader,
			HTTPResponse: body,
		}
		return issuesData, nil
	}
	return body, nil
}
