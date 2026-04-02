package handler

import (
    "os"
    "fmt"
    "bytes"
    "image"
    "strconv"
    "image/png"
    "image/jpeg"
    "image/color"
    "github.com/nfnt/resize"
    imagerecog20190930  "github.com/alibabacloud-go/imagerecog-20190930/v2/client"
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

/*颜色推荐*/
func colorRecommend(colorList []ColorOut, data *imagerecog20190930.RecognizeImageColorResponseBodyData) ([]ColorOut, error) {
    var colorSlice []color.Color
    colorMap := make(map[string]ColorOut)
    for _, item := range colorList {
        colorSlice = append(colorSlice, color.RGBA{item.R, item.G, item.B, 255})
        colorMap[fmt.Sprintf("%d%d%d", item.R, item.G, item.B)] = item
    }
    var err error
    var flag bool
    var r, g, b uint64
    var rr, gg, bb uint32
    var rrr, ggg, bbb uint8
    var nearestColor color.Color
    var palette color.Palette = colorSlice
    colors := []ColorOut{}
    for _, item := range data.ColorTemplateList {
        if *item.Percentage > 0.15 && len(*item.Color) == 6 {
            r, err = strconv.ParseUint((*item.Color)[0:2], 16, 8)
            if err != nil {
                return nil, err
            }
            g, err = strconv.ParseUint((*item.Color)[2:4], 16, 8)
            if err != nil {
                return nil, err
            }
            b, err = strconv.ParseUint((*item.Color)[4:6], 16, 8)
            if err != nil {
                return nil, err
            }
            nearestColor = palette.Convert(color.RGBA{uint8(r), uint8(g), uint8(b), 0xff}) 
            rr, gg, bb, _ = nearestColor.RGBA()
            rrr = uint8(rr)
            ggg = uint8(gg)
            bbb = uint8(bb)
            flag = true
            for _, recommend := range colors {
                if rrr == recommend.R && ggg == recommend.G && bbb == recommend.B {
                    flag = false
                    break
                }
            }
            if flag {
                colors = append(colors, colorMap[fmt.Sprintf("%d%d%d", rrr, ggg, bbb)])
            }
        }
    }
    return colors, nil
}
