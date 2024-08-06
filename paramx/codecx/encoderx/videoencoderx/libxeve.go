package videoencoderx

import "github.com/elizabevil/ffmpegx/paramx/typex"

// // LIBXEVE 9.19 libxeve
type LIBXEVE struct {
	B typex.BitrateInt `json:"b" flag:"-b"` // (encoding,audio,video)
	//Set target video bitrate in bits/s. Note that FFmpeg’s b option is expressed in bits/s, while xeve’s bitrate is in kilobits/s.

	Bf typex.UNumber `json:"bf" flag:"-bf"` // (bframes)
	//Set the maximum number of B frames (1,3,7,15).

	G typex.Size `json:"g" flag:"-g"` // (keyint)
	//Set the GOP size (I-picture period).

	Preset int `json:"preset" flag:"-preset"` // (preset)
	//Set the xeve preset. Set the encoder preset value to determine encoding speed [fast, medium, slow, placebo]

	Tune int `json:"tune" flag:"-tune"` // (tune)
	//Set the encoder tune parameter [psnr, zerolatency]

	Profile int `json:"profile" flag:"-profile"` // (profile)
	//Set the encoder profile [0: baseline; 1: main]

	Crf typex.UI8 `json:"crf" flag:"-crf"` // (crf)
	//Set the quality for constant quality mode. Constant rate factor <10..49> [default: 32]

	Qp typex.Level `json:"qp" flag:"-qp"` // (qp)
	//Set constant quantization rate control method parameter. Quantization parameter qp <0..51> [default: 32]

	Threads typex.UNumber `json:"threads" flag:"-threads"` // (threads)
	//Force to use a specific number of threads

}
