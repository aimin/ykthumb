package ykimg

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	os.Remove("testdata/resize.jpg")
	os.Remove("testdata/watermark.jpg")
	os.Remove("testdata/webp2jpg.jpg")
	os.Remove("testdata/jpg2webp.jpg")
	os.Remove("testdata/png2webp.webp")
	os.Remove("testdata/png2jpg.jpg")
	os.Remove("testdata/gif2webp.webp")
	os.Remove("2@small.jpg")
}

func Test_handler_resize(t *testing.T) {
	err := resize_img("testdata/source.jpg", 200, 200, "testdata/resize.jpg")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = os.Stat("testdata/resize.jpg")
	if err != nil {
		t.Error("failed!")
	}
	fmt.Print(".")
}

func Test_handler_watermark(t *testing.T) {
	err := watermark("testdata/source.jpg", "testdata/watermark.jpg", "testdata/w.jpg")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = os.Stat("testdata/watermark.jpg")
	if err != nil {
		t.Error("failed!")
	}
	fmt.Print(".")
}

func Test_handler_webp2jpg(t *testing.T) {
	err := webp2jpg("testdata/source.webp", "testdata/webp2jpg.jpg")
	if err != nil {
		fmt.Println("xxxx")
		fmt.Println(err.Error())
		return
	}
	_, err = os.Stat("testdata/webp2jpg.jpg")
	if err != nil {
		t.Error("failed!")
	}
	fmt.Print(".")
}

func Test_handler_jpg2webp(t *testing.T) {
	err := jpg2webp("testdata/source.jpg", "testdata/jpg2webp.webp")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = os.Stat("testdata/jpg2webp.webp")
	if err != nil {
		t.Error("failed!")
	}
	fmt.Print(".")
}

func Test_handler_png2webp(t *testing.T) {
	err := towebp("testdata/source.png", "testdata/png2webp.webp", "png")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = os.Stat("testdata/png2webp.webp")
	if err != nil {
		t.Error("failed!")
	}
	fmt.Print(".")
}

func Test_handler_gif2webp(t *testing.T) {
	err := towebp("testdata/source.gif", "testdata/gif2webp.webp", "gif")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = os.Stat("testdata/gif2webp.webp")
	if err != nil {
		t.Error("failed!")
	}
	fmt.Print(".")
}

func Test_handler_gif2jpg(t *testing.T) {
	err := tojpg("testdata/source.gif", "testdata/gif2jpg.jpg", "gif")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = os.Stat("testdata/gif2jpg.jpg")
	if err != nil {
		t.Error("failed!")
	}
	fmt.Print(".")
}
