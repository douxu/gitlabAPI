package main

import (
	"flag"
	"fmt"
	"gitlabAPI/modules"
	"gitlabAPI/util"
	"io/ioutil"
	"log"
	"strconv"

	gitlab "github.com/douxu/go-gitlab"
	"gopkg.in/yaml.v2"
)

func main() {

	var iessueT gitlab.Issue
	fmt.Println(iessueT)

	// page 当前的页码
	// per_page 当前的页面存放issue的个数

	config := new(modules.Config)
	resChan := make(chan interface{})
	issueList := make([]modules.Issue, 0)
	confPath := flag.String("C", "./conf.yaml", "config file path")
	flag.Parse()
	fmt.Println("configPath:", confPath)
	conf, err := ioutil.ReadFile(*confPath)
	if err != nil {
		log.Println("read config file failed, err is :", err.Error())
		return
	}
	err = yaml.Unmarshal(conf, config)
	if err != nil {
		log.Println("unmarshal config data failed, err is :", err.Error())
		return
	}
	fmt.Println(config)
	// targetID, err := util.GetProjectInfo(config)
	// if err != nil {
	// 	log.Println("get project ID failed, err is :", err.Error())
	// 	return
	// }
	targetID := config.ProjectInfo.ID
	fmt.Println("targetID", targetID)
	issuesPara := modules.IssuePara{
		TargetID: targetID,
		FirstGet: true,
	}
	issues, err := util.GetIssuesInfo(config, issuesPara, nil)
	if err != nil {
		log.Println("get issues by project ID failed, err is :", err.Error())
		return
	}
	firstRes := issues.(modules.IssueWithHeader)
	issueList = append(issueList, firstRes.Issues...)
	totalPage := firstRes.Header.TotalPages
	curPage := firstRes.Header.CurrentPage
	fmt.Println("totalPage", totalPage)
	fmt.Println("curPage", curPage)
	issuesPara.FirstGet = false
	for i := curPage + 1; i <= totalPage; i++ {
		issuesPara.Index = strconv.Itoa(i)
		go util.GetIssuesInfo(config, issuesPara, resChan)
	}
	fmt.Println("go routine end")
	total := totalPage - curPage
	for {
		fmt.Println("total", total)
		select {
		case value, ok := <-resChan:
			if ok {
				if value.(modules.IssueResult).OK == nil {
					total--
				}
				// process issues slice
				issueList = append(issueList,
					value.(modules.IssueResult).Data.([]modules.Issue)...)
			}
		}
		if total == 0 {
			fmt.Println("target complete")
			close(resChan)
			break
		}
	}
	fmt.Println("get data success")
	// generate Markdown file
	err = util.GenerateMarkdownFile(config, issueList)
	if err != nil {
		fmt.Println("generate markdown file failed,err is :" + err.Error())
		return
	}
}
