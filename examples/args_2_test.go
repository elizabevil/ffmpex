package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/paramx/flagx"
	"github.com/elizabevil/ffmpegx/paramx/formatx/muxerx"
	"github.com/elizabevil/ffmpegx/paramx/optionx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"github.com/elizabevil/ffmpegx/transcoderx/interfacex"
	"testing"
)

const FfmpegBin = "/usr/bin/ffmpeg"
const FfprobeBin = "/usr/bin/ffprobe"
const InputVedio = "./file/input.mp4"
const OutputVideo = "./file/output.mp4"
const OutputVideoHls = "./file/output.m3u8"

func hlsParams() interfacex.IArgs {
	input := optionx.Input{Input: InputVedio}
	output := optionx.Output{Output: OutputVideoHls}
	args := paramx.AnyArgs{"-c": "copy"}
	//argsx := typex.Args{"-c", "copy"}
	hls := muxerx.HLS{
		HlsFlags:    flagx.HlsFlags_delete_segments,
		HlsListSize: &typex.ZeroUN,
		//HlsTime:            typex.TimeDurationSecondI(61 * time.Second),
		//HlsTime:            typex.TimeDurationParseSecondI("00:01:01"),
		HlsTime:            typex.TimeDurationParseSecondI("0m2s"),
		Strftime:           typex.True,
		HlsDeleteThreshold: 1,
		HlsSegmentFilename: "%Y%m%d%H%M%S.ts",
	}
	argInterface := paramx.BuildIArgInterface(input, args, hls, output)
	return argInterface
}
func TestArgsParse(t *testing.T) {
	fmt.Println(hlsParams().Args())
	//[-i ./file/input.mp4 -c copy -hls_flags delete_segments -strftime 1 -hls_segment_filename %Y%m%d%H%M%S.ts -hls_delete_threshold 1 -hls_list_size 0 -hls_time 2 ./file/output.m3u8]
}
