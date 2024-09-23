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

### Metadata 
THe metadata conatins the file names, size, start and end postiion which can be used for decoding 
Example
```json
[{"FileSize":87282,"FileName":"Resumee","FileType":"pdf","Position":{"FrameNumber":0,"StartOffset":0,"EndOffset":87282}},{"FileSize":1011218,"FileName":"What if...","FileType":"png","Position":{"FrameNumber":1,"StartOffset":87282,"EndOffset":1098500}},{"FileSize":292599,"FileName":"combination_sum_recursion_graph","FileType":"png","Position":{"FrameNumber":11,"StartOffset":1098500,"EndOffset":1391099}},{"FileSize":775,"FileName":"demo","FileType":"js","Position":{"FrameNumber":14,"StartOffset":1391099,"EndOffset":1391874}},{"FileSize":1646,"FileName":"matrix","FileType":"py","Position":{"FrameNumber":15,"StartOffset":1391874,"EndOffset":1393520}},{"FileSize":1455,"FileName":"placing","FileType":"txt","Position":{"FrameNumber":16,"StartOffset":1393520,"EndOffset":1394975}},{"FileSize":2363,"FileName":"server","FileType":"py","Position":{"FrameNumber":17,"StartOffset":1394975,"EndOffset":1397338}},{"FileSize":77890,"FileName":"subsets_recursion_graph","FileType":"png","Position":{"FrameNumber":18,"StartOffset":1397338,"EndOffset":1475228}},{"FileSize":334856,"FileName":"wallpaperflare.com_wallpaper","FileType":"jpg","Position":{"FrameNumber":19,"StartOffset":1475228,"EndOffset":1810084}}]
```
