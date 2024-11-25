package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func GetPostsLocation() string {
	// 获取当前文件的绝对路径
	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)
	currentDir = filepath.Dir(currentDir)
	currentDir = filepath.Dir(currentDir)

	return fmt.Sprintf("%s/_posts/", currentDir)
}

func GetRootLocation() string {
	// 获取当前文件的绝对路径
	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)
	currentDir = filepath.Dir(currentDir)
	currentDir = filepath.Dir(currentDir)

	return fmt.Sprintf(currentDir + "/")
}

func GetBlackListFileLocation() string {
	return GetRootLocation() + "blacklist/black.txt"
}

func GetArticleLocation(file string) string {
	// 获取当前文件的绝对路径
	posts := GetPostsLocation()
	currentDir := filepath.Dir(posts)
	currentDir = filepath.Dir(currentDir)
	log.Println(posts, currentDir, file)
	relativePath, err := filepath.Rel(currentDir, file)
	if err != nil {
		return ""
	}
	return relativePath
}

func GetArticleContent(url string) (string, error) {
	content, err := ioutil.ReadFile(GetRootLocation() + url)
	log.Println("url is:", url, "root location is: ", GetRootLocation(), "post location is: ", GetPostsLocation())
	if err != nil {
		log.Println("读取文件失败：", err)
		return "", InternalErr
	}
	return string(content), nil
}

func UpdateBlackListFile(ip string) {
	log.Printf("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@  %s is added into blacklist", ip)
	// 获取当前文件的绝对路径
	filePath := GetBlackListFileLocation()
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// 文件不存在，创建文件并写入内容
		file, err := os.Create(filePath)
		if err != nil {
		}
		defer file.Close()

		_, _ = file.WriteString(ip)
	} else {
		// 文件存在，打开并追加写入内容
		file, _ := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		defer file.Close()
		ip = "\n" + ip
		_, _ = file.WriteString(ip)
	}
}

func GetBlackList() []string {
	// 获取当前文件的绝对路径
	filePath := GetBlackListFileLocation()
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil
	}
	lines, err := readLines(filePath)
	if err != nil {
		log.Fatalf("failed to read lines: %s", err)
	}
	return lines
}
func readLines(filePath string) ([]string, error) {
	// 使用 os.ReadFile 来读取整个文件内容到内存
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// 将文件内容按换行符分割成多个行
	lines := strings.Split(string(data), "\n")

	// 过滤空行和仅包含空白字符的行
	var filteredLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			filteredLines = append(filteredLines, trimmed)
		}
	}
	return filteredLines, nil
}

func AnyToJsonStr(v interface{}) (string, error) {
	// 将传入的变量序列化为 JSON 字符串
	jsonData, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	// 将字节数组转化为字符串返回
	return string(jsonData), nil
}

func JsonStrToAny(jsonStr string, v interface{}) error {
	// 将 JSON 字符串解析到传入的变量 v 中
	return json.Unmarshal([]byte(jsonStr), v)
}
func JsonStrSliceToAny(jsonStrs []string, v interface{}) error {

	// 确保传递的 v 是一个指向切片的指针
	slicePtr := reflect.ValueOf(v)
	if slicePtr.Kind() != reflect.Ptr || slicePtr.Elem().Kind() != reflect.Slice {
		return errors.New("v must be a pointer to a slice")
	}

	// 获取切片的值，并确定其元素类型
	sliceValue := slicePtr.Elem()
	structType := sliceValue.Type().Elem()

	// 遍历 JSON 字符串切片
	for _, jsonStr := range jsonStrs {
		// 创建新的元素，并确保是指针类型以便于解码
		elemPtr := reflect.New(structType).Interface()

		// 进行 JSON 解码
		err := json.Unmarshal([]byte(jsonStr), elemPtr)
		if err != nil {
			return err
		}

		// 将解码结果追加到切片中
		sliceValue.Set(reflect.Append(sliceValue, reflect.ValueOf(elemPtr).Elem()))
	}
	return nil
}

func NowTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
