package muxerx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// ALP 21.10 alp ALP
// 4.10 alp
type ALP struct {
	Type ALPType `json:"type" flag:"-type"`
	//Set file type.

}
type ALPType = typex.String

const (
	tun = "tun"
	//Set file type as music. Must have a sample rate of 22050 Hz.

	pcm = "pcm"
	//Set file type as sfx.

	auto = "auto"
	//Set file type as per output file extension. .pcm results in type pcm else type tun is set. (default)

)
