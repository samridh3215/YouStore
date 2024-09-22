package main

import (
	"log/slog"

	"github.com/samridh3215/YouStore/internal/encoder"
)

func main() {
	slog.Info("Client Started")
	conf := encoder.EncoderConfig{
		FrameRate:     5,
		ImageWidth:    1080,
		ImageHeight:   786,
		MultipleFiles: true,
	}
	data_encoder := encoder.NewDataEncoder("/home/samridh/PicoCTF", conf)
	data_encoder.Encode()

}
