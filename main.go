package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/sunshineplan/imgconv"
)

func setWhiteBackground(imgSrc image.Image) image.Image {
	backgroundImg := image.NewRGBA(imgSrc.Bounds())

	draw.Draw(backgroundImg, backgroundImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(backgroundImg, backgroundImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)

	return backgroundImg
}

func getBase64Image(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Failed to open file in \"%v\": %v", filepath, err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatalf("Failed to get file size: %v", err)
	}

	bs := make([]byte, stat.Size())

	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		log.Fatalf("Failed to read bite slice: %v", err)
	}

	mimeType := http.DetectContentType(bs)

	base64Encoded := "data:"+mimeType+";base64,"+base64.StdEncoding.EncodeToString(bs)
	return base64Encoded
}

func main() {
	args := os.Args

	mapFormat := []string{"jpg","png","gif","tiff","bmp"}

	var output string	
	format := imgconv.JPEG
	resizeOpts := imgconv.ResizeOption{}

	for argsIndex := 2; argsIndex < len(args); argsIndex+=2 {
		switch flag := args[argsIndex]; flag {
		case "-o":
			output = args[argsIndex+1]
		case "-f":
			switch form := args[argsIndex+1]; form {
			case "JPEG":
				format = imgconv.JPEG
			case "PNG":
				format = imgconv.PNG
			case "BMP":
				format = imgconv.BMP
			case "GIF":
				format = imgconv.GIF
			case "TIFF":
				format = imgconv.TIFF
			default:
				log.Fatal("Output format invalid. The supported values are JPEG, PNG, BMP, GIT and TIFF.")
			}
		case "-w":
			width, err := strconv.Atoi(args[argsIndex+1])
			if err != nil {
				log.Fatalf("Invalid value for width parameter: %v", err)
			}

			resizeOpts.Width = width
		case "-h":
			height, err := strconv.Atoi(args[argsIndex+1])
			if err != nil {
				log.Fatalf("Invalid value for height value: %v", err)
			}

			resizeOpts.Height = height
		case "--base64":
			fmt.Println(getBase64Image(args[1]))
			return
		default:
			log.Fatal("Invalid arguments.\n"+
				"\tTry: convertimg <file> [-o <output-filename>] [-f JPEG|PNG|BMP|GIF|TIFF] [-w <int>] [-h <int>]\n"+
				"\t     convertimg <file> --base64")
		}
	}

	imgSrc, err := imgconv.Open(args[1])
	if err != nil {
		log.Fatalf("Failed to open image: %v", err)
	}

	if format == imgconv.JPEG | imgconv.BMP {
		imgSrc = setWhiteBackground(imgSrc)
	}

	if (resizeOpts != imgconv.ResizeOption{}) {
		imgSrc = imgconv.Resize(imgSrc, &resizeOpts)
	}

	err = imgconv.Write(io.Discard, imgSrc, &imgconv.FormatOption{Format: format})
	if err != nil {
		log.Fatalf("Failed to write image: %v", err)
	}

	if output == "" {
		arrString := strings.Split(args[1],".")
		file := strings.Join(arrString[:len(arrString)-1],".")
		output = fmt.Sprintf("%v.%v",file,mapFormat[format])
	}

	err = imgconv.Save(output, imgSrc, &imgconv.FormatOption{Format: format})
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}
}
