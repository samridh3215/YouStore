package video

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

type Video struct {
	Frames   []image.Image
	MetaData []FileData
}

func removeAllFilesInDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %v", err)
	}
	for _, file := range files {
		err := os.Remove(filepath.Join(dir, file.Name()))
		if err != nil {
			return fmt.Errorf("failed to remove file %s: %v", file.Name(), err)
		}
	}
	return nil
}

func writeMetaDataToFile(metaDataSting string) {

}

func (v *Video) CreateFromFrames(videoName string, outputDir string, frameRate int) error {
	_, existErr := os.Open(videoName)
	if os.IsExist(existErr) {
		slog.Error("Video name already exists. Choose other name", "name", videoName)
		fmt.Scanln(&videoName)
	}
	if _, err := os.Stat(outputDir); !os.IsNotExist(err) {
		err := removeAllFilesInDir(outputDir)
		if err != nil {
			return fmt.Errorf("failed to clean output directory: %v", err)
		}
	} else {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create output directory: %v", err)
		}
	}

	for i, frame := range v.Frames {
		fileName := fmt.Sprintf("%s/frame_%04d.png", outputDir, i)
		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("failed to create image file %s: %v", fileName, err)
		}
		defer file.Close()
		err = png.Encode(file, frame)
		if err != nil {
			return fmt.Errorf("failed to encode image %s: %v", fileName, err)
		}
	}

	// Prepare metadata
	metadataJSON, err := json.Marshal(v.MetaData)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %v", err)
	}
	_ = base64.StdEncoding.EncodeToString(metadataJSON)

	// Construct FFmpeg command with metadata as a comment
	cmd := exec.Command(
    "ffmpeg",
    "-framerate", fmt.Sprintf("%d", frameRate),
    "-i", filepath.Join(outputDir, "frame_%04d.png"),
    "-c:v", "mpeg4",  // Fallback to mpeg4 codec
    "-q:v", "5",      // Adjust quality (lower values mean higher quality)
    "-metadata", fmt.Sprintf("comment=%s", string(metadataJSON)),
    videoName,
)

	cmd.Stdout = os.Stdout	
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to create video: %v", err)
	}

	fmt.Println("Video created successfully:", videoName)
	return nil
}
