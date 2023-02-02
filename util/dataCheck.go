package util

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
)

// 校验 URL 地址
func checkStoreAddr(addr string) error {
	if len(addr) == 0 {
		return errors.New("url is empty")
	}
	parse, err := url.Parse(addr)
	if err != nil {
		return err
	}

	if parse.Scheme != "http" && parse.Scheme != "https" {
		return errors.New("url Scheme illegal. err:" + parse.Scheme)
	}
	re := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)+([A-Za-z]|[A-Za-z][A-Za-z0-9\-]*[A-Za-z0-9])$`)

	u := strings.Split(parse.Host, ":")
	if len(u) == 2 && len(u[1]) != 0 {
		result := re.FindAllStringSubmatch(u[0], -1)
		if result == nil {
			return errors.New("url illegal. err:" + u[0])
		}
		return nil
	}
	result := re.FindAllStringSubmatch(parse.Host, -1)
	if result == nil {
		return errors.New("url illegal. err:" + parse.Host)
	}
	return nil
}

func GetInput() (downUrl string) {
	fmt.Printf("Download Page url: ")
	_, err := fmt.Scanln(&downUrl)
	if err != nil {
		log.Fatal("暂未输入 URL")
		return
	}
	err = checkStoreAddr(downUrl)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	return
}
