package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

var (
	errImageSizeMismatch = errors.New("image size mismatch")
)

func OpenImage(path string) (image.Image, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	img, err := png.Decode(file)

	if err != nil {
		return nil, err
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return img, nil
}

func writeImg(img image.Image, path string) error {
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	err = png.Encode(file, img)

	if err != nil {
		return err
	}

	return nil
}

func encode(srcPath, auxPath, outPath string) error {
	src, err := OpenImage(srcPath)

	if err != nil {
		return err
	}

	aux, err := OpenImage(auxPath)

	if err != nil {
		return err
	}

	sx, sy := src.Bounds().Dx(), src.Bounds().Dy()
	ax, ay := aux.Bounds().Dy(), aux.Bounds().Dy()

	if sx != ax || sy != ay {
		return errImageSizeMismatch
	}

	out := image.NewRGBA(image.Rect(0, 0, sx, sy))

	for y := 0; y < sy; y++ {
		for x := 0; x < sx; x++ {
			sr, sg, sb, sa := src.At(x, y).RGBA()
			ar, ag, ab, aa := aux.At(x, y).RGBA()

			sr, sg, sb, sa = sr/256, sg/256, sb/256, sa/256
			ar, ag, ab, aa = ar/256, ag/256, ab/256, aa/256

			nr := uint8((sr & 0xfc) | (ar / 64))
			ng := uint8((sg & 0xfc) | (ag / 64))
			nb := uint8((sb & 0xfc) | (ab / 64))
			na := uint8((sa & 0xfc) | (aa / 64))

			c := color.RGBA{
				R: nr,
				G: ng,
				B: nb,
				A: na,
			}

			out.Set(x, y, c)
		}
	}

	err = writeImg(out, outPath)

	if err != nil {
		return err
	}

	return nil
}

func decode(srcPath, outPath string) error {
	src, err := OpenImage(srcPath)

	if err != nil {
		return err
	}

	sx, sy := src.Bounds().Dx(), src.Bounds().Dy()

	out := image.NewRGBA(image.Rect(0, 0, sx, sy))

	for y := 0; y < sy; y++ {
		for x := 0; x < sx; x++ {
			sr, sg, sb, sa := src.At(x, y).RGBA()

			sr, sg, sb, sa = sr/256, sg/256, sb/256, sa/256

			nr := uint8((sr & 3) * 64)
			ng := uint8((sg & 3) * 64)
			nb := uint8((sb & 3) * 64)
			na := uint8((sa & 3) * 64)

			c := color.RGBA{
				R: nr,
				G: ng,
				B: nb,
				A: na,
			}

			out.Set(x, y, c)
		}
	}

	err = writeImg(out, outPath)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	var (
		srcPath string
		auxPath string
		outPath string
		mode    string
	)

	flag.StringVar(&srcPath, "s", "", "The path of the source or encoded image.")
	flag.StringVar(&auxPath, "a", "", "The path of the image to encode into the source.")
	flag.StringVar(&outPath, "o", "", "The path of the output image.")
	flag.StringVar(&mode, "m", "", "The mode to use. Options are `encode` or `decode`.")

	flag.Parse()

	if mode != "encode" && mode != "decode" {
		if mode == "" {
			fmt.Println("No mode was specified. Please use -m `encode` or `decode`.")
		} else {
			fmt.Println("Invalid mode. Please use -m `encode` or `decode`.")
		}
		return
	}

	if srcPath == "" {
		fmt.Println("Source path not specified. Please specify a source path with -s")
		return
	}

	if outPath == "" {
		fmt.Println("Output path not specified Please specify a source path with -o")
		return
	}

	if auxPath == "" && mode == "encode" {
		fmt.Println("An auxiliary path must be specified for this operation. Please specify a path with -a")
		return
	}

	var err error

	if mode == "encode" {
		err = encode(srcPath, auxPath, outPath)
	} else if mode == "decode" {
		err = decode(srcPath, outPath)
	}

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}
