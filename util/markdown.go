package util

import (
	"bufio"
	"fmt"
	"gitlabAPI/modules"
	"os"
	"strconv"
	"strings"
)

// GenerateMarkdownFile func of generate markdown file by issues
func GenerateMarkdownFile(conf *modules.Config, issues []modules.Issue) error {
	// open a markdown file
	var file *os.File
	var err error
	filePath := "./" + conf.FileInfo.Title + ".md"
	defer file.Close()
	fmt.Println("fileName:", filePath)
	if checkFileIsExist(filePath) {
		file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
		fmt.Println("文件存在")
	} else {
		file, err = os.Create(filePath)
		fmt.Println("文件不存在,创建名为" + conf.FileInfo.Title + "的文件")
	}
	bug, new := groupingAndWriteIssue(issues)
	writer := bufio.NewWriter(file)
	_, err = fmt.Fprintln(writer, "# "+conf.FileInfo.Title)
	if err != nil {
		return err
	}
	// bug fix tile
	_, err = fmt.Fprintln(writer, "* "+conf.FileInfo.Subtitle[0])
	if err != nil {
		return err
	}
	// write bug fix's issue slice
	err = wtriteIssueSlice(writer, bug)
	if err != nil {
		return err
	}
	// what's new tilte
	_, err = fmt.Fprintln(writer, "* "+conf.FileInfo.Subtitle[1])
	if err != nil {
		return err
	}
	// write what's new issue slice
	err = wtriteIssueSlice(writer, new)
	if err != nil {
		return err
	}
	return writer.Flush()
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func groupingAndWriteIssue(issues []modules.Issue) (bug, new []modules.Issue) {
	for _, issue := range issues {
		var flag bool
		for _, label := range issue.Labels {
			if strings.Contains(strings.ToLower(label), "bug") {
				flag = true
			}
		}
		if flag == true {
			bug = append(bug, issue)
		} else {
			new = append(new, issue)
		}
		// reset flag singal
		flag = false
	}
	return bug, new
}

func wtriteIssueSlice(w *bufio.Writer, issues []modules.Issue) error {
	for index, issue := range issues {
		number := index + 1
		indexStr := strconv.Itoa(number)
		// writer issue's web URL
		fmt.Fprintln(w, indexStr+". "+"["+issue.Title+"]"+"("+issue.WebURL+")")
	}
	return nil
}
