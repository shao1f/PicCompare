package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/disintegration/imageorient"
	_ "golang.org/x/image/webp"
)

func GetImageFromUrl(url string) (image.Image, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	img, _, err := imageorient.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func main() {
	i1, err := GetImageFromUrl("xxx.png") // 替换为网络图片
	if err != nil {
		log.Fatal(err)
	}
	i2, err := GetImageFromUrl("xxx1.png") // 替换为网络图片
	if err != nil {
		log.Fatal(err)
	}
	width := i1.Bounds().Dx()
	height := i1.Bounds().Dy()
	var t, r int
	for y := 20; y < height-20; y++ {
		for x := 20; x < width-20; x++ {
			r1, g1, b1, _ := i1.At(x, y).(color.NRGBA).R, i1.At(x, y).(color.NRGBA).G, i1.At(x, y).(color.NRGBA).B, i1.At(x, y).(color.NRGBA).A
			r2, g2, b2, _ := i2.At(x, y).(color.NRGBA).R, i2.At(x, y).(color.NRGBA).G, i2.At(x, y).(color.NRGBA).B, i2.At(x, y).(color.NRGBA).A

			rgb1 := RGB{
				red:   int64(r1),
				green: int64(g1),
				blue:  int64(b1),
			}
			h1 := rgb1.rgb2hex()
			ix1, err := strconv.ParseInt(h1.str, 16, 64)
			if err != nil {
				log.Fatal(err)
			}

			rgb2 := RGB{
				red:   int64(r2),
				green: int64(g2),
				blue:  int64(b2),
			}
			h2 := rgb2.rgb2hex()
			ix2, err := strconv.ParseInt(h2.str, 16, 64)
			if err != nil {
				log.Fatal(err)
			}

			if math.Abs(float64(ix1-ix2)) > 1800000 {
				t++
				r += x
			}
		}
	}
	fmt.Println(int(math.Round(float64(r/t))) - 55)
}

type RGB struct {
	red, green, blue int64
}

type HEX struct {
	str string
}

func t2x(t int64) string {
	result := strconv.FormatInt(t, 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}
func (color RGB) rgb2hex() HEX {
	r := t2x(color.red)
	g := t2x(color.green)
	b := t2x(color.blue)
	return HEX{r + g + b}
}
