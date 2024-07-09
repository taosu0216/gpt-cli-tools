package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

//go:embed main.py
var pythonScript embed.FS

var (
	flag        int
	tmpFile     *os.File
	contextText string
	person      int
	ai          int
)

func init() {
	person = 1
	ai = 1
	scriptBytes, err := pythonScript.ReadFile("main.py")
	if err != nil {
		log.Fatalf("Error reading embedded Python script: %v", err)
	}

	tmpFile, err = ioutil.TempFile("", "main.py")
	if err != nil {
		log.Fatalf("Error creating temporary file: %v", err)
	}

	if _, err = tmpFile.Write(scriptBytes); err != nil {
		log.Fatalf("Error writing to temporary file: %v", err)
	}
	if err = tmpFile.Close(); err != nil {
		log.Fatalf("Error closing temporary file: %v", err)
	}

}

func main() {
	flag = 0
	err := syscall.Setenv("PYTHONIOENCODING", "utf-8")
	if err != nil {
		panic(err)
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(tmpFile.Name())

	reader := bufio.NewReader(os.Stdin)
	for {
		input := ""
		if flag == 0 {
			input = first(reader)
		} else {
			input = other(reader)
		}
		contextText += fmt.Sprintf("%d. ", person) + "User: " + input + "\n"
		person++
		cmd := exec.Command("python", tmpFile.Name(), contextText)
		out, er := cmd.CombinedOutput()
		if er != nil {
			fmt.Printf("combined out err :\n%s\n", string(out))
			log.Fatalf("cmd.Run() failed with %s\n", er)
		}
		contextText += fmt.Sprintf("%d. ", ai) + "AI: " + string(out) + "\n"
		ai++
		log.Println("ai answer: ")
		color.Cyan("%s\n", string(out))
		flag++
	}
}

func first(reader *bufio.Reader) string {
	color.Blue("Ask anything...")
	return readUntilDoubleNewline(reader)
}
func other(reader *bufio.Reader) string {
	color.Blue("Your: ")
	return readUntilDoubleNewline(reader)
}

func readUntilDoubleNewline(reader *bufio.Reader) string {
	var inputLines []string
	var emptyLineCount int // 新增变量，用来计数空行的数量

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入错误:", err)
			return ""
		}

		trimmedInput := strings.TrimSpace(input) // 使用strings.TrimSpace去除首尾空白字符

		if trimmedInput == "" {
			emptyLineCount++         // 空行计数增加
			if emptyLineCount >= 2 { // 如果连续两个空行
				color.Red("loading...") // 打印提示信息
				break                   // 结束循环
			}
		} else {
			inputLines = append(inputLines, trimmedInput) // 追加非空行
			emptyLineCount = 0                            // 重置空行计数
		}
	}

	return strings.Join(inputLines, "\n")
}
