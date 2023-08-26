package main

import (
	"bytes"
	"encoding/base64"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"github.com/nfnt/resize"
	"os"
)

func convertToBase64(inputPath string, maxSize int) (string, error) {
	// 读取图片文件
	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 解码图片文件
	imageData, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	// 压缩图片
	decodedImage := Resize(imageData, maxSize)

	// 创建一个缓冲区
	buffer := new(bytes.Buffer)

	// 编码并写入缓冲区
	err = jpeg.Encode(buffer, decodedImage, nil)
	if err != nil {
		return "", err
	}

	// 将缓冲区的内容转换为base64字符串
	encodedString := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return encodedString, nil
}

func Resize(img image.Image, maxSize int) image.Image {
	size := img.Bounds().Size()
	width := size.X
	height := size.Y

	// 如果图片尺寸小于最大大小，则直接返回原图
	if width <= maxSize && height <= maxSize {
		return img
	}

	// 计算缩放比例
	var newWidth, newHeight uint

	if width > height {
		scale := float64(maxSize) / float64(width)
		newWidth = uint(maxSize)
		newHeight = uint(float64(height) * scale)
	} else {
		scale := float64(maxSize) / float64(height)
		newWidth = uint(float64(width) * scale)
		newHeight = uint(maxSize)
	}

	// 执行缩放操作
	resized := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	return resized
}
