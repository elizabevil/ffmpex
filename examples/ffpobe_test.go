package examples

import (
	"fmt"
	"github.com/elizabevil/ffmpegx/transcoderx"
	"strconv"
	"strings"
	"testing"
)

func TestProbe(t *testing.T) {
	transcode := transcoderx.NewTranscoder(transcoderx.NewConfig(func(config *transcoderx.Config) {
		config.FFprobeBin = FfprobeBin
	}))
	metadata, err := transcode.GetMetadata(InputVedio)
	if err != nil {
		panic(err)
	}
	fmt.Println(metadata.Format)
	for _, stream := range metadata.Streams {
		fmt.Println(stream.Index, stream.CodecName, stream.CodecType, stream.StreamVideoOnly.Width)
	}
}
func TestProgressSpeed(t *testing.T) {
	float, err := strconv.ParseFloat(strings.TrimRight("5.73e+03x", "x"), 10)
	fmt.Println(err, float)
}
