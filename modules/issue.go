package modules

import "gitlabAPI/modules/httppara"

// Issue the struct of gitlab's issue
type Issue struct {
	ProjectID int       `json:"project_id"` // issue所在project的ID
	Title     string    `json:"title"`      // issue的标题
	WebURL    string    `json:"web_url"`    // issue的网页
	State     string    `json:"state"`      // issue的状态
	Tag       Milestone `json:"milestone"`  // issue的tag
	Labels    []string  `json:"labels"`     // issue的labels
}

// IssueWithHeader struct of first issue request
type IssueWithHeader struct {
	Header httppara.HTTPHeader
	Issues []Issue
}

// IssuePara the struct of issus's request
type IssuePara struct {
	TargetID int    // issue所在project的ID
	FirstGet bool   // 是否为首次请求
	Index    string // issue所在的页码
}

// IssueResult the struct of issus's process result
type IssueResult struct {
	Data interface{} // 处理后的issue数据
	OK   error       // 错误
}
