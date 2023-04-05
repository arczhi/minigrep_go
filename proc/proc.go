package proc

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"unsafe"
)

// 运行逻辑
func Run(config *Config) error {
	//读取文件
	f, err := os.Open(config.FilePath)
	if err != nil {
		return err
	}
	//退出前关闭文件句柄
	defer f.Close()
	buf, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	//将文件字节转为字符串
	contents := *(*string)(unsafe.Pointer(&buf))
	//进行关键词检索
	result, err := search(config.Query, &contents)
	if err != nil {
		return err
	}
	//打印结果
	printResult(result)
	return nil
}

// 构建配置
type Config struct {
	Query    string
	FilePath string
}

func (c *Config) Build(args []string) error {
	if len(args) < 3 {
		return errors.New("no enough arguments")
	}
	c.Query = args[1]
	c.FilePath = args[2]
	return nil
}

// 检索文件
type SearchResult struct {
	Word    string
	LineNum int
}

// query要检索的关键词 contents文本字符串
func search(query string, contents *string) (*[]SearchResult, error) {
	//结果集
	result := &[]SearchResult{}
	//换行符
	lineBreak, err := getLineBreak()
	if err != nil {
		return nil, err
	}

	//匹配字符串中的每一行
	for line_num, line := range strings.Split((*contents), lineBreak) {
		if strings.Contains(strings.ToUpper(line), query) ||
			strings.Contains(strings.ToLower(line), query) {
			*result = append(*result, SearchResult{
				Word:    line,
				LineNum: line_num + 1, //行数为下标+1
			})
		}
	}

	return result, nil
}

// 根据系统来决定换行符
func getLineBreak() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return "\r\n", nil
	case "linux":
		return "\n", nil
	case "darwin":
		return "\r", nil
	default:
		return "", errors.New("unknown arch !")
	}
}

// 打印输出
func printResult(result *[]SearchResult) {
	for _, r := range *result {
		fmt.Printf("line %d : %s\n", r.LineNum, r.Word)
	}
}
