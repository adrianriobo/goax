package screenshot

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/kbinani/screenshot"
)

var seq int = 0

func CaptureScreen(outputPath, outputFilename string) error {
	return CaptureBounds(screenshot.GetDisplayBounds(0), outputPath, outputFilename)
}

func CaptureBounds(rect image.Rectangle, outputPath, outputFilename string) error {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return fmt.Errorf("error capturing the screen: %v", err)
	}
	fileName := fmt.Sprintf("%s-%s.png", outputFilename, screenshotSequece())
	file, err := os.Create(filepath.Join(outputPath, fileName))
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

func screenshotSequece() string {
	s := fmt.Sprintf("%d", seq)
	seq++
	return s
}
