# NetSpider
基于Go 1.19 的 站点模板爬虫

## 拉取并构建 NetSpider
```shell
git clone https://github.com/Kongxiaoaaa/NetSpider.git
cd NetSpider
go mod init NetSpider
go mod tidy
go run main.go
```

## 更改设置
`main.go`
```go
func main() {
  // 启动爬虫
  downUrl := util.GetInput()
  app.Fetch(downUrl, "./res")  // Fetch(url, downloadDir)
}
```
