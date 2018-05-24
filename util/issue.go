package util

import (
	"encoding/json"
	"fmt"
	"gitlabAPI/modules"
	"gitlabAPI/modules/httppara"
	"gitlabAPI/modules/token"
	"strconv"
)

// GetIssuesInfo func of get all issues
// by Personal Access Tokens
func GetIssuesInfo(conf *modules.Config, para modules.IssuePara, resChan chan interface{}) (interface{}, error) {
	requestPara := new(httppara.HTTPParameters)
	issues := make([]modules.Issue, 0)
	URL := initIssueQuery(conf, para)
	token := token.PersonalToken{Name: "PRIVATE-TOKEN", Value: conf.AccessInfo.PersonalToken}
	fmt.Println("URL", URL)
	requestPara.InitPara(URL, "GET", token)
	issuesInfo, err := HTTPRequestWithHeader(requestPara, "issue")
	if err != nil {
		return nil, err
	}
	if para.FirstGet == true {
		result := issuesInfo.(httppara.HTTPResponse)
		err = json.Unmarshal(result.HTTPResponse, &issues)
		if err != nil {
			return nil, err
		}
		issueList, err := processIssueInfo(issues, conf, para.TargetID)
		if err != nil {
			return nil, err
		}
		data := modules.IssueWithHeader{
			Header: result.HTTPHeader,
			Issues: issueList,
		}
		return data, nil
	}
	err = json.Unmarshal(issuesInfo.(httppara.HTTPResponse).HTTPResponse, &issues)
	if err != nil {
		return nil, err
	}
	issueList, err := processIssueInfo(issues, conf, para.TargetID)
	fmt.Println(issueList)
	resChan <- modules.IssueResult{Data: issueList, OK: err}
	return nil, err
}

func processIssueInfo(list []modules.Issue, conf *modules.Config, targetID int) (issues []modules.Issue, err error) {
	issues = make([]modules.Issue, 0)
	for _, value := range list {
		if value.ProjectID == targetID && value.Tag.Title == conf.IssueInfo.Tag {
			issues = append(issues, value)
		}
	}
	return issues, nil
}

func initIssueQuery(conf *modules.Config, para modules.IssuePara) (url string) {
	var issuesQuery, requestPage string
	projectID := strconv.Itoa(para.TargetID)
	scope := conf.IssueInfo.CreatedBySelf
	if !scope {
		issuesQuery = "scope=all&"
	}
	status := conf.IssueInfo.State
	if status == "" {
		// if not state the status of the issue
		// default inquire closed issue
		issuesQuery += "state=closed"
	} else {
		issuesQuery += "state=" + status
	}
	// set every page display limit is 100
	issuesQuery += "&per_page=100"
	if para.Index == "" {
		requestPage = "&page=" + "1"
	} else {
		requestPage = "&page=" + para.Index
	}
	// set request page's index
	issuesQuery += requestPage
	return conf.AccessInfo.WebURL + "/api/v4/projects/" + projectID + "/issues?" + issuesQuery
}
