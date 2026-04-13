package handler

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

/*解析图片数据流*/
func parseImageData(imgData *bytes.Buffer) (image.Image, error) {
    var err error
    var img image.Image
    img, _, err = image.Decode(imgData)
    return img, err
}

/*保存图片*/
func saveJpgImage(img image.Image, path string) error {
    /*保存到的文件*/
    outFile, err := os.Create(path)
    if err != nil {
        return err
    }
    defer outFile.Close()
    /*开始保存*/
    return jpeg.Encode(outFile, img, &jpeg.Options{Quality:100})      
}

/*保存图片*/
func savePngImage(img image.Image, path string) error {
    /*保存到的文件*/
    outFile, err := os.Create(path)
    if err != nil {
        return err
    }
    defer outFile.Close()
    /*开始保存*/
    var b bytes.Buffer
    err = png.Encode(&b, img)
    if err == nil {
        _, err = outFile.Write(b.Bytes())
    }
    return err
}

/*图片缩放*/
func resizeImg(img image.Image, width uint, height uint) (image.Image, error) {
    return resize.Resize(width, height, img, resize.Lanczos3), nil
}
