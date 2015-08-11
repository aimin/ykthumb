package ykimg

import "os"
import "net/http"
import "io"

// 将图片下载并保存到本地
func save_img(url string, to string) error {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return err
	}

	// 创建文件
	dst, err := os.Create(to)
	if err != nil {
		return err
	}
	// 生成文件
	io.Copy(dst, res.Body)
	return nil
}

func save_img_channel(url string, to string, c chan string) error {
	err := save_img(url, to)
	if err != nil {
		return err
	}
	c <- "sucessed"

	return nil
}
