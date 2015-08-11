package ykimg

import (
	"bytes"
	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func loadImg(file string) (image.Image, error) {
	return loadImgSpecFormat(file, "jpg")
}

func loadImgSpecFormat(file string, ft string) (image.Image, error) {
	file_origin, err := os.Open(file)
	if err != nil {
		return nil, Newcuserr(10000)
	}
	defer file_origin.Close()
	var img0 image.Image
	switch ft {
	case "gif":
		img0, err = gif.Decode(file_origin)
	case "png":
		img0, err = png.Decode(file_origin)
	case "webp":
		img0, err = webp.Decode(file_origin)
	default:
		img0, err = jpeg.Decode(file_origin)
	}
	if err != nil {
		return nil, Newcuserr(10001)
	}
	return img0, nil
}

func jpg2webp(file string, to string) error {
	return towebp(file, to, "jpg")
}

//to webp
func towebp(file string, to string, sft string) error {
	img0, err := loadImgSpecFormat(file, sft)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = webp.Encode(buf, img0, &webp.Options{Quality: 75})
	if err = ioutil.WriteFile(to, buf.Bytes(), 0666); err != nil {
		return err
	}
	return nil
}

//to jpg
func webp2jpg(file string, to string) error {
	return tojpg(file, to, "webp")
}

func tojpg(file string, to string, sft string) error {
	var buf bytes.Buffer

	img0, err := loadImgSpecFormat(file, sft)

	// Encode lossless webp
	err = jpeg.Encode(&buf, img0, &jpeg.Options{100})
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(to, buf.Bytes(), 0666); err != nil {
		return err
	}
	return nil
}

//改变大小
func resize_img(file string, width uint, height uint, to string) error {
	return resize_img_qua(file, width, height, to, int(80))
}

// 改变大小及清晰度
func resize_img_qua(file string, width uint, height uint, to string, qua int) error {
	origin, err := loadImg(file)

	canvas := resize.Resize(width, height, origin, resize.Lanczos3)

	file_out, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer file_out.Close()

	jpeg.Encode(file_out, canvas, &jpeg.Options{qua})

	return nil

}

func cmd_thumbnail(file string, width uint, height uint, to string) {
	// 打开图片并解码
	file_origin, _ := os.Open(file)
	origin, _ := jpeg.Decode(file_origin)
	defer file_origin.Close()

	canvas := resize.Thumbnail(width, height, origin, resize.Lanczos3)
	file_out, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer file_out.Close()

	jpeg.Encode(file_out, canvas, &jpeg.Options{80})
}

// 水印
func watermark(file string, to string, w string) error {
	// 打开图片并解码
	origin, err := loadImg(file)

	// 打开水印图并解码
	watermark, err := loadImgSpecFormat(w, "png")
	if err != nil {
		return err
	}

	//原始图界限
	origin_size := origin.Bounds()

	//创建新图层
	canvas := image.NewNRGBA(origin_size)
	// 贴原始图
	draw.Draw(canvas, origin_size, origin, image.ZP, draw.Src)
	// 贴水印图
	draw.Draw(canvas, watermark.Bounds().Add(image.Pt(30, 30)), watermark, image.ZP, draw.Over)

	//生成新图片
	create_image, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	jpeg.Encode(create_image, canvas, &jpeg.Options{95})
	defer create_image.Close()
	return nil
}
