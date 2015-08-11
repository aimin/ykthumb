package ykimg

import (
	"errors"
)

var errs map[int]string

func init() {
	errs = map[int]string{
		10000: "源图片无效!",
		10001: "源图片格式有误!",
		10401: "水印源图无效!",
		10402: "水印源图有误!",
	}
}

func Newcuserr(code int) error {
	return errors.New(errs[code])
}
