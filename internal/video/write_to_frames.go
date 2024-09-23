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

func WriteToFrames(frameWidth, frameHeight int, metadata []FileData) ([]image.Image, []FileData) {
	frames := []image.Image{}
	totalPixels := frameHeight * frameWidth
	currentByte := uint(0)
	currentFrame := uint(0)

	for i := range metadata {
		binaryString := bytesToBinary(metadata[i].Content)
		remainingBits := len(binaryString)
		bitIndex := 0

		metadata[i].Position.FrameNumber = currentFrame
		metadata[i].Position.StartOffset = currentByte

		for remainingBits > 0 {
			frame := image.NewGray(image.Rect(0, 0, frameWidth, frameHeight))
			bitsInThisFrame := min(totalPixels, remainingBits)

			for j := 0; j < bitsInThisFrame; j++ {
				r := j % frameWidth
				c := j / frameWidth
				if binaryString[bitIndex+j] == '0' {
					frame.SetGray(r, c, color.Gray{Y: 0})
				} else {
					frame.SetGray(r, c, color.Gray{Y: 255})
				}
			}

			for j := bitsInThisFrame; j < totalPixels; j++ {
				r := j % frameWidth
				c := j / frameWidth
				frame.SetGray(r, c, color.Gray{Y: 0}) // Black pixel for padding
			}

			frames = append(frames, frame)
			bitIndex += bitsInThisFrame
			remainingBits -= bitsInThisFrame
			currentFrame++
			currentByte += uint(bitsInThisFrame / 8)
		}

		metadata[i].Position.EndOffset = currentByte
	}

	return frames, metadata
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
