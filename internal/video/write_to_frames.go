package video

import (
	"fmt"
	"image"
	"image/color"
	"strings"
)

func bytesToBinary(data []byte) string {
	var builder strings.Builder
	builder.Grow(len(data) * 8)

	for _, b := range data {
		builder.WriteString(fmt.Sprintf("%08b", b))
	}
	return builder.String()
}
func WriteToFrames(data []byte, frameWidth, frameHeight int) []image.Image {
	frames := []image.Image{}
	binaryString := bytesToBinary(data)
	noOfFrames := len(binaryString) / (frameHeight * frameWidth)
	leftover := len(binaryString) % (frameHeight * frameWidth)
	for i := 0; i < noOfFrames; i++ {
		frame := image.NewGray(image.Rect(0, 0, frameWidth, frameHeight))
		frameContent := ""
		if i == noOfFrames-1 {
			frameContent = binaryString[i*frameHeight*frameWidth:]
		}
		frameContent = binaryString[i*frameHeight*frameWidth : (i+1)*frameHeight*frameWidth]
		for c := 0; c < frameHeight; c++ {
			for r := 0; r < frameWidth; r++ {
				bit := frameContent[c*frameHeight+r]
				if bit == '0' {
					frame.SetGray(r, c, color.Gray{Y: 0}) 
				} else {
					frame.SetGray(r, c, color.Gray{Y: 255}) 
				}
			}
		}
		frames = append(frames, frame)
	}
	if leftover > 0 {
		frame := image.NewGray(image.Rect(0, 0, frameWidth, frameHeight))
		frameContent := binaryString[noOfFrames*frameHeight*frameWidth:]

		// Fill the remaining pixels
		for i := 0; i < len(frameContent); i++ {
			c := i / frameWidth
			r := i % frameWidth

			bit := frameContent[i]
			if bit == '0' {
				frame.SetGray(r, c, color.Gray{Y: 0}) // Black pixel
			} else {
				frame.SetGray(r, c, color.Gray{Y: 255}) // White pixel
			}
		}
		frames = append(frames, frame)
	}
	return frames
}
