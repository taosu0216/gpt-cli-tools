# gpt-cli-tools
A tool that allows you to chat with GPT in the cli

# use
## 依赖
```shell
pip install -r requirements.txt
```
## 配置
在main.py中填充自己的api_key和baseurl,以及model(对话模型)
## 使用
### 生成可执行文件
```shell
git clone https://github.com/taosu0216/gpt-cli-tools.git
cd gpt-cli-tools
go mod init gpt-cli-tools
go mod tidy
go build -o ai.exe main.go
```
### 添加到环境变量
只需将该exe文件所在目录添加到环境变量即可
设置 -> 系统 -> 系统信息 -> 高级系统设置 -> 环境变量 -> 用户环境变量 -> Path -> 将exe文件所在目录添加到环境变量中
### 效果图
![img](https://img.picui.cn/free/2024/07/09/668c9245b4f3a.png)
然后就可以在电脑的任一位置使用跟ai对话的cli了
进行提问的判断是界面有两个空行,也就是三行回车

# v2
直接在main.go修改baseurl和key,然后go build即可