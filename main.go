package main

import (
	"image"
	"image/color"
	"image/draw"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/sunshineplan/imgconv"
)

func setWhiteBackground(imgSrc image.Image) image.Image {
	newImg := image.NewRGBA(imgSrc.Bounds())

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)

	return newImg
}

func main() {
	args := os.Args

	output := args[1]+"-copy"
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
		default:
			log.Fatal(`Invalid arguments.
				Try: convertimg <file> [-o <output-filename>] [-f JPEG|PNG|BMP|GIF|TIFF] [-w <int>] [-h <int>]`)
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

	err = imgconv.Save(output, imgSrc, &imgconv.FormatOption{Format: format})
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}
}