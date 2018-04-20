package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	var format = flag.String("f", "png", "encode format")
	flag.Parse()
	img, err := readImage(os.Stdin)
	if err != nil {
		log.Fatalln("image read error:%v", err)
	}
	switch *format {
	case "png":
		err = toPNG(img, os.Stdout)
	case "jpeg", "jpg":
		err = toJPEG(img, os.Stdout)
	case "gif":
		err = toGIF(img, os.Stdout)
	default:
	}
	if err != nil {
		log.Fatalln("convert error:%v", err)
	}
}
func readImage(in io.Reader) (image.Image, error) {
	img, kind, err := image.Decode(in)
	if err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr, "Imput format =", kind)
	return img, nil
}
func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}
func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toGIF(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil})
}
