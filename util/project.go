package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlabAPI/modules"
	"gitlabAPI/modules/httppara"
	"gitlabAPI/modules/token"
)

// GetProjectInfo func of get all projects
// by Personal Access Tokens
func GetProjectInfo(conf *modules.Config) (int, error) {
	var URL string
	requestPara := new(httppara.HTTPParameters)
	projects := make([]modules.Project, 0)
	token := token.PersonalToken{Name: "PRIVATE-TOKEN", Value: conf.AccessInfo.PersonalToken}
	baseURL := conf.AccessInfo.WebURL
	URL = baseURL + "/api/v4/projects"
	fmt.Println("URL", URL)
	requestPara.InitPara(URL, "GET", token)
	projectInfo, err := HTTPRequestWithHeader(requestPara, "project")
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(projectInfo.([]byte), &projects)
	if err != nil {
		return 0, err
	}
	return processProjectInfo(projects, conf.ProjectInfo.Name)
}

func processProjectInfo(list []modules.Project, targetName string) (ProjectID int, err error) {
	for _, value := range list {
		if value.Name == targetName {
			return value.ID, err
		}
	}
	return 0, errors.New("No match Project Info")
}
