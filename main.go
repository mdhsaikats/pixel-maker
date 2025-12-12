package main

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"golang.org/x/image/draw"
)

func main() {
	fmt.Println("## Application Started ##")
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the image path: ")
		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)
		path = strings.Trim(path, "\"")
		if path == "exit" {
			fmt.Println("Exiting application...")
			break
		}

		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			panic(err)
		}

		fmt.Println("Image loaded successfully")
		fmt.Println("Width: ", img.Bounds().Dx())
		fmt.Println("Height:", img.Bounds().Dy())

		smallWidth := 32
		smallHeight := 32
		smallImg := image.NewRGBA(image.Rect(0, 0, smallWidth, smallHeight))
		draw.NearestNeighbor.Scale(smallImg, smallImg.Bounds(), img, img.Bounds(), draw.Over, nil)

		scaleFactor := 10
		pixelated := image.NewRGBA(image.Rect(0, 0, smallWidth*scaleFactor, smallHeight*scaleFactor))
		draw.NearestNeighbor.Scale(pixelated, pixelated.Bounds(), smallImg, smallImg.Bounds(), draw.Over, nil)

		outFile, err := os.Create("pixelated.png")
		if err != nil {
			panic(err)
		}
		defer outFile.Close()

		err = png.Encode(outFile, pixelated)
		if err != nil {
			panic(err)
		}

		fmt.Println("Pixelated image saved as pixelated.png")
	}

}
