package main

import (
	"flag"
	"log/slog"

	"github.com/samridh3215/YouStore/internal/encoder"
)

func main() {

	width := flag.Uint("w", 1080, "To define the width of the video")
	height := flag.Uint("h", 786, "To define the height of the video")
	frameRate := flag.Int("fr", 20, "Frame rate of the video")
	srcDir := flag.String("s", "", "Directory to read the files from")
	_ = flag.Int("m", 1, "0 to decode, 1 to encode")
	_ = flag.String("d", "", "Video to decode")
	flag.Parse()

	slog.Info("Client Started")
	slog.Info("Arguments passed", "flags", flag.Args())
	conf := encoder.EncoderConfig{
		FrameRate:     *frameRate,
		ImageWidth:    *width,
		ImageHeight:   *height,
		MultipleFiles: true,
	}
	data_encoder := encoder.NewDataEncoder(*srcDir, conf)
	data_encoder.Encode()

}
