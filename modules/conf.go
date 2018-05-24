package modules

// IssueConf struct of conf file issue info
type IssueConf struct {
	CreatedBySelf bool   `yaml:"createdbyself"` // 表明查询的issue是否由自己创建,false为查询所有可见的issue
	State         string `yaml:"state"`         // 表明issue的状态是opened/closed
	Tag           string `yaml:"tag"`           // 表明issue所属的tag
}

// ProjectConf struct of conf file's project info
type ProjectConf struct {
	ID   int    `yaml:"id"`   // issue所属项目的ID
	Name string `yaml:"name"` // issue所属项目的Name
}

// AccessConf struct of conf file's project info
type AccessConf struct {
	WebURL        string `yaml:"webURL"`        //gitlab地址
	PersonalToken string `yaml:"personalToken"` // 个人的Token
}

// FileConf struct of conf file's makrdown file info
type FileConf struct {
	Title    string   `yaml:"title"`    // Markdown文件的一级标题
	Subtitle []string `yaml:"subtitle"` // Markdown文件的四级标题
}

// Config struct of config file
type Config struct {
	AccessInfo  AccessConf  `yaml:"accessInfo"`  // config文件中的access信息
	ProjectInfo ProjectConf `yaml:"projectInfo"` // config文件中project信息
	IssueInfo   IssueConf   `yaml:"issueInfo"`   // config文件中issue信息
	FileInfo    FileConf    `yaml:"fileInfo"`    // conf文件中markdown文件信息
}
