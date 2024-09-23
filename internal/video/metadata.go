package video

type FileData struct {
	FileSize uint
	FileName string
	FileType string
	Position FramePosition
	Content  []byte `json:"-"`
}
