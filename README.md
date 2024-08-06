# # FFmpex

> ###### go control ffmpeg program

Encapsulates commonly used methods and parameters

> All From =>FFmpegDoc https://www.ffmpeg.org/documentation.html

---

### Usage notes

> go get github.com/elizabevil/ffmpegx@version

- args

```go
//ffmpeg [global_options] {[input_file_options] -i input_url} ... {[output_file_options] output_url} ...

```

- Make args

```go
//make cmd args
//https://www.ffmpeg.org/documentation.html

input := optionx.Input{Input: InputVedio}
output := optionx.Output{Output: "output.mp4"}
args := paramx.AnyArgs{"-c": "copy"} //yourself free args
//argsx := typex.Args{"-c", "copy"}  // or this
hls := muxerx.HLS{
    HlsFlags:           flagx.HlsFlags_delete_segments,
	//HlsTime:            typex.TimeDurationI(61 * time.Second),
    //HlsTime:            typex.TimeDurationParseI("00:01:01"),
    HlsTime:            typex.TimeDurationParseI("1m1s"), // time 
    HlsDeleteThreshold: 1,
    HlsListSize:        &typex.ZeroUN, //default 5 --> force 0
    HlsSegmentFilename: "%Y%m%d-%s.ts",
}

//ffmpeg [global_options] {[input_file_options] -i input_url} ... {[output_file_options] output_url} ...

argInterface := paramx.BuildIArgInterface(input, args, hls, output)
fmt.Println(argInterface.Args())
//[-i ./file/input.mp4 -c copy -hls_flags delete_segments -hls_segment_filename %Y%m%d-%s.ts -hls_delete_threshold 1 -hls_list_size 5 -hls_time 10 output.mp4]

// or you can impl Arg interface
type Arg interface {
    Args() typex.Args
}
```

- YourParse Args

```
type YourParse struct {// your paesr method
}

func (y YourParse) ParamParse(input any) typex.Args {
	return []string{"-c", "copy"}
}

func (y YourParse) ParamItemType(of reflect.Value) (string, bool) {
	return "c", false
}
//================
type YourHls muxerx.HLS // your type
func (r YourHls) Args() typex.Args {
	return YourParse{}.ParamParse(r)
}
fmt.Println(YourHls{}.Args()) //see args
//paramx.BuildIArgInterface(input, args, hls, output)
paramx.BuildIArgInterface(input, args, YourHls{}, output) // complete args
```

- Get Metadata

```go
metadata, err := transcoderx.Metadata(ffprobeBinPath, InputVedio) // use ffprobeBin to get InputVedio's Metadata 
fmt.Println(metadata.Format)  // println Format
fmt.Println(metadata.Streams) // println Streams
```

- Audio Channel_layout

```go
stream := metadata.Streams[1] // if this stream `s codec_type is audio
layout := stream.GetChannelLayout() //{stereo {1 2 3}}
fmt.Println(metadatax.ChannelNames[layout.Layout.Mask]) //{LFE low frequency}
```

- Directly run Or transcoder run

```go
transcoderx.Cmd(ffmpegBinPath, args) //  run run with args

transcode := transcoderx.NewTranscode(transcoderx.NewConfig())// with default config
tran.Args = makeParams()                                      // set args
transcode.Cmd() //run

handleFun, err := transcode.PipelineX() //run
handleFun()
```

- Use transcoder to Run

```go
// make transcode
transcoderx.SetMode(true) // show full cmd

transcode := transcoderx.NewTranscode(transcoderx.NewConfig()) // with default config
tran.Args = makeParams() // must set cmd params

pipline, err := transcoder.Pipeline(func (cmd *exec.Cmd) {
    getwd, _ := os.Getwd()
    cmd.Dir = filepath.Join(getwd, "file") // set cmd dir 
})
if err != nil {
    log.Panicln(err)
}
//transcode.Ph=interfacex.ProgressHandle // set your self method to parse progress
for progress := range pipline { // progress chan
    //progress {271 181 -1.0 N/A 00:00:09.01 N/A  6.02x +Inf}
    fmt.Println("progress", progress)
}
//[FFmpegx-Debug] /usr/bin/ffmpeg -y -re -i ./file/input.mp4 -c copy -vcodec libx264 ./file/output.mp4
```

- some types fun

```go
fmt.Println(typex.TimeZero) //0000-01-01 00:00:00 +0000 UTC
duration := 3*time.Second + 500*time.Millisecond //3.5s
fmt.Println(typex.TimeDuration(duration, time.Millisecond)) //3500 Millisecond
fmt.Println(typex.TimeDurationSecond(duration))//3.5s
fmt.Println(typex.TimeDurationSecondI(duration))//3 s
fmt.Println(typex.TimeDurationParseSecondF("200ms"))//0.2s
fmt.Println(typex.TimeDurationParseSecondF("200000us"))//0.2s
fmt.Println(typex.TimeDurationParseSecondF("00:01:01"))//61s
fmt.Println(typex.TimeDurationParseSecondF("61s"))//61s


fmt.Println(typex.NewVideoSize(1024, 720))//720x1024
fmt.Println(typex.NewRatio(1024, 720))//720:1024
fmt.Println(typex.NewRate(25, 1))  //25/1
```

- More examples ==>see examples dir 