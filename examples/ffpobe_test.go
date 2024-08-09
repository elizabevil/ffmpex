package examples

import (
	"encoding/xml"
	"fmt"
	"github.com/elizabevil/ffmpegx/paramx"
	"github.com/elizabevil/ffmpegx/paramx/optionx"
	"github.com/elizabevil/ffmpegx/paramx/typex"
	"github.com/elizabevil/ffmpegx/transcoderx"
	"strconv"
	"strings"
	"testing"
)

func TestProbe(t *testing.T) {
	transcoderx.Debug = true
	transcode := transcoderx.NewTranscoder(transcoderx.NewConfig(func(config *transcoderx.Config) {
		config.FFprobeBin = FfprobeBin
	}))
	metadata, err := transcode.Metadata(InputVedio)
	if err != nil {
		panic(err)
	}
	fmt.Println(metadata.Format)
	for _, stream := range metadata.Streams {
		fmt.Println(stream.Index, stream.CodecType, stream.CodecName, stream.CodecLongName, stream.StreamVideoOnly.Width)
	}
}
func TestFFprobeParams(t *testing.T) {
	fprobe := optionx.FFprobe{
		I: InputVedio,
		//PrintFormat:         optionx.OutputFormat_json,
		PrintFormat:         optionx.OutputFormat_xml,
		ShowData:            false, //packets_and_frames data
		ShowError:           false,
		ShowFormat:          false,
		ShowFrames:          false,
		ShowPackets:         false,
		ShowPrograms:        true,
		ShowStreamGroups:    false,
		ShowStreams:         true,
		ShowChapters:        true,
		ShowProgramVersion:  false,
		ShowLibraryVersions: false,
		ShowVersions:        false, //all version
		ShowPixelFormats:    false,
		ShowPrivateData:     false,
		CountFrames:         true,
		CountPackets:        true,
		//ShowDataHash:        flagx.Hash_MD5,
		//ShowEntries: "stream=codec_type:stream_disposition=default", //    stream>>disposition>>default
		ShowLog: 0,
		//ShowOptionalFields: false,
	}
	fmt.Println(fprobe.Args())
	//x, err := transcoderx.MetadataWithArgs(FfprobeBin, paramx.BuildIArgInterface(fprobe), transcoderx.JsonUnmarshal)
	x, err := transcoderx.MetadataWithArgs(FfprobeBin, paramx.BuildIArgInterface(fprobe), XmlUnmarshal{})
	if err != nil {
		fmt.Println("MetadataX", err)
		return
	}
	fmt.Println(x.Format)
	fmt.Println(x.StreamType(typex.AVMEDIA_TYPE_VIDEO))
}

type XmlUnmarshal struct {
}

func (x XmlUnmarshal) Unmarshal(bytes []byte, data any) error {
	return xml.Unmarshal(bytes, data)
}

func TestProgressSpeed(t *testing.T) {
	float, err := strconv.ParseFloat(strings.TrimRight("5.73e+03x", "x"), 10)
	fmt.Println(err, float)
}
