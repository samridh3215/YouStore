## Flags:

- **Width (`-w`)**
  Type: `uint`
  Default: 1080
  Description: Defines the width of the video in pixels.
- **Height (`-h`)**
  Type: `uint`
  Default: 786
  Description: Defines the height of the video in pixels.
- **Frame Rate (`-fr`)**
  Type: `int`
  Default: 20
  Description: Sets the frame rate of the video, i.e., the number of frames per second.
- **Source Directory (`-s`)**
  Type: string
  Default: ""
  Description: Specifies the directory path from which the video files are to be read.


### Complete usage
```bash
go run cmd/src/main.go -w 1280 -h 720 -fr 25 -s /path/to/files
```
