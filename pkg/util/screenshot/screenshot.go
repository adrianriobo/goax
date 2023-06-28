package screenshot

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func CaptureScreen(outputPath, outputFilename string) error {
	return CaptureBounds(screenshot.GetDisplayBounds(0), outputPath, outputFilename)
}

func CaptureBounds(rect image.Rectangle, outputPath, outputFilename string) error {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return fmt.Errorf("error capturing the screen: %v", err)
	}
	file, err := os.Create("test.png")
	if err != nil {
		return fmt.Errorf("error creating the file: %v", err)
	}
	defer file.Close()
	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("error encoding the capture: %v", err)
	}
	return nil
}
