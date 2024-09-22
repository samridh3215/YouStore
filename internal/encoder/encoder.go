package encoder

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/samridh3215/YouStore/internal/video"
)

var BytesWritten int

type FileData struct {
	fileSize uint
	fileName string
	fileType string
	position video.FramePosition
	content  []byte
}

type encoder struct {
	dataPath string
	config   EncoderConfig
}

func NewDataEncoder(dataPath string, config EncoderConfig) *encoder {

	slog.Info("Created encoder with ", "config", config)
	return &encoder{
		dataPath: dataPath,
		config:   config,
	}
}

func (enc *encoder) Encode() {
	files, readDirErr := os.ReadDir(enc.dataPath)
	if readDirErr != nil {
		slog.Error("Could not read files from directory", "msg", readDirErr.Error())
	}
	allBytes := []byte{}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		nameParts := strings.Split(f.Name(), ".")
		fileName := strings.Join(nameParts[:len(nameParts)-1], ".")
		fileType := nameParts[len(nameParts)-1]
		bytesData, readFileErr := os.ReadFile(filepath.Join(enc.dataPath, f.Name()))

		if readFileErr != nil {
			slog.Error("Could not read file in bytes", "msg", readFileErr.Error())
		}
		_ = FileData{
			fileSize: uint(len(bytesData)),
			fileName: fileName,
			fileType: fileType,
			content:  bytesData,
			position: video.FramePosition{
				FrameNumber: uint(BytesWritten) / uint(enc.config.ImageHeight*enc.config.ImageWidth),
				StartOffset: uint(BytesWritten),
				EndOffset:   uint(BytesWritten + len(bytesData)),
			},
		}
		allBytes = append(allBytes, bytesData...)

	}
	videoObject := video.Video{
		Frames: video.WriteToFrames(allBytes, int(enc.config.ImageWidth), int(enc.config.ImageHeight)),
	}
	videoObject.CreateFromFrames("sample.mp4", "output", enc.config.FrameRate)

}
