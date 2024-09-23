package encoder

type EncoderConfig struct {
	FrameRate     int
	ImageWidth    uint
	ImageHeight   uint
	MultipleFiles bool // If True, encodes all files in the directory. If False, then encodes the first file
}
