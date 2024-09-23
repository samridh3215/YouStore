package encoder

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/samridh3215/YouStore/internal/video"
)

var BytesWritten int

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
	var fileDataList []video.FileData
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
		metadata := video.FileData{
			FileSize: uint(len(bytesData)),
			FileName: fileName,
			FileType: fileType,
			Content:  bytesData,
		}
		fileDataList = append(fileDataList, metadata)
		allBytes = append(allBytes, bytesData...)

	}
	frames, updatedMetaData := video.WriteToFrames(int(enc.config.ImageWidth), int(enc.config.ImageHeight), fileDataList)
	videoObject := video.Video{
		Frames:   frames,
		MetaData: updatedMetaData,
	}
	err := videoObject.CreateFromFrames("sample.mp4", "output", enc.config.FrameRate)
	if err != nil {
		slog.Error("Error creating video", "msg", err.Error())
	}

}
